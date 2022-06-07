package async

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hpb-project/srng-service/cache"
	"github.com/hpb-project/srng-service/contracts"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"math/big"
	"net/http"
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

func GetTokenInfo() map[string]interface{} {
	return pullTask.GetTokenInfo()
}

type HoldInfo struct {
	Data json.RawMessage `json:"data"`
	Code int	`json:"code"`
	Msg string `json:"enMsg"`
}

type RequestParam struct {
	Start string  	`json:"start"`
	Length string 	    `json:"length"`
	TokenAddr string `json:"token_address"`
	TokenType int `json:"token_types"`
}

func (p *PullEvent) GetTokenInfo() map[string]interface{} {
	var info = make(map[string]interface{})
	info["name"] = "HRG Token"
	info["symbol"] = "HRG"
	info["supply"] = "0"
	info["csupply"] = "0"
	info["price"] = "0.0"
	info["holdlist"] = nil

	tokenAddr := beego.AppConfig.String("tokenAddr")
	{
		addr := common.HexToAddress(tokenAddr)
		hrg,_ := contracts.NewToken(addr, p.client)
		unit := new(big.Int).SetInt64(1000000000000000000)
		opt := bind.CallOpts{}
		supply,_ := hrg.TotalSupply(&opt)
		supply = new(big.Int).Div(supply, unit)
		csupply := new(big.Int).Sub(supply,big.NewInt(50000000))

		info["supply"] = supply.Text(10)
		info["csupply"] = csupply.Text(10)
	}
	{
		param := RequestParam{
			Start: "1",
			Length: "10",
			TokenAddr: strings.ToLower(tokenAddr),
			TokenType: 1,
		}
		body,_ := json.Marshal(param)
		logs.Debug("request param", "body", string(body))
		resp, err := http.Post("https://hscan.org/blockBrowser/tokens/address/tokenHolder", "application/json", bytes.NewBuffer(body))
		if err != nil {
			logs.Error("get holder list failed")
			return info
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//Failed to read response.
			return info
		}
		var holdinfo = HoldInfo{}
		json.Unmarshal(data, &holdinfo)
		if holdinfo.Code != 0 {
			logs.Error("get holder info failed", "err", holdinfo.Msg)
			return info
		}
		info["holdlist"] = holdinfo.Data
	}
	return info
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
