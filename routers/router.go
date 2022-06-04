package routers

import (
	"github.com/astaxie/beego"
	"github.com/hpb-project/srng-service/controllers"
)

func init() {
	beego.Router("/srng/api/v1/consumedhistory", &controllers.StatController{}, "post::ConsumedHistory")
}
