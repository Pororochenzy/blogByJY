package errcode

import (
	"fmt"
	"net/http"
)

//错误处理，编写一些常用的错误处理公共方法，用以标准化我们的错误输出
type Error struct {
	code     int      `json:"code"`
	msg      string   `json:"msg"`
	detatils []string `json:"detatils"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码%d已经存在，请更换一个,code"))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码:%d,错误信息：%s", e.code, e.msg)
}
func (e *Error) Code() int {
	return e.code
}
func (e *Error) Msg() string {
	return e.msg
}
func (e *Error) Msgf(args []interface{}) string { // ?如何用
	return fmt.Sprintf(e.msg, args...)
}
func (e *Error) Details() []string {
	return e.detatils
}
func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.detatils = []string{}
	for _, d := range details {
		newError.detatils = append(newError.detatils, d)
	}
	return &newError
}

//主要针对一些特定错误码进行状态码转换
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.code:
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		fallthrough
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
