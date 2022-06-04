package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"math/big"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/globalsign/mgo/bson"
)

type OperaController struct {
	BaseController
}

// 开始比赛
func (e *OperaController) StartGame() {
	account := beego.AppConfig.String("account")
	arenaExtra3Contract := getContractAddress("arenaExtra3Contract")
	addr := ethcommon.HexToAddress(arenaExtra3Contract[0])
	//arenaExtra3ContractAbi := horseArenaExtra3.HorseArenaExtra3ABI
	privateKey := beego.AppConfig.String("privateKey")
	var param models.StartGameParam
	data := e.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &param); err != nil {
		logs.Error(err)
		e.ResponseInfo(500, "参数解析异常", nil)
		return
	}
	if param.TokenId == "" || param.GameId == "" || param.Distance == "" || param.Level == "" || param.Types == "" ||
		len(param.HorseIds) != 12 {
		e.ResponseInfo(500, "参数异常,缺少参数", nil)
		return
	}
	{
		client := common.EthClient()
		nonce := common.GetCount()
		gasPrice, _ := client.SuggestGasPrice(context.Background())
		defaultOpt := &bind.TransactOpts{
			From:     ethcommon.HexToAddress(account),
			Signer:   getSignFunc(privateKey),
			Context:  context.Background(),
			GasLimit: 10000000,
			GasPrice: gasPrice,
			Nonce:    big.NewInt(int64(nonce)),
		}
		contract, err := horseArenaExtra3.NewHorseArenaExtra3(addr, client)
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		tokenid, _ := new(big.Int).SetString(param.TokenId, 10)
		gameid, _ := new(big.Int).SetString(param.GameId, 10)
		typeid, _ := new(big.Int).SetString(param.Types, 10)
		level, _ := new(big.Int).SetString(param.Level, 10)
		distance, _ := new(big.Int).SetString(param.Distance, 10)
		horseids := make([]*big.Int, 0)
		for _, id := range param.HorseIds {
			h, _ := new(big.Int).SetString(id, 10)
			horseids = append(horseids, h)
		}

		tx, err := contract.StartGame(defaultOpt, tokenid, gameid, horseids, typeid, level, distance)
		if err != nil {
			logs.Error("start game failed", "err", err.Error())
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		e.ResponseInfo(200, nil, tx.Hash())
	}
}

// 取消比赛
func (e *OperaController) CancelGame() {
	account := beego.AppConfig.String("account")
	arenaExtra3Contract := getContractAddress("arenaExtra3Contract")
	addr := ethcommon.HexToAddress(arenaExtra3Contract[0])
	//arenaExtra3ContractAbi := horseArenaExtra3.HorseArenaExtra3ABI
	privateKey := beego.AppConfig.String("privateKey")
	var param models.CancelGameParam
	data := e.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &param); err != nil {
		logs.Error(err)
		e.ResponseInfo(500, "参数解析异常", nil)
		return
	}
	if param.TokenId == "" || len(param.GameIds) == 0 {
		e.ResponseInfo(500, "参数异常,缺少参数", nil)
		return
	}
	{
		client := common.EthClient()
		nonce := common.GetCount()
		gasPrice, _ := client.SuggestGasPrice(context.Background())
		defaultOpt := &bind.TransactOpts{
			From:     ethcommon.HexToAddress(account),
			Signer:   getSignFunc(privateKey),
			Context:  context.Background(),
			GasLimit: 10000000,
			GasPrice: gasPrice,
			Nonce:    big.NewInt(int64(nonce)),
		}
		contract, err := horseArenaExtra3.NewHorseArenaExtra3(addr, client)
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		tokenid, _ := new(big.Int).SetString(param.TokenId, 10)
		gameids := make([]*big.Int, 0)
		for _, id := range param.GameIds {
			h, _ := new(big.Int).SetString(id, 10)
			gameids = append(gameids, h)
		}

		tx, err := contract.CancelGame(defaultOpt, tokenid, gameids)
		if err != nil {
			logs.Error("cancel game failed", "err", err.Error())
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		e.ResponseInfo(200, nil, tx.Hash())
	}
}

