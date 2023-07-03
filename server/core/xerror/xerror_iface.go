package xerror

type Error interface {
	GetErr() error
	SetErr(error)
	GetCode() int32
	SetCode(int32)
	GetMsg() string
	SetMsg(string)
	GetType() int8
	SetType(int8)
	Error() string
	AddError(Error) Error
	GetErrorList() []Error
	Copy() Error
	Is(error) bool
	Contain(error) bool
}
