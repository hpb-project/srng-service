package controllers

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

/**
常量获取接口
*/
type ConstantController struct {
	BaseController
}

func (e *ConstantController) GetConstant() {
	var res = &models.Constant{}
	constantContractAddr := getContractAddress("constantContractAddr")
	addr := ethcommon.HexToAddress(constantContractAddr[0])
	defaultOpt := &bind.CallOpts{}
	client := common.EthClient()

	contract, err := constant.NewConstantCaller(addr, client)
	if err != nil {
		e.ResponseInfo(500, err.Error(), nil)
		return
	}
	{
		v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, err := contract.GetConstant(defaultOpt)
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		result := []*big.Int{v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12}
		formatConstant(result, res)
	}
	{
		v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, err := contract.GetConstant2(defaultOpt)
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		result := []*big.Int{v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12}
		formatConstant2(result, res)
	}
	{
		v0, v1, v2, v3, v4, err := contract.GetConstant3(defaultOpt)
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		result := []*big.Int{v0, v1, v2, v3, v4}
		formatConstant3(result, res)
	}
	{
		v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, err := contract.GetConstant4(defaultOpt)
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		result := []*big.Int{v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11}
		formatConstant4(result, res)
	}
	{
		v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, err := contract.GetConstant5(defaultOpt)
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		result := []*big.Int{v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11}
		formatConstant5(result, res)
	}
	{
		v0, v1, v2, v3, v4, v5, v6, v7, err := contract.GetConstant6(defaultOpt)
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		result := []*big.Int{v0, v1, v2, v3, v4, v5, v6, v7}
		formatConstant6(result, res)
	}
	{
		fee_1200, err := contract.GetApplyGameAmountByDistance(defaultOpt, big.NewInt(1200))
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		fee_2400, err := contract.GetApplyGameAmountByDistance(defaultOpt, big.NewInt(2400))
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		fee_3600, err := contract.GetApplyGameAmountByDistance(defaultOpt, big.NewInt(3600))
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		result := make(map[string]int64)
		result["1200"] = fee_1200.Int64()
		result["2400"] = fee_2400.Int64()
		result["3600"] = fee_3600.Int64()
		res.ApplyGameFee = result
	}
	{
		award_1200, err := contract.GetAwardAmountByDistance(defaultOpt, big.NewInt(1200))
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		award_2400, err := contract.GetAwardAmountByDistance(defaultOpt, big.NewInt(2400))
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		award_3600, err := contract.GetAwardAmountByDistance(defaultOpt, big.NewInt(3600))
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		result := make(map[string]int64)
		result["1200"] = award_1200.Int64()
		result["2400"] = award_2400.Int64()
		result["3600"] = award_3600.Int64()
		res.AllAwardAmount = result
	}

	e.ResponseInfo(200, nil, res)
}

func getSignFunc(privk string) bind.SignerFn {
	privateKey, err := crypto.HexToECDSA(privk[2:])
	if err != nil {
		beego.Error("hex to ecdsa failed, err :", err)
		return nil
	}
	chainId, err := common.EthClient().ChainID(context.Background())
	if err != nil {
		beego.Error("get chainId failed, err :", err)
		return nil
	}
	return func(addr ethcommon.Address, tx *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	}
}

// 向系统设置当前时间
func (e *ConstantController) SetDate() {
	account := beego.AppConfig.String("account")
	privateKey := beego.AppConfig.String("privateKey")

	var param models.ParamDate
	constantContractAddr := getContractAddress("constantContractAddr")
	addr := ethcommon.HexToAddress(constantContractAddr[0])
	client := common.EthClient()

	gasPrice, _ := client.SuggestGasPrice(context.Background())
	nonce := common.GetCount()
	defaultOpt := &bind.TransactOpts{
		From:     ethcommon.HexToAddress(account),
		Signer:   getSignFunc(privateKey),
		Context:  context.Background(),
		GasLimit: 1000000,
		GasPrice: gasPrice,
		Nonce:    big.NewInt(int64(nonce)),
	}

	data := e.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &param); err != nil {
		logs.Error(err)
		e.ResponseInfo(500, "参数解析异常", nil)
		return
	}
	if true {
		contract, err := constant.NewConstantTransactor(addr, client)
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		tx, err := contract.SetDate(defaultOpt, uint32(param.Year), uint32(param.Month), uint32(param.Week))
		if err != nil {
			e.ResponseInfo(500, err.Error(), nil)
			return
		}
		e.ResponseInfo(200, nil, tx.Hash())
	} else {
		e.ResponseInfo(500, "参数异常", nil)
		return
	}
}

