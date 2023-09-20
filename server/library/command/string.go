package command

import (
	"math/rand"
	"unsafe"
)

// StringGenRandom 生成随机字符串
func StringGenRandom(count int, letters ...byte) []byte {
	if len(letters) == 0 {
		letters = []byte(`abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ`)
	}
	length := len(letters)
	rnd := rand.New(rand.NewSource(GenRandom()))
	rnd.Shuffle(length, func(i, j int) {
		letters[i], letters[j] = letters[j], letters[i]
	})
	newString := make([]byte, 0, count)
	for m := 0; m < count; m++ {
		newString = append(newString, letters[rnd.Intn(length)])
	}

	return newString
}

// StringShuffle 随机打乱字符串
func StringShuffle(s string) string {
	re := []rune(s)
	rnd := rand.New(rand.NewSource(GenRandom()))
	rnd.Shuffle(len(re), func(i, j int) {
		re[i], re[j] = re[j], re[i]
	})

	return string(re)
}

// BytesToString []byte转string
func BytesToString(b []byte) string {
	return unsafe.String(&b[0], len(b))
}

// StringToBytes string 转[]byte
func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
