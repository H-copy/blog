package errcode

import (
	"testing"
)


func TestNewEror(t *testing.T){
	
	code := 1
	msg := "error 1"
	statusCode := 200
	
	e := NewError(code, msg, statusCode)

	if _, ok := codes[code]; !ok {
		t.Error(`错误类型为添加到缓存队列中`)
	}

	if e.Code() != code {
		t.Error(`code 设置错误`)
	}

	if e.Msg() != msg{
		t.Error(`msg 设置错误`)
	}

	if e.StatusCode() != statusCode{
		t.Error(`statusCode 设置错误`)
	}

	
}