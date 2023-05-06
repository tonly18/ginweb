package merror

type Error interface {
	GetError() error
	GetCode() int32
	GetMsg() string
	GetType() int8
}
