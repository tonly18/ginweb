package command

import (
	"math/rand"
	"reflect"
	"unsafe"
)

//StringGenRandom 生成随机字符串
func StringGenRandom(count int, letters ...byte) []byte {
	rand.Seed(GenRandomSeed())
	if len(letters) == 0 {
		letters = []byte(`abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ`)
	}
	length := len(letters)
	for i := 0; i < 5; i++ {
		rand.Seed(GenRandomSeed())
		rand.Shuffle(length, func(i, j int) {
			letters[i], letters[j] = letters[j], letters[i]
		})
	}
	newStr := make([]byte, 0, count)
	for m := 0; m < count; m++ {
		rand.Seed(GenRandomSeed())
		newStr = append(newStr, letters[rand.Intn(length)])
	}

	return newStr
}

//StringShuffle 随机打乱字符串
func StringShuffle(s string) string {
	ru := []rune(s)
	rand.Seed(GenRandomSeed())
	rand.Shuffle(len(ru), func(i, j int) {
		ru[i], ru[j] = ru[j], ru[i]
	})

	return string(ru)
}

//B2String []byte转string
func B2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//S2Byte string 转[]byte
func S2Byte(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}
