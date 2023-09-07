package xerror

import (
	"errors"
	"fmt"
)

type NewError struct {
	Err     error
	Code    uint32
	Message string
	error   []Error
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

func (e *NewError) addError(err Error) Error {
	if len(e.error) == 0 {
		e.error = make([]Error, 0, 10)
	}
	e.error = append(e.error, err)

	//设置err为当前最新的Error
	e.SetErr(err.GetErr())
	e.SetCode(err.GetCode())
	e.SetMsg(err.GetMsg())

	return e
}

func (e *NewError) GetError() []Error {
	return e.error
}

func (e *NewError) Copy() Error {
	return &NewError{
		Err:     e.Err,
		Code:    e.Code,
		Message: e.Message,
		error:   e.error,
	}
}

func (e *NewError) Is(err error) bool {
	return errors.Is(e.GetErr(), err)
}

func (e *NewError) Contain(err error) bool {
	for _, v := range e.error {
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
func Wrap(originalError, newError Error) Error {
	if newError == nil {
		panic("the parameter newError cannot be nil")
	}

	//error
	if originalError == nil {
		originalError = newError.Copy()
	}
	originalError.addError(newError)

	//return
	return originalError
}
