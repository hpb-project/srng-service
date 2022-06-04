package routers

import (
	"github.com/astaxie/beego"
	"github.com/hpb-project/srng-service/controllers"
)

func init() {

	/**
	马匹装备路由
	*/
	beego.Router("/race/api/v1/mintEquipment", &controllers.EquipmentController{}, "post:MintEquipment")

	/**
	马匹信息路由
	*/
	beego.Router("/race/api/v1/mintHorses", &controllers.HorseInvokeController{}, "post:MintHorses")

	/**
	常量值查询
	*/
	beego.Router("/race/api/v1/getConstant", &controllers.ConstantController{}, "get:GetConstant")

	/**
	游戏中心服操作
	*/
	beego.Router("/race/api/v1/startGame", &controllers.OperaController{}, "post:StartGame")
	beego.Router("/race/api/v1/cancelGame", &controllers.OperaController{}, "post:CancelGame")
	beego.Router("/race/api/v1/endPointGame", &controllers.OperaController{}, "post:EndPointGame")
	beego.Router("/race/api/v1/endGrandGame", &controllers.OperaController{}, "post:EndGrandGame")
	beego.Router("/race/api/v1/mintEquip", &controllers.OperaController{}, "post:MintEquip")

	beego.Router("/race/api/v1/getEvent/?:pageNo/?:pageSize", &controllers.OperaController{}, "get:GetEvent")
	beego.Router("/race/api/v1/getEventByTxHash/?:hash", &controllers.OperaController{}, "get:GetEventByTxHash")
	beego.Router("/race/api/v1/getEventByEventName/?:name", &controllers.OperaController{}, "get:GetEventByEventName")
	beego.Router("/race/api/v1/sign/?:name", &controllers.OperaController{}, "get:Sign")
	beego.Router("/race/api/v1/checkToken/?:address/?:token", &controllers.OperaController{}, "get:CheckToken")
	beego.Router("/race/api/v1/verifySign", &controllers.OperaController{}, "post:VerifySign")
	beego.Router("/race/api/v1/checkTx/?:hash", &controllers.OperaController{}, "get:CheckTx")

	beego.Router("/race/api/v1/setWeekRank", &controllers.OperaController{}, "post:SetWeekRank")
	beego.Router("/race/api/v1/setMonthRank", &controllers.OperaController{}, "post:SetMonthRank")
	beego.Router("/race/api/v1/setYearsRank", &controllers.OperaController{}, "post:SetYearsRank")
	beego.Router("/race/api/v1/initHorseGrade", &controllers.OperaController{}, "post:InitHorseGrade")
	beego.Router("/race/api/v1/setDate", &controllers.ConstantController{}, "post:SetDate")
}
