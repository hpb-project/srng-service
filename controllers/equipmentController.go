package controllers

import (
	"encoding/json"
	"math/big"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

/**
马匹装备接口
*/
type EquipmentController struct {
	BaseController
}

// 增发装备
func (e *EquipmentController) MintEquipment() {
	account := beego.AppConfig.String("account")
	equipContract := getContractAddress("horseEquipContract")
	equipContractAbi := horseEquipContract.HorseEquipContractABI
	privateKey := beego.AppConfig.String("privateKey")
	var param models.Param
	data := e.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &param); err != nil {
		logs.Error(err)
		e.ResponseInfo(500, "参数解析异常", nil)
		return
	}
	if len(param.Ids) != len(param.Prices) || len(param.Ids) != len(param.Styles) || len(param.Ids) != len(param.Types) {
		e.ResponseInfo(500, "参数个数异常", nil)
		return
	}
	var ids []string
	var types []*big.Int
	var styles []*big.Int
	var prices []*big.Int
	for index, value := range param.Ids {
		ids = append(ids, value)
		types = append(types, big.NewInt(int64(param.Types[index])))
		styles = append(styles, big.NewInt(int64(param.Styles[index])))
		prices = append(prices, common.FormatPrice(param.Prices[index]))
	}

	result, err := common.SignTx(account, equipContract[0], common.MintEQUIP, equipContractAbi, privateKey, param.To, ids, types, styles, prices)
	if err != nil {
		e.ResponseInfo(500, err.Error(), nil)
		return
	}
	e.ResponseInfo(200, nil, result)
}