// 结束积分赛
func (e *OperaController) EndPointGame() {
	account := beego.AppConfig.String("account")
	arenaExtraContract := getContractAddress("arenaExtraContract")
	addr := ethcommon.HexToAddress(arenaExtraContract[0])
	//arenaExtraContractAbi := horseArenaExtra.HorseArenaExtraABI
	privateKey := beego.AppConfig.String("privateKey")
	var param models.EndGameParam
	data := e.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &param); err != nil {
		logs.Error(err)
		e.ResponseInfo(500, "参数解析异常", nil)
		return
	}
	if param.TokenId == "" || param.GameId == "" || len(param.Rank) != 12 || len(param.Comp) != 12 || len(param.Score) != 12 {
		e.ResponseInfo(500, "参数异常,缺少参数", nil)
		return
	}
	{
		client := common.EthClient()
		nonce := common.GetCount()
		gasPrice, _ := client.SuggestGasPrice(context.Background())
		defaultOpt := &bind.TransactOpts{
			From:     ethcommon.HexToAddress(account),
			Signer:   getSignFunc(privateKey),
			Context:  context.Background(),
			GasLimit: 10000000,
			GasPrice: gasPrice,
			Nonce:    big.NewInt(int64(nonce)),
		}
		contract, err := horseArenaExtra.NewHorseArenaExtra(addr, client)
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		tokenid, _ := new(big.Int).SetString(param.TokenId, 10)
		gameid, _ := new(big.Int).SetString(param.GameId, 10)
		ranks := make([]*big.Int, len(param.Rank))
		scores := make([]*big.Int, len(param.Score))
		comps := make([]*big.Int, len(param.Comp))
		for i := 0; i < len(param.Rank); i++ {
			ranks[i], _ = new(big.Int).SetString(param.Rank[i], 10)
			scores[i], _ = new(big.Int).SetString(param.Score[i], 10)
			comps[i], _ = new(big.Int).SetString(param.Comp[i], 10)
		}
		//tokenId *big.Int, gameId *big.Int, rank []*big.Int, score []*big.Int, comp []*big.Int
		tx, err := contract.EndPointGame(defaultOpt, tokenid, gameid, ranks, scores, comps)
		if err != nil {
			logs.Error("EndPointGame game failed", "err", err.Error())
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		e.ResponseInfo(200, nil, tx.Hash())
	}
}

// 结束大奖赛
func (e *OperaController) EndGrandGame() {
	account := beego.AppConfig.String("account")
	arenaExtraContract := getContractAddress("arenaExtraContract")
	addr := ethcommon.HexToAddress(arenaExtraContract[0])
	//arenaExtraContractAbi := horseArenaExtra.HorseArenaExtraABI
	privateKey := beego.AppConfig.String("privateKey")
	var param models.EndGameParam
	data := e.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &param); err != nil {
		logs.Error(err)
		e.ResponseInfo(500, "参数解析异常", nil)
		return
	}
	if param.TokenId == "" || param.GameId == "" || len(param.Rank) != 12 || len(param.Comp) != 12 || len(param.Score) != 12 {
		e.ResponseInfo(500, "参数异常,缺少参数", nil)
		return
	}
	{
		client := common.EthClient()
		nonce := common.GetCount()
		gasPrice, _ := client.SuggestGasPrice(context.Background())
		defaultOpt := &bind.TransactOpts{
			From:     ethcommon.HexToAddress(account),
			Signer:   getSignFunc(privateKey),
			Context:  context.Background(),
			GasLimit: 10000000,
			GasPrice: gasPrice,
			Nonce:    big.NewInt(int64(nonce)),
		}
		contract, err := horseArenaExtra.NewHorseArenaExtra(addr, client)
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		tokenid, _ := new(big.Int).SetString(param.TokenId, 10)
		gameid, _ := new(big.Int).SetString(param.GameId, 10)
		ranks := make([]*big.Int, len(param.Rank))
		scores := make([]*big.Int, len(param.Score))
		comps := make([]*big.Int, len(param.Comp))
		for i := 0; i < len(param.Rank); i++ {
			ranks[i], _ = new(big.Int).SetString(param.Rank[i], 10)
			scores[i], _ = new(big.Int).SetString(param.Score[i], 10)
			comps[i], _ = new(big.Int).SetString(param.Comp[i], 10)
		}
		//tokenId *big.Int, gameId *big.Int, rank []*big.Int, score []*big.Int, comp []*big.Int
		tx, err := contract.EndGrandGame(defaultOpt, tokenid, gameid, ranks, scores, comps)
		if err != nil {
			logs.Error("EndPointGame game failed", "err", err.Error())
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		e.ResponseInfo(200, nil, tx.Hash())
	}
}

