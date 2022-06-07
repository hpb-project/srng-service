package async

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hpb-project/srng-service/contracts"
	"github.com/hpb-project/srng-service/db"
	"github.com/hpb-project/srng-service/utils"
	"strings"
	"time"
)

func ToTbSubScribe(vLog *types.Log, sub *contracts.OracleSubscribe) *db.TbSubscribe {
	tbSubscribe := new(db.TbSubscribe)
	tbSubscribe.Id = utils.GetSnowflakeId()
	tbSubscribe.BlockHash = strings.ToLower(vLog.BlockHash.String())
	tbSubscribe.TxTime = time.Unix(sub.Time.Int64(), 0)
	tbSubscribe.TxHash = strings.ToLower(vLog.TxHash.String())
	tbSubscribe.BlockId = int64(vLog.BlockNumber)
	tbSubscribe.Commiter = strings.ToLower(sub.Commiter.String())
	tbSubscribe.CommitHash = strings.ToLower(common.BytesToHash(sub.Hash[:]).String())
	tbSubscribe.Consumer = strings.ToLower(sub.Consumer.String())

	tbSubscribe.Status = 1
	tbSubscribe.IsDeleted = 0
	return tbSubscribe
}

func ToTbUnSubScribe(vLog *types.Log, unsub *contracts.OracleUnSubscribe) *db.TbUnSubscribe {
	tbUnSubscribe := new(db.TbUnSubscribe)
	tbUnSubscribe.Id = utils.GetSnowflakeId()
	tbUnSubscribe.BlockHash = strings.ToLower(vLog.BlockHash.String())
	tbUnSubscribe.TxTime = time.Unix(unsub.Time.Int64(), 0)
	tbUnSubscribe.TxHash = strings.ToLower(vLog.TxHash.String())
	tbUnSubscribe.BlockId = int64(vLog.BlockNumber)
	tbUnSubscribe.Commiter = strings.ToLower(unsub.Commiter.String())
	tbUnSubscribe.CommitHash = strings.ToLower(common.BytesToHash(unsub.Hash[:]).String())
	tbUnSubscribe.Consumer = strings.ToLower(unsub.Consumer.String())

	tbUnSubscribe.Status = 1
	tbUnSubscribe.IsDeleted = 0
	return tbUnSubscribe
}

func ToTbCommitHash(vLog *types.Log, commit *contracts.OracleCommitHash) *db.TbCommitHash {
	tbCommitHash := new(db.TbCommitHash)
	tbCommitHash.Id = utils.GetSnowflakeId()
	tbCommitHash.BlockHash = strings.ToLower(vLog.BlockHash.String())
	tbCommitHash.TxTime = time.Unix(commit.Time.Int64(), 0)
	tbCommitHash.TxHash = strings.ToLower(vLog.TxHash.String())
	tbCommitHash.BlockId = int64(vLog.BlockNumber)
	tbCommitHash.Commiter = strings.ToLower(commit.Sender.String())
	tbCommitHash.CommitHash = strings.ToLower(common.BytesToHash(commit.Hash[:]).String())

	tbCommitHash.Status = 1
	tbCommitHash.IsDeleted = 0
	return tbCommitHash
}

func ToTbRandomConsumed(vLog *types.Log, consumed *contracts.OracleRandomConsumed) *db.TbRandomConsumed {
	tbRandomConsumed := new(db.TbRandomConsumed)
	tbRandomConsumed.Id = utils.GetSnowflakeId()
	tbRandomConsumed.BlockHash = strings.ToLower(vLog.BlockHash.String())
	tbRandomConsumed.TxTime = time.Unix(consumed.Time.Int64(), 0)
	tbRandomConsumed.TxHash = strings.ToLower(vLog.TxHash.String())
	tbRandomConsumed.BlockId = int64(vLog.BlockNumber)
	tbRandomConsumed.Committer = strings.ToLower(consumed.Commiter.String())
	tbRandomConsumed.Random = strings.ToLower(common.BytesToHash(consumed.Random[:]).String())
	tbRandomConsumed.Consumer = strings.ToLower(consumed.Consumer.String())

	tbRandomConsumed.Status = 1
	tbRandomConsumed.IsDeleted = 0
	return tbRandomConsumed
}

func ToTbReveal(vLog *types.Log, reveal *contracts.OracleRevealSeed) *db.TbReveal {
	tbReveal := new(db.TbReveal)
	tbReveal.Id = utils.GetSnowflakeId()
	tbReveal.BlockHash = strings.ToLower(vLog.BlockHash.String())
	tbReveal.TxTime = time.Unix(reveal.Time.Int64(), 0)
	tbReveal.TxHash = strings.ToLower(vLog.TxHash.String())
	tbReveal.BlockId = int64(vLog.BlockNumber)
	tbReveal.Committer = strings.ToLower(reveal.Commiter.String())
	tbReveal.Seed = strings.ToLower(common.BytesToHash(reveal.Seed[:]).String())
	tbReveal.Hash = strings.ToLower(common.BytesToHash(reveal.Hash[:]).String())

	tbReveal.Status = 1
	tbReveal.IsDeleted = 0
	return tbReveal
}
