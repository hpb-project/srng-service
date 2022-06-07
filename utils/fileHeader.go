package utils

import (
	"bytes"
	"encoding/hex"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"sync"
)

var fileTypeMap sync.Map

func init() {
	fileTypeMap.Store(".jpg", "ffd8ff")  //JPEG (jpg)
	fileTypeMap.Store(".jpeg", "ffd8ff")  //JPEG (jpg)
	fileTypeMap.Store(".png", "89504e470d0a")  //PNG (png)
	fileTypeMap.Store(".gif", "474946383961")  //GIF (gif)
	fileTypeMap.Store(".doc", "d0cf11e0a1b11ae1")  //MS Excel 注意：word、msi 和 excel的文件头一样
	fileTypeMap.Store(".docx", "504b0304")       //docx文件
	fileTypeMap.Store(".pdf", "255044462D312E")  //Adobe Acrobat (pdf)
	fileTypeMap.Store(".mp4", "000000")
	fileTypeMap.Store(".mp3", "494433")
	fileTypeMap.Store(".zip", "504b0304")
	fileTypeMap.Store(".rar", "52617221")
	fileTypeMap.Store(".jar", "0x504B0304")
	fileTypeMap.Store(".gz", "1f8b08")         //gz文件
	fileTypeMap.Store(".xls", "D0CF11E0")
	fileTypeMap.Store(".xlsx", "504b0304")
	fileTypeMap.Store(".ppt", "D0CF11E0A1B11AE1")
	fileTypeMap.Store(".pptx", "504b0304")
}
// 获取前面结果字节的二进制
func bytesToHexString(src []byte) string {
	res := bytes.Buffer{}
	if src == nil || len(src) <= 0 {
		return ""
	}
	temp := make([]byte, 0)
	for _, v := range src {
		sub := v & 0xFF
		hv := hex.EncodeToString(append(temp, sub))
		if len(hv) < 2 {
			res.WriteString(strconv.FormatInt(int64(0), 10))
		}
		res.WriteString(hv)
	}
	return res.String()
}
// 用文件前面几个字节来判断
// fSrc: 文件字节流（就用前面几个字节）
func GetFileType(fSrc []byte,suffix string) bool {
	fileCode 				:= 		bytesToHexString(fSrc)
	checkFileStatus 	    := 		beego.AppConfig.String("checkFile::checkFileStatus")
	fileVale, ok 			:= 		fileTypeMap.Load(suffix)
	if checkFileStatus == ""{
		checkFileStatus = "1"
	}
	if ok {
		keyValue := fileVale.(string)
		if strings.HasPrefix(fileCode, strings.ToLower(keyValue)) ||
			strings.HasPrefix(keyValue, strings.ToLower(fileCode)) {
			return true
		}
		return false
	}else{
		if checkFileStatus == "1"{
			return true
		}else	if checkFileStatus == "2"{
			return false
		}
		return false
	}
}