package utils

import (
	"github.com/astaxie/beego"
)

func GetFileUrl(fileurl string) string {
	domain := beego.AppConfig.String("domain")
	return domain + fileurl
}
