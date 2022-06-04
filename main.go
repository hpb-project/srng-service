package main

import (
	"github.com/hpb-project/srng-service/async"
)

/**
赛马调用合约服务
*/

func main() {
	go async.SyncLogs()
	beego.Run()
}
