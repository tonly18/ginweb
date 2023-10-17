package command

import (
	"math/rand"
	"strconv"
	"unsafe"
)

// GenTraceID 生成链路跟踪码
func GenTraceID() string {
	traceId := GenRandom()
	return strconv.Itoa(int(traceId))
}

// GenRand 生成rand.rand
func GenRand() *rand.Rand {
	return rand.New(rand.NewSource(GenRandom()))
}

// IsValueNil 值判空
func IsValueNil(v any) bool {
	if v == nil {
		return true
	}

	// 判断值是否为空
	type eface struct {
		v   int64
		ptr unsafe.Pointer
	}
	efacePtr := (*eface)(unsafe.Pointer(&v))
	if efacePtr == nil {
		return true
	}

	// ok := efaceptr == nil || uintptr(efaceptr.ptr) == 0
	return uintptr(efacePtr.ptr) == 0x0
}
