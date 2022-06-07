package utils

import (
	"encoding/hex"
	"strings"
)

func FormatAddress(str string, length int) string {
	desStr := str
	tmp := "00000000000000000000000000000000000000000000000000000000000000000000000000000000"
	if len(desStr) >= 2 {
		if desStr[:2] == "0x" {
			desStr = desStr[2:]
		}
	}
	if len(desStr) > length {
		desStr = desStr[len(desStr)-length:]
	} else {
		desStr = tmp[:length-len(desStr)] + desStr
	}
	return "0x" + desStr
}

func FormatValue(str string) string {
	resStr := str
	desStr := "0"
	if len(resStr) >= 2 {
		if resStr[:2] == "0x" {
			resStr = resStr[2:]
		}
	}
	for i := 0; i < len(resStr); i++ {
		if resStr[i:i+1] != "0" {
			desStr = resStr[i:]
			break
		}
	}
	return "0x" + desStr
}

func FormatString(str string) string {
	desStr := str
	// 去除null
	null, _ := hex.DecodeString("00")
	desStr = strings.Replace(desStr, string(null), "", -1)
	// 去除\n
	enter, _ := hex.DecodeString("0A")
	desStr = strings.Replace(desStr, string(enter), "", -1)
	//去除\f
	newPage, _ := hex.DecodeString("0C")
	desStr = strings.Replace(desStr, string(newPage), "", 1)
	// 去除第一个空格
	space, _ := hex.DecodeString("20")
	desStr = strings.Replace(desStr, string(space), "", 1)
	// 去除第一个制表符
	horizontalTab, _ := hex.DecodeString("09")
	desStr = strings.Replace(desStr, string(horizontalTab), "", 1)
	// 去除\a
	bel, _ := hex.DecodeString("07")
	desStr = strings.Replace(desStr, string(bel), "", -1)
	// 去除第一个\b
	backspace, _ := hex.DecodeString("08")
	desStr = strings.Replace(desStr, string(backspace), "", 1)
	// 去除end of text
	endOfText, _ := hex.DecodeString("03")
	desStr = strings.Replace(desStr, string(endOfText), "", -1)
	// 去除vertical tab
	vt, _ := hex.DecodeString("0B")
	desStr = strings.Replace(desStr, string(vt), "", -1)
	// 去除cancel
	cancel, _ := hex.DecodeString("18")
	desStr = strings.Replace(desStr, string(cancel), "", -1)

	return desStr
}

func MarshalText(b []byte) []byte {
	result := make([]byte, len(b)*2+2)
	copy(result, `0x`)
	hex.Encode(result[2:], b)
	return result
}
