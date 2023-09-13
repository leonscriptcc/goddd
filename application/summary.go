package application

import "github.com/leonscriptcc/goddd/infrastructure/gconsts"

type HttpResponse struct {
	Code    int64  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func HttpResSuccess(data any) HttpResponse {
	return HttpResponse{
		Code:    gconsts.ResSuccessCode,
		Message: gconsts.ResSuccessMessage,
		Data:    data,
	}
}

func HttpResSystemErr() HttpResponse {
	return HttpResponse{
		Code:    gconsts.ResSysErrCode,
		Message: gconsts.ResSysErrMessage,
		Data:    nil,
	}
}

func HttpResIllegalParam() HttpResponse {
	return HttpResponse{
		Code:    gconsts.ResIllegalParamCode,
		Message: gconsts.ResIllegalParamMessage,
		Data:    nil,
	}
}
