package command

import (
	"encoding/json"
	"strconv"
	"unsafe"
)

// StringGenRandom 生成随机字符串
func StringGenRandom(count int, letters ...byte) []byte {
	if len(letters) == 0 {
		letters = []byte(`abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ`)
	}
	length := len(letters)
	rnd := GenRand()
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
	rnd := GenRand()
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

// ConvertToString 任何类型转换成字符串
func ConvertToString(value any) string {
	var str string
	if value == nil {
		return str
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		str = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		str = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		str = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		str = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		str = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		str = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		str = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		str = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		str = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		str = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		str = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		str = strconv.FormatUint(it, 10)
	case string:
		str = value.(string)
	case []byte:
		str = BytesToString(value.([]byte))
	default:
		valueByte, _ := json.Marshal(value)
		str = BytesToString(valueByte)
	}

	return str
}
