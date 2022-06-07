package async

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hpb-project/srng-service/contracts"
	"strings"
)

func OracleContractHandler(vLog types.Log, client *ethclient.Client, addr common.Address) error {
	logs.Info("handler oracle contract logs")
	ormer := orm.NewOrm()
	filter, err := contracts.NewOracleFilterer(addr, client)
	if err != nil {
		logs.Error("NewOracleFilter failed", "err", err)
		return err
	}
	{
		method := vLog.Topics[0]
		logs.Info("got event ", vLog)
		switch strings.ToLower(method.String()) {
		case EventSubscribe:
			sub, err := filter.ParseSubscribe(vLog)
			if err != nil {
				logs.Error("parse subscribe event failed", "err", err)
				return err
			}
			logs.Info("got event subscribe", sub)
			tbSubscribe := ToTbSubScribe(&vLog, sub)
			_, err = ormer.Insert(tbSubscribe)
			if err != nil {
				beego.Error("ormer.Insert(tbSubscribe) error ", err.Error())
				return err
			}

		case EventUnSubscribe:
			unsub, err := filter.ParseUnSubscribe(vLog)
			if err != nil {
				logs.Error("parse unsubscribe event failed", "err", err)
				return err
			}
			logs.Info("got event unsubscribe", unsub)
			tbUnSubscribe := ToTbUnSubScribe(&vLog, unsub)
			_, err = ormer.Insert(tbUnSubscribe)
			if err != nil {
				beego.Error("ormer.Insert(tbUnSubscribe) error ", err.Error())
				return err
			}
		case EventCommitHash:
			commit, err := filter.ParseCommitHash(vLog)
			if err != nil {
				logs.Error("parse commit event failed", "err", err)
				return err
			}
			logs.Info("got event commit", commit)
			tbCommitHash := ToTbCommitHash(&vLog, commit)
			_, err = ormer.Insert(tbCommitHash)
			if err != nil {
				beego.Error("ormer.Insert(tbCommitHash) error ", err.Error())
				return err
			}
		case EventRevealSeed:
			reveal, err := filter.ParseRevealSeed(vLog)
			if err != nil {
				logs.Error("parse reveal event failed", "err", err)
				return err
			}
			logs.Info("got event reveal", reveal)
			tbReveal := ToTbReveal(&vLog, reveal)
			_, err = ormer.Insert(tbReveal)
			if err != nil {
				beego.Error("ormer.Insert(tbRandomConsume) error ", err.Error())
				return err
			}
		case EventRandomConsumed:
			consume, err := filter.ParseRandomConsumed(vLog)
			if err != nil {
				logs.Error("parse consume event failed", "err", err)
				return err
			}
			logs.Info("got event consume", consume)
			tbRandomConsume := ToTbRandomConsumed(&vLog, consume)
			_, err = ormer.Insert(tbRandomConsume)
			if err != nil {
				beego.Error("ormer.Insert(tbRandomConsume) error ", err.Error())
				return err
			}
		}
	}
	return nil
}
