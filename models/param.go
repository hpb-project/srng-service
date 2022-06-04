package models

import (
	"encoding/json"
)

func Marshal(param interface{}) []byte {
	byteInfo, _ := json.Marshal(param)
	return byteInfo
}
