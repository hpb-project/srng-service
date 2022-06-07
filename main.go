package main

import (
	"github.com/astaxie/beego"
	"github.com/hpb-project/srng-service/async"
	_ "github.com/hpb-project/srng-service/routers"
)

func main() {
	go async.SyncLogs()
	beego.Run()
}
