package command

import (
	"math/rand"
	"strconv"
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
