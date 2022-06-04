package controllers

import (
	"fmt"
	"github.com/hpb-project/srng-service/db"
)

type StatController struct {
	BaseController
}

func (e *StatController) ConsumedHistory() {
	result := db.GetHistoryConsumedCount(14)
	if len(result) == 0 {
		e.ResponseInfo(200, nil, nil)
		return
	}

	list := make([]map[string]interface{}, len(result))
	for i, v := range result {
		row := make(map[string]interface{})
		row["y"] = v.Count
		row["dt"] = fmt.Sprintf("%d", v.Date.Unix())
		list[i] = row
	}
	e.ResponseInfo(200, nil, list)
}

func (e *StatController) ConsumedOneDay() {
	result := db.GetHistoryConsumedCount(1)
	if len(result) == 0 {
		e.ResponseInfo(200, nil, nil)
		return
	}
	e.ResponseInfo(200, nil, result[0].Count)
}
