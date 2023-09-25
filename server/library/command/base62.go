package command

import (
	"math"
	"strings"
)

var chars string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// 10进制转62进制
func Encode62(num int) string {
	bytes := []byte{}
	for num > 0 {
		bytes = append(bytes, chars[num%62])
		num = num / 62
	}
	for left, right := 0, len(bytes)-1; left < right; left, right = left+1, right-1 {
		bytes[left], bytes[right] = bytes[right], bytes[left]
	}

	return string(bytes)
}

// 62进制转10进制
func Decode62(str string) int64 {
	var num int64
	n := len(str)
	for i := 0; i < n; i++ {
		pos := strings.IndexByte(chars, str[i])
		num += int64(math.Pow(62, float64(n-i-1)) * float64(pos))
	}
	return num
}

//func main() {
//	mobile := 15821793512
//	str := Encode62(mobile)
//	fmt.Println(str)//output: HGkdFw
//	num := Decode62(str)
//	fmt.Println(num)//output: 15821793512
//｝
