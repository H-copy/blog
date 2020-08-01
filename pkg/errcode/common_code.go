package errcode

import (
	"net/http"
)

var (

	Success = NewError(0, "成功", http.StatusOK)
	ServerError = NewError(100, "服务内部错误", http.StatusInternalServerError )
	InvalidParams = NewError(101, "参数错误", http.StatusBadRequest )
	NotFound = NewError(102, "未找到", http.StatusNotFound )
	UnauthorizedAuthNotExist = NewError(103, "未找到Appkey", http.StatusInternalServerError )
	unauthorizedTokenError = NewError(104, "Token错误", http.StatusInternalServerError )
	unauthorizedTokenTimeout = NewError(105, "Token超时", http.StatusInternalServerError )
	unauthorizedTokenGenerate = NewError(106, "Token生成失败", http.StatusInternalServerError )
	TooManyRequests = NewError(107, "请求过多", http.StatusInternalServerError )

)