func formatConstant(result []*big.Int, constant *models.Constant) {
	constant.BuyTaxRate = result[0].String()
	constant.ModifyNameTimes = result[1].String()
	constant.TakeBackFromMarketCD = result[2].String()
	constant.PutMarketCD = result[3].String()
	constant.GrowUpTime = result[4].String()
	constant.TakeBackFromStudFarmCD = result[5].String()
	constant.PutStudFarmCD = result[6].String()
	constant.MareBreedCD = result[7].String()
	constant.StudBreedCD = result[8].String()
	constant.BreedCostType1 = result[9].String()
	constant.BreedCostValue1 = result[10].String()
	constant.BreedCostType2 = result[11].String()
	constant.BreedCostValue2 = result[12].String()
}
func formatConstant2(result []*big.Int, constant *models.Constant) {
	constant.TrainingValueMin = result[0].String()
	constant.TrainingValueMax = result[1].String()
	constant.TrainingAddValue = result[2].String()
	constant.TrainingCD = result[3].String()
	constant.TrainingReduceCD = result[4].String()
	constant.TrainingReduceValue = result[5].String()
	constant.TrainingCostType = result[6].String()
	constant.TrainingCostValue = result[7].String()
	constant.RacecourseDayLimit = result[8].String()
	constant.ApplyRaceCD = result[9].String()
	constant.CancelRaceCD = result[10].String()
	constant.RaceCostType = result[11].String()
	constant.RaceCostValue = result[12].String()
}
func formatConstant3(result []*big.Int, constant *models.Constant) {
	constant.BreedTaxRate = result[0].String()
	constant.OwnerGainRate = result[1].String()
	constant.EnergyCost = result[2].String()
	constant.EnergyRecoverCD = result[3].String()
	constant.RacecourseCloseCD = result[4].String()
}

func formatConstant4(result []*big.Int, constant *models.Constant) {
	constant.WeekRankCount = result[0].String()
	constant.MonthRankCount = result[1].String()
	constant.YearRankCount = result[2].String()
	constant.MaxEnergy = result[3].String()
	constant.InitScore = result[4].String()
	constant.InitGrade = result[5].String()
	constant.MaxApplyCount = result[6].String()
	constant.TrackNumber = result[7].String()
	constant.BreCoefficient = result[8].String()
	constant.MinDiscountOfHorse = result[9].String()
	constant.MaxRewardOfHorse = result[10].String()
	constant.MortAmount = result[11].String()
}

func formatConstant5(result []*big.Int, constant *models.Constant) {
	constant.MortToken = result[0].String()
	constant.MaxSpacing = result[1].String()
	constant.DistApplyFee = result[2].String()
	constant.AwardToken = result[3].String()
	constant.AwardAmount = result[4].String() // deprecated
	constant.ExtraAwardToken = result[5].String()
	constant.ExtraAwardAmount = result[6].String()
	constant.MaxScore = result[7].String()
	constant.FeeRateOfEquip = result[8].String()
	constant.MinDiscountOfEquip = result[9].String()
	constant.MaxRewardOfEquip = result[10].String()
	constant.MinSpacing = result[11].String()
}
func formatConstant6(result []*big.Int, constant *models.Constant) {
	constant.Resting = result[0].String()
	constant.OnSale = result[1].String()
	constant.Breeding = result[2].String()
	constant.Signing = result[3].String()
	constant.InGame = result[4].String()
	constant.OnHold = result[5].String()
	constant.OnSell = result[6].String()
	constant.Equipped = result[7].String()
}
