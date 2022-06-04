package controllers

import (
	"encoding/json"
	"github.com/hpb-project/srng-service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type HorseInvokeController struct {
	BaseController
}

// 批量增发马匹
func (e *HorseInvokeController) MintHorses() {
	account := beego.AppConfig.String("account")
	horseContract := getContractAddress("horseRaceContract")
	horseContractAbi := horseRaceContract.HorseRaceContractABI
	privateKey := beego.AppConfig.String("privateKey")
	var param models.Param
	data := e.Ctx.Input.RequestBody
	if err := json.Unmarshal(data, &param); err != nil {
		logs.Error(err)
		e.ResponseInfo(500, "参数解析异常", nil)
		return
	}
	if len(param.Ids) != len(param.GeneControl) || len(param.Ids) != len(param.GeneEndurance) || len(param.Ids) != len(param.GeneGrip) {
		e.ResponseInfo(500, "参数个数异常", nil)
		return
	}
	if len(param.Ids) != len(param.GeneSpeed) || len(param.Ids) != len(param.GeneSpeedUp) || len(param.Ids) != len(param.GeneSwerve) {
		e.ResponseInfo(500, "参数个数异常", nil)
		return
	}
	var ids []string
	var geneGrip []string
	var geneSpeedUp []string
	var geneEndurance []string
	var geneSpeed []string
	var geneSwerve []string
	var geneControl []string
	for index, value := range param.Ids {
		ids = append(ids, value)
		geneGrip = append(geneGrip, param.GeneGrip[index])
		geneSpeedUp = append(geneSpeedUp, param.GeneSpeedUp[index])
		geneEndurance = append(geneEndurance, param.GeneEndurance[index])
		geneSpeed = append(geneSpeed, param.GeneSpeed[index])
		geneSwerve = append(geneSwerve, param.GeneSwerve[index])
		geneControl = append(geneControl, param.GeneControl[index])
	}

	result, err := common.SignTx(account, horseContract[0], common.MintHORSERS, horseContractAbi, privateKey, param.To, ids, geneGrip, geneSpeedUp, geneEndurance, geneSpeed, geneSwerve, geneControl)
	if err != nil {
		e.ResponseInfo(500, err.Error(), nil)
		return
	}
	e.ResponseInfo(200, nil, result)
}
