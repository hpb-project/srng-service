package utils

import (
	"strings"
	"encoding/base64"
)
//字符串部分截取的操作
func SubString(str string, begin, length int) string {
	rs := []rune(str)
	lth := len(rs)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	return string(rs[begin:end])
}
//去除空格和特殊字符
func StringReplace(str string) string{
	//去除空格
	str = strings.Replace(str, " ", "", -1)
	//去除换行符
	str = strings.Replace(str, "\n", "", -1)
	return str
}
//去除特殊字符`
func StringReplaces(str string) string{
	//去除换行符
	str = strings.Replace(str, "`", "", -1)
	return str
}
//字符串转换字节
func Base64ToByte(str string) ([]byte,error){
	resBytes,err := base64.StdEncoding.DecodeString(str)
	if err != nil{
		return resBytes,err
	}
	return resBytes,nil
}
//字节转字符串
func ByteToBase64(str []byte) (res string){
	res	 	= base64.StdEncoding.EncodeToString(str)
	return res
}