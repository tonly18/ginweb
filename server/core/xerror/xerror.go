package xerror

import (
	"context"
	"errors"
	"fmt"
	"server/core/logger"
)

type TempError struct {
	Err        error
	Code       uint32
	Message    string
	ErrorStack []Error
	Type       int8 //1 error|2 debug
}

func (e *TempError) GetErr() error {
	return e.Err
}

func (e *TempError) SetErr(err error) {
	e.Err = err
}

func (e *TempError) GetCode() uint32 {
	return e.Code
}

func (e *TempError) SetCode(code uint32) {
	e.Code = code
}

func (e *TempError) GetMsg() string {
	return e.Message
}

func (e *TempError) SetMsg(msg string) {
	e.Message = msg
}

func (e *TempError) AddError(err Error) Error {
	if len(e.ErrorStack) == 0 {
		e.ErrorStack = make([]Error, 0, 10)
	}
	e.ErrorStack = append(e.ErrorStack, err)

	//设置Error为当前最新的Error
	e.SetErr(err.GetErr())
	e.SetCode(err.GetCode())
	e.SetMsg(err.GetMsg())
	e.SetType(err.GetType())

	return e
}

func (e *TempError) GetErrorStack() []Error {
	return e.ErrorStack
}

func (e *TempError) GetType() int8 {
	return e.Type
}

func (e *TempError) SetType(itype int8) {
	e.Type = itype
}

func (e *TempError) Error() string {
	return fmt.Sprintf(`code:%v, message:%v`, e.Code, e.Message)
}

func (e *TempError) Copy() Error {
	return &TempError{
		Err:        e.Err,
		Code:       e.Code,
		Message:    e.Message,
		ErrorStack: e.ErrorStack,
		Type:       e.Type,
	}
}

func (e *TempError) Is(err error) bool {
	return errors.Is(e.GetErr(), err)
}

func (e *TempError) Contain(err error) bool {
	for _, v := range e.ErrorStack {
		if errors.Is(v.GetErr(), err) {
			return true
		}
	}
	return false
}

//Wrap 老的错误信息包裹新的错误信息
//
//@params
//	ctx				context.Context	上下文
//	originalError	Error			原始Error
//	newError		Error			新的Error
//@return
//	Error
func Wrap(ctx context.Context, originalError, newError Error) Error {
	if newError == nil {
		panic("the parameter newError cannot be nil")
	}

	//error
	if originalError == nil {
		originalError = newError.Copy()
	}
	originalError.AddError(newError)

	//log
	if newError.GetType() == 1 {
		for _, e := range originalError.GetErrorStack() {
			//fmt.Println("error-list:", e.GetCode(), e.GetErr(), e.GetMsg(), e.GetType())
			//xlog.Errorf(`[%d] message:%v, error:%v`, e.GetCode(), e.GetMsg(), e.GetErr())
			logger.Error(ctx, fmt.Sprintf(`[%d] message:%v, error:%v`, e.GetCode(), e.GetMsg(), e.GetErr()))
		}
	}

	//return
	return originalError
}
