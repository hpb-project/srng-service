package routers

import (
	"github.com/astaxie/beego"
	"github.com/hpb-project/srng-service/controllers"
)

func init() {
	beego.Router("/srng/api/v1/consumedHistory", &controllers.StatController{}, "post::ConsumedHistory")
	beego.Router("/srng/api/v1/consumedOneday", &controllers.StatController{}, "post::ConsumedOneDay")
}
