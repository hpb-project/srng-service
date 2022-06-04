package main

import (
	"github.com/astaxie/beego"
	"github.com/hpb-project/srng-service/async"
)

func main() {
	go async.SyncLogs()
	beego.Run()
}
