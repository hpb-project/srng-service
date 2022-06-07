package utils

import (
	"encoding/json"
)
//json解析
func JsonUnmarshal(data string) (results interface{},err error){
	err = json.Unmarshal([]byte(data),&results)
	if err != nil{
		return results,err
	}
	return results,nil
}
