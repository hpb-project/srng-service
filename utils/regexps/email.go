package regexps

import (
	"regexp"
)
//邮箱地址验证
func VerifyEmailFormat(emailNum string) bool {
	regular := "^[a-z0-9A-Z\u4e00-\u9fa5]+[- | a-z0-9A-Z . _]+@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)?\\.)+[a-z]{2,}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(emailNum)
}