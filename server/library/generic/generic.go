package generic

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Integer interface {
	Int | Uint
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Integer | Float
}

type String interface {
	~string
}

type NumberString interface {
	Number | String
}
