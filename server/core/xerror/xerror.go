package xerror

import (
	"fmt"
)

type NewError struct {
	Err     error
	Code    uint32
	Message string
	stack   []Error
}

func (e *NewError) Error() string {
	return fmt.Sprintf(`code:%v, message:%v`, e.Code, e.Message)
}

func (e *NewError) GetErr() error {
	return e.Err
}

func (e *NewError) SetErr(err error) {
	e.Err = err
}

func (e *NewError) GetCode() uint32 {
	return e.Code
}

func (e *NewError) SetCode(code uint32) {
	e.Code = code
}

func (e *NewError) GetMsg() string {
	return e.Message
}

func (e *NewError) SetMsg(msg string) {
	e.Message = msg
}

func (e *NewError) GetStack() []Error {
	return e.stack
}

func (e *NewError) addStack(err Error) {
	if len(e.stack) == 0 {
		e.stack = make([]Error, 0, 10)
	}
	e.stack = append(e.stack, err)

	//设置err为当前最新的Error
	e.SetErr(err.GetErr())
	e.SetCode(err.GetCode())
	e.SetMsg(err.GetMsg())
}

func (e *NewError) Is(err error) bool {
	return e.GetErr() == err
}

// Wrap 老的错误信息包裹新的错误信息
//
// @params
//
//	originalError	Error			原始Error
//	newErrors		[]Error			新的Error
//
// @return
//
//	Error
func Wrap(originalError, newError Error) Error {
	if originalError == nil {
		panic("the parameter originalError is nil")
	}

	if newError == nil {
		originalError.addStack(&NewError{
			Err:     originalError.GetErr(),
			Code:    originalError.GetCode(),
			Message: originalError.GetMsg(),
		})
	} else {
		originalError.addStack(newError)
	}

	//return
	return originalError
}
