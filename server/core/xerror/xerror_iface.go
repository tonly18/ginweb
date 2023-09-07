package xerror

type Error interface {
	Error() string

	GetErr() error
	SetErr(error)
	GetCode() uint32
	SetCode(uint32)
	GetMsg() string
	SetMsg(string)

	addError(Error) Error
	GetError() []Error

	Copy() Error
	Is(error) bool
	Contain(error) bool
}
