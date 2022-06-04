package async

import (
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hpb-project/srng-service/contracts"
	"strings"
)

func OracleContractHandler(vLog types.Log, client *ethclient.Client, addr common.Address) error {
	logs.Info("handler arena contract logs")
	filter, err := contracts.NewOracleFilterer(addr, client)
	if err != nil {
		logs.Error("NewOracleFilter failed", "err", err)
		return err
	}
	{
		method := vLog.Topics[0]
		switch strings.ToLower(method.String()) {
		case EventSubscribe:
			sub, err := filter.ParseSubscribe(vLog)
			if err != nil {
				logs.Error("parse subscribe event failed", "err", err)
				return err
			}
			logs.Info("got event subscribe", sub)
		case EventUnSubscribe:
			unsub, err := filter.ParseUnSubscribe(vLog)
			if err != nil {
				logs.Error("parse unsubscribe event failed", "err", err)
				return err
			}
			logs.Info("got event unsubscribe", unsub)
		case EventCommitHash:
			commit, err := filter.ParseCommitHash(vLog)
			if err != nil {
				logs.Error("parse commit event failed", "err", err)
				return err
			}
			logs.Info("got event commit", commit)
		case EventRevealSeed:
			reveal, err := filter.ParseRevealSeed(vLog)
			if err != nil {
				logs.Error("parse reveal event failed", "err", err)
				return err
			}
			logs.Info("got event reveal", reveal)
		case EventRandomConsumed:
			consume, err := filter.ParseRandomConsumed(vLog)
			if err != nil {
				logs.Error("parse consume event failed", "err", err)
				return err
			}
			logs.Info("got event consume", consume)
		}
	}
	return nil
}
