package regexps

import (
	"regexp"
)
//执行密码的正则验证(8~20)
func VerifyPasswordFormat(res,str,end string) bool {
	if len(str) <= 0{
		str = "8"
	}
	if len(end) <= 0{
		str = "20"
	}
	regular := "^[a-zA-Z0-9_\\@\\#\\$\\%\\^\\&\\*\\(\\)\\!\\,\\.\\?\\-\\+\\|\\=]{"+str+","+end+"}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(res)
}
//执行验证码的正则验证
func VerifyCodeFormat(res,str string) bool {
	if len(str) <= 0{
		str = "6"
	}
	regular := "^[0-9]{"+str+"}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(res)
}