// 增发装备
func (e *OperaController) MintEquip() {
	account := beego.AppConfig.String("account")
	contract := getContractAddress("horseEquipContract")
	horseEquipContractAbi := horseEquipContract.HorseEquipContractABI
	privateKey := beego.AppConfig.String("privateKey")
	var param models.MintEquipParam
	data := e.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &param); err != nil {
		logs.Error(err)
		e.ResponseInfo(500, "参数解析异常", nil)
		return
	}
	if param.To == "" || param.EquipStyle == "" || param.EquipTypes == "" {
		e.ResponseInfo(500, "参数异常,缺少参数", nil)
		return
	}
	result, err := common.SignTx(account, contract[0], common.MintEquipFun, horseEquipContractAbi, privateKey, param.To,
		param.EquipTypes, param.EquipStyle)
	if err != nil {
		e.ResponseInfo(500, err.Error(), nil)
		return
	}
	e.ResponseInfo(200, nil, result)
}

// 周排名上传
func (e *OperaController) SetWeekRank() {
	account := beego.AppConfig.String("account")
	arenaExtra2Contract := getContractAddress("arenaExtra2Contract")
	arenaExtra2ContractAbi := horseArenaExtra2.HorseArenaExtra2ABI
	privateKey := beego.AppConfig.String("privateKey")
	var param models.Rank
	data := e.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &param); err != nil {
		logs.Error(err)
		e.ResponseInfo(500, "参数解析异常", nil)
		return
	}
	if len(param.Rank) == len(param.HorseIds) && len(param.Rank) == len(param.TotalScores) &&
		len(param.Rank) == len(param.Accounts) && len(param.Rank) > 0 {
		result, err := common.SignTx(account, arenaExtra2Contract[0], common.WeekRank, arenaExtra2ContractAbi, privateKey, param.HorseIds,
			param.TotalScores, param.Accounts, param.Rank)
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		e.ResponseInfo(200, nil, result)
	} else {
		e.ResponseInfo(500, "参数异常", nil)
		return
	}
}

// 月排名上传
func (e *OperaController) SetMonthRank() {
	account := beego.AppConfig.String("account")
	arenaExtra2Contract := getContractAddress("arenaExtra2Contract")
	arenaExtra2ContractAbi := horseArenaExtra2.HorseArenaExtra2ABI
	privateKey := beego.AppConfig.String("privateKey")
	var param models.Rank
	data := e.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &param); err != nil {
		logs.Error(err)
		e.ResponseInfo(500, "参数解析异常", nil)
		return
	}
	if len(param.Rank) == len(param.HorseIds) && len(param.Rank) == len(param.TotalScores) &&
		len(param.Rank) == len(param.Accounts) && len(param.Rank) > 0 {
		result, err := common.SignTx(account, arenaExtra2Contract[0], common.MonthRank, arenaExtra2ContractAbi, privateKey, param.HorseIds,
			param.TotalScores, param.Accounts, param.Rank)
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		e.ResponseInfo(200, nil, result)
		return
	} else {
		e.ResponseInfo(500, "参数异常", nil)
		return
	}
}

// 年排名上传
func (e *OperaController) SetYearsRank() {
	account := beego.AppConfig.String("account")
	arenaExtra2Contract := getContractAddress("arenaExtra2Contract")
	arenaExtra2ContractAbi := horseArenaExtra2.HorseArenaExtra2ABI
	privateKey := beego.AppConfig.String("privateKey")
	var param models.Rank
	data := e.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &param); err != nil {
		logs.Error(err)
		e.ResponseInfo(500, "参数解析异常", nil)
		return
	}
	if len(param.Rank) == len(param.HorseIds) && len(param.Rank) == len(param.TotalScores) &&
		len(param.Rank) == len(param.Accounts) && len(param.Rank) > 0 {
		result, err := common.SignTx(account, arenaExtra2Contract[0], common.YearsRank, arenaExtra2ContractAbi, privateKey, param.HorseIds,
			param.TotalScores, param.Accounts, param.Rank)
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		e.ResponseInfo(200, nil, result)
		return
	} else {
		e.ResponseInfo(500, "参数异常", nil)
		return
	}
}

// 查询接口
func (e *OperaController) GetEvent() {
	pageNo := e.Ctx.Input.Param(":pageNo")
	pageSize := e.Ctx.Input.Param(":pageSize")
	if pageNo == "" || pageSize == "" {
		e.ResponseInfo(500, "参数异常", nil)
		return
	}
	pn, _ := strconv.Atoi(pageNo)
	ps, _ := strconv.Atoi(pageSize)
	if ps > 50 {
		e.ResponseInfo(500, "No more than 50", nil)
		return
	}
	strconv.Atoi(pageSize)
	var result []interface{}
	models.FindPage(pn, ps, &result)
	e.ResponseInfo(200, nil, result)
}

