package errcode

import (
	"fmt"
)

type Error struct{
	code int `json:"code"`
	msg string `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}
var statusCodeMap = map[int]int{}

func NewError(code int, msg string, statusCode int) *Error{
	
	if _, ok := codes[code]; ok{
		panic(fmt.Sprintf("错误码： %d 已存在", code))
	}

	
	codes[code] = msg
	statusCodeMap[code] = statusCode	

	return &Error{code: code, msg:msg}
	
}

func (e *Error) Error() string{
	return e.msg
}

func (e *Error) Code() int{
	return e.code
}

func (e *Error) Msg() string{
	return e.msg
}

func (e *Error) Msgf(args[]interface{}) string{
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string{
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error{
	
	e.details = []string{}

	for _, item := range details{
		e.details = append(e.details, item)
	}
	
	return e
}

func (e *Error) StatusCode() int{
	return statusCodeMap[e.code]
}