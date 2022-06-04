package async

import (
	"context"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hpb-project/srng-service/cache"
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

type logHandler func(log types.Log, client *ethclient.Client, addr common.Address) error

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
		oracleAddr := beego.AppConfig.String("oracleAddr")
		{
			if len(oracleAddr) > 0 {
				addr := common.HexToAddress(oracleAddr)
				pullTask.contractList = append(pullTask.contractList, addr)
				pullTask.contractHandler[addr] = OracleContractHandler
				log.Info("add contract to fillter log", "address ", oracleAddr)
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
					handle(vlog, p.client, vlog.Address)
				}
			}
		}
		p.lastBlock = new(big.Int).Add(query.ToBlock, bigOne)
		cache.Redis.Set(LastSyncBlockKey, p.lastBlock.Text(10))
	}
}
