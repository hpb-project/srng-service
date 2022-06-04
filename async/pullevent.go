package async

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/common/log"
	"math/big"
	"strings"
	"time"
)

const (
	LastSyncBlockKey = "lastSyncBlock"
)

var (
	pullTask   *PullEvent
	bigOne     = big.NewInt(1)
	bigTen     = big.NewInt(10)
	bighundred = big.NewInt(100)
	bigK       = big.NewInt(1000)
)

type logHandler func(log types.Log) error

//type ContractEvent struct {
//	Addr common.Address
//	Handler logHandler
//}

type PullEvent struct {
	ctx             context.Context
	client          *ethclient.Client
	lastBlock       *big.Int
	contractList    []common.Address
	contractHandler map[common.Address]logHandler
}

func getAddress(multiAddrs string) []string {
	addrs := strings.Split(multiAddrs, ",")
	ret := make([]string, 0)
	for _, addr := range addrs {
		if len(addr) > 0 {
			ret = append(ret, addr)
		}
	}
	return ret
}

func init() {
	var err error
	rpc := beego.AppConfig.String("url")
	deployBlock := beego.AppConfig.String("deployedBlock")
	fmt.Println("get rpc ", rpc)
	pullTask = &PullEvent{contractHandler: make(map[common.Address]logHandler)}
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(fmt.Sprintf("ethclient dial failed, err:", err))
	} else {
		pullTask.client = client
		lastBlock := cache.Redis.Get(LastSyncBlockKey)
		if len(lastBlock) == 0 {
			lastBlock = deployBlock
		}
		{
			blockNumber, _ := new(big.Int).SetString(lastBlock, 10)
			pullTask.lastBlock = blockNumber
		}
	}
	pullTask.ctx = context.Background()
	pullTask.contractList = make([]common.Address, 0)
	{
		arenaAddrs := getAddress(beego.AppConfig.String("arenaContract"))
		for _, arenaAddr := range arenaAddrs {
			if len(arenaAddr) > 0 {
				addr := common.HexToAddress(arenaAddr)
				pullTask.contractList = append(pullTask.contractList, addr)
				pullTask.contractHandler[addr] = HorseArenaContractHandler
				log.Info("add contract to fillter log", "address ", arenaAddr)
			}
		}

		arenaExtraAddrs := getAddress(beego.AppConfig.String("arenaExtraContract"))
		for _, arenaExtraAddr := range arenaExtraAddrs {
			if len(arenaExtraAddr) > 0 {
				addr := common.HexToAddress(arenaExtraAddr)
				pullTask.contractList = append(pullTask.contractList, addr)
				pullTask.contractHandler[addr] = HorseArenaExtraHandler
				log.Info("add contract to fillter log", "address ", arenaExtraAddr)
			}
		}

		arenaExtra2Addrs := getAddress(beego.AppConfig.String("arenaExtra2Contract"))
		for _, arenaExtra2Addr := range arenaExtra2Addrs {
			if len(arenaExtra2Addr) > 0 {
				addr := common.HexToAddress(arenaExtra2Addr)
				pullTask.contractList = append(pullTask.contractList, addr)
				pullTask.contractHandler[addr] = HorseArenaExtra2Handler
				log.Info("add contract to fillter log", "address ", arenaExtra2Addr)
			}
		}

		arenaExtra3Addrs := getAddress(beego.AppConfig.String("arenaExtra3Contract"))
		for _, arenaExtra3Addr := range arenaExtra3Addrs {
			if len(arenaExtra3Addr) > 0 {
				addr := common.HexToAddress(arenaExtra3Addr)
				pullTask.contractList = append(pullTask.contractList, addr)
				pullTask.contractHandler[addr] = HorseArenaExtra3Handler
				log.Info("add contract to fillter log", "address ", arenaExtra3Addr)
			}
		}

		horseEquipAddrs := getAddress(beego.AppConfig.String("horseEquipContract"))
		for _, horseEquipAddr := range horseEquipAddrs {
			if len(horseEquipAddr) > 0 {
				addr := common.HexToAddress(horseEquipAddr)
				pullTask.contractList = append(pullTask.contractList, addr)
				pullTask.contractHandler[addr] = HorseEquipContractHandler
				log.Info("add contract to fillter log", "address ", horseEquipAddr)
			}
		}

		horseRaceAddrs := getAddress(beego.AppConfig.String("horseRaceContract"))
		for _, horseRaceAddr := range horseRaceAddrs {
			if len(horseRaceAddr) > 0 {
				addr := common.HexToAddress(horseRaceAddr)
				pullTask.contractList = append(pullTask.contractList, addr)
				pullTask.contractHandler[addr] = HorseRaceContractHandler
				log.Info("add contract to fillter log", "address ", horseRaceAddr)
			}
		}

		horseRaceExtraAddrs := getAddress(beego.AppConfig.String("horseRaceExtra1Contract"))
		for _, horseRaceExtraAddr := range horseRaceExtraAddrs {
			if len(horseRaceExtraAddr) > 0 {
				addr := common.HexToAddress(horseRaceExtraAddr)
				pullTask.contractList = append(pullTask.contractList, addr)
				pullTask.contractHandler[addr] = HorseRaceExtra1Handler
				log.Info("add contract to fillter log", "address ", horseRaceExtraAddr)
			}
		}

		horseRaceExtra2Addrs := getAddress(beego.AppConfig.String("horseRaceExtra2Contract"))
		for _, horseRaceExtra2Addr := range horseRaceExtra2Addrs {
			if len(horseRaceExtra2Addr) > 0 {
				addr := common.HexToAddress(horseRaceExtra2Addr)
				pullTask.contractList = append(pullTask.contractList, addr)
				pullTask.contractHandler[addr] = HorseRaceExtra2Handler
				log.Info("add contract to fillter log", "address ", horseRaceExtra2Addr)
			}
		}
	}
}

func SyncLogs() {
	pullTask.GetLogs()
}

func (p *PullEvent) GetLogs() {
	query := ethereum.FilterQuery{}
	query.FromBlock = p.lastBlock
	query.ToBlock = new(big.Int).Add(p.lastBlock, big.NewInt(1))
	query.Addresses = p.contractList
	for {
		query.FromBlock = p.lastBlock

		log.Info("start fileter start at ", p.lastBlock.Text(10))
		height, err := p.client.BlockNumber(p.ctx)
		if height <= p.lastBlock.Uint64() {
			time.Sleep(time.Second)
			continue
		} else if (height - 1000) >= p.lastBlock.Uint64() {
			query.ToBlock = new(big.Int).Add(p.lastBlock, bigK)
		} else if (height - 100) >= p.lastBlock.Uint64() {
			query.ToBlock = new(big.Int).Add(p.lastBlock, bighundred)
		} else if (height - 10) >= p.lastBlock.Uint64() {
			query.ToBlock = new(big.Int).Add(p.lastBlock, bigTen)
		} else {
			query.ToBlock = new(big.Int).Add(p.lastBlock, bigOne)
		}

		allLogs, err := p.client.FilterLogs(p.ctx, query)
		if err != nil {
			log.Error("filter logs failed", err)
			continue
		}
		if len(allLogs) > 0 {
			for _, vlog := range allLogs {
				handle, exist := p.contractHandler[vlog.Address]
				if exist {
					handle(vlog)
				}
			}
		}
		p.lastBlock = new(big.Int).Add(query.ToBlock, bigOne)
		cache.Redis.Set(LastSyncBlockKey, p.lastBlock.Text(10))
	}
}
