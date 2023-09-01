package xerror

type Error interface {
	GetErr() error
	SetErr(error)
	GetCode() uint32
	SetCode(uint32)
	GetMsg() string
	SetMsg(string)
	Error() string
	AddError(Error) Error
	GetErrorStack() []Error
	Copy() Error
	Is(error) bool
	Contain(error) bool
}
