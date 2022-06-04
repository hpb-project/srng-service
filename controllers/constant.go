package controllers

type StatController struct {
	BaseController
}

func (e *StatController) ConsumedHistory() {

	e.ResponseInfo(200, nil, nil)
}
