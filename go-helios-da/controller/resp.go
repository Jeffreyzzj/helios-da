package controller

import (
	"context"
	"go-helios-da/resource"
	"go.uber.org/zap"
)

func BuildOkResponse(ctx context.Context, data interface{}) interface{} {
	return ResponseInfo{
		Data: data,
		Code: 0,
		Msg:  "",
	}
}

func BuildErrResponse(ctx context.Context, data interface{}, err error) interface{} {
	resource.LOGGER.Error("EsBaseSearch has err：", zap.Error(err))
	return ResponseInfo{
		Data: nil,
		Code: 1,
		Msg:  "response has err",
	}
}

type ResponseInfo struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
}