func (e *OperaController) GetEventByTxHash() {
	hash := e.Ctx.Input.Param(":hash")
	var result []interface{}
	err := models.FindAllByHash(bson.M{"tx_hash": hash}, &result)
	if err != nil {
		logs.Error(err)
		e.ResponseInfo(500, err.Error(), nil)
		return
	}
	e.ResponseInfo(200, nil, result)
}

func (e *OperaController) GetEventByEventName() {
	name := e.Ctx.Input.Param(":name")
	var result []interface{}
	err := models.FindAllByHash(bson.M{"event_name": name}, &result)
	if err != nil {
		logs.Error(err)
		e.ResponseInfo(500, err.Error(), nil)
		return
	}
	e.ResponseInfo(200, nil, result)
}

//签名函数
func (e *OperaController) Sign() {
	name := e.Ctx.Input.Param(":name")
	result, err := common.Sign(name)
	if err != nil {
		logs.Error(err)
		e.ResponseInfo(500, err.Error(), nil)
		return
	}
	e.ResponseInfo(200, nil, result)
}

// 验证登录token
func (e *OperaController) CheckToken() {
	address := e.Ctx.Input.Param(":address")
	token := e.Ctx.Input.Param(":token")
	if address == "" || token == "" {
		e.ResponseInfo(500, "Parameters of the abnormal", nil)
		return
	}
	loginContract := beego.AppConfig.String("loginContract")
	loginContractAbi := userLogin.UserLoginABI
	result, err := common.CallContractOpera(loginContractAbi, common.CheckToken, loginContract, address, token)
	if err != nil {
		e.ResponseInfo(500, err.Error(), nil)
		return
	}
	e.ResponseInfo(200, nil, result[0])
}

// 初始化马匹等级
func (e *OperaController) InitHorseGrade() {
	account := beego.AppConfig.String("account")
	horseRaceExtra1Contract := getContractAddress("horseRaceExtra1Contract")
	horseRaceExtra1ContractAbi := horseRaceExtra1.HorseRaceExtra1ABI
	privateKey := beego.AppConfig.String("privateKey")
	var param models.InitGrade
	data := e.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &param); err != nil {
		logs.Error(err)
		e.ResponseInfo(500, "参数解析异常", nil)
		return
	}
	if len(param.HorseIds) != len(param.Comp) {
		e.ResponseInfo(500, "参数异常", nil)
		return
	}
	result, err := common.SignTx(account, horseRaceExtra1Contract[0], common.InitHorseGradeFun, horseRaceExtra1ContractAbi, privateKey, param.HorseIds,
		param.Comp)
	if err != nil {
		e.ResponseInfo(500, err.Error(), nil)
		return
	}
	e.ResponseInfo(200, nil, result)
}

// 开始比赛
func (e *OperaController) VerifySign() {
	var param models.VerifySignParam
	data := e.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &param); err != nil {
		logs.Error(err)
		e.ResponseInfo(500, "参数解析异常", nil)
		return
	}
	logs.Info("get param ", param)
	if param.Address == "" || param.Message == "" || param.Signature == "" {
		e.ResponseInfo(500, "参数异常,缺少参数", nil)
		return
	}
	sig := ethcommon.FromHex(param.Signature)
	if len(sig) != 65 {
		e.ResponseInfo(500, "Signature 参数异常", nil)
		return
	}
	userAddr := ethcommon.HexToAddress(param.Address)
	logs.Debug("signature verify with recover1")
	recover1 := common.Sender(ethcommon.FromHex(param.Message), sig)
	if bytes.Compare(userAddr.Bytes(), recover1.Bytes()) == 0 {
		logs.Debug("signature verify passed with recover1.")
		e.ResponseInfo(200, nil, true)
	} else {
		logs.Debug("signature verify with recover2")
		sig2 := ethcommon.FromHex(param.Signature)
		recover2 := common.Sender([]byte(param.Message), sig2)
		if bytes.Compare(userAddr.Bytes(), recover2.Bytes()) == 0 {
			logs.Debug("signature verify passed with recover2.")
			e.ResponseInfo(200, nil, true)
		}
	}

	e.ResponseInfo(200, nil, false)
}

// 验证登录token
func (e *OperaController) CheckTx() {
	txhash := e.Ctx.Input.Param(":hash")
	if txhash == "" {
		e.ResponseInfo(500, "Parameters of the abnormal", nil)
		return
	}
	result := make(map[string]int)
	result["status"] = 2 // pending
	client := common.EthClient()
	h := ethcommon.HexToHash(txhash)
	r, err := client.TransactionReceipt(context.Background(), h)
	if err == nil {
		result["status"] = int(r.Status) // 0 : failed  1 : succeed
	}
	e.ResponseInfo(200, nil, result)
}
