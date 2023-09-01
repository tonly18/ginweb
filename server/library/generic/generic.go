package generic

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type String interface {
	~string
}

type Number interface {
	Int | Uint | Float
}

type NumberInt interface {
	Int | Uint
}

type NumberString interface {
	Int | Uint | Float | String
}
