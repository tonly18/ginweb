package merror

import (
	"context"
	"fmt"
	"server/core/logger"
)

type errorStruct struct {
	Error   error
	Code    int32
	Message string
	Type    int8 //1 error|2 debug
}

//ErrorTemp Struct
type ErrorTemp errorStruct

//myError Struct
type myError errorStruct

func NewError(c context.Context, err *ErrorTemp) Error {
	if err.Type == 1 {
		logger.Error(c, fmt.Sprintf(`[%v] message:%v, error:%v`, err.Code, err.Message, err.Error))
	} else if err.Type == 2 {
		logger.Debug(c, fmt.Sprintf(`[%v] message:%v, error:%v`, err.Code, err.Message, err.Error))
	}

	//return
	return &myError{
		Error:   err.Error,
		Code:    err.Code,
		Message: err.Message,
	}
}

func (e *myError) GetError() error {
	return e.Error
}

func (e *myError) GetCode() int32 {
	return e.Code
}

func (e *myError) GetMsg() string {
	return e.Message
}

func (e *myError) GetType() int8 {
	return e.Type
}
