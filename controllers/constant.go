package controllers

/**
常量获取接口
*/
type ConstantController struct {
	BaseController
}

func (e *ConstantController) GetConstant() {

	e.ResponseInfo(200, nil, nil)
}
