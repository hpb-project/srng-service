package routers

import (
	"github.com/astaxie/beego"
	"github.com/hpb-project/srng-service/controllers"
)

func init() {
//	beego.Router("/srng/api/v1/consumedHistory", &controllers.StatController{}, "post::ConsumedHistory")
//	beego.Router("/srng/api/v1/consumedOneday", &controllers.StatController{}, "post::ConsumedOneDay")
	ns := beego.NewNamespace("/srng",
		beego.NSNamespace("api",
			//用户信息
			beego.NSRouter("/consumedHistory", &controllers.StatController{}, "post:ConsumedHistory"),
			beego.NSRouter("/consumedOneday", &controllers.StatController{}, "post:ConsumedOneDay"),
			beego.NSRouter("/tokeninfo", &controllers.TokenController{}, "post:TokenInfo"),
		),
	)
	beego.AddNamespace(ns)
}
