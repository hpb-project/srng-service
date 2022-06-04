package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

func getContractAddress(contract string) []string {
	multiAddrs := beego.AppConfig.String(contract)
	addrs := strings.Split(multiAddrs, ",")
	ret := make([]string, 0)
	for _, addr := range addrs {
		if len(addr) > 0 {
			ret = append(ret, addr)
		}
	}
	return ret
}
