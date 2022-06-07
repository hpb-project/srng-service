package controllers

import (
	"github.com/hpb-project/srng-service/async"
)
type TokenController struct {
	BaseController
}

func (e *TokenController) TokenInfo() {
	var result = async.GetTokenInfo()
	e.ResponseInfo(200, nil, result)
}
