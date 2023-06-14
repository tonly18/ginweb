package merror

import (
	"context"
	"fmt"
	"server/core/logger"
)

type errorStruct struct {
	Err     error
	Code    int32
	Message string
	Type    int8 //1 error|2 debug
}

//ErrorTemp Struct
type ErrorTemp errorStruct

//myError Struct
type myError errorStruct

//NewError
//
//@params
//	err *ErrorTemp
//		Type	int8	1 error|2 debug
//@return
func NewError(ctx context.Context, err *ErrorTemp) Error {
	if err.Type == 1 {
		logger.Error(ctx, fmt.Sprintf(`[%v] message:%v, error:%v`, err.Code, err.Message, err.Err))
	} else if err.Type == 2 {
		logger.Debug(ctx, fmt.Sprintf(`[%v] message:%v, error:%v`, err.Code, err.Message, err.Err))
	}

	//return
	return &myError{
		Err:     fmt.Errorf(`%w`, err.Err),
		Code:    err.Code,
		Message: err.Message,
	}
}

func (e *myError) GetError() error {
	return e.Err
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

func (e *myError) Error() string {
	return e.Err.Error()
}
