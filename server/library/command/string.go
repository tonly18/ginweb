package command

import (
	"math/rand"
)

//StringGenRandom 生成随即字符串
func StringGenRandom(count int, str ...string) string {
	rand.Seed(ShuffleUnixNano())
	letters := []byte(`abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ`)
	if len(str) > 0 {
		letters = []byte(str[0])
	}
	length := len(letters)
	for i := 0; i < length; i++ {
		rand.Shuffle(length, func(i, j int) {
			letters[i], letters[j] = letters[j], letters[i]
		})
	}
	newStr := make([]byte, 0, count)
	for m := 0; m < count; m++ {
		newStr = append(newStr, letters[rand.Intn(length)])
	}

	return string(newStr)
}

//StringShuffle 随机打乱字符串
func StringShuffle(s string) string {
	ru := []rune(s)
	rand.Seed(ShuffleUnixNano())
	rand.Shuffle(len(ru), func(i, j int) {
		ru[i], ru[j] = ru[j], ru[i]
	})

	return string(ru)
}
