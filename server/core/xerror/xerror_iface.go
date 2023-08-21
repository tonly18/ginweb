package xerror

type Error interface {
	GetErr() error
	SetErr(error)
	GetCode() uint32
	SetCode(uint32)
	GetMsg() string
	SetMsg(string)
	GetType() int8
	SetType(int8)
	Error() string
	AddError(Error) Error
	GetErrorStack() []Error
	Copy() Error
	Is(error) bool
	Contain(error) bool
}
