package command

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"server/library/generic"
	"strings"
)

//SliceTrans slice数据类型之间的相互转换
func SliceTrans[T generic.Int | generic.Uint, R generic.Int | generic.Uint](s []T) []R {
	slice := make([]R, 0, len(s))
	for _, v := range s {
		slice = append(slice, R(v))
	}

	//return
	return slice
}

//SliceIntShuffle 随机打乱[]int
func SliceIntShuffle(s []int) []int {
	rand.Seed(ShuffleUnixNano())
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})

	return s
}

//SliceIntGetRandomElement 切片中随机N个元素
func SliceIntGetRandomElement(s []int, count int) []int {
	rand.Seed(ShuffleUnixNano())
	length := len(s)
	if length == 0 || count < 1 || length < count {
		return s
	}

	rand.Shuffle(length, func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})

	ret := make([]int, 0, count)
	for i := 0; i < count; i++ {
		length = len(s)
		ind := rand.Intn(length)
		ret = append(ret, s[ind])
		s = SliceIntRemoveElement(s, s[ind])
	}

	//return
	return ret
}

//SliceIntRemoveElement 切片中删除指定的元素
func SliceIntRemoveElement(s []int, ele int) []int {
	j := 0
	for _, v := range s {
		if v != ele {
			s[j] = v
			j++
		}
	}
	return s[:j]
}

//SliceIntCopy 切片复制
func SliceIntCopy(data []int) []int {
	b := make([]int, len(data))
	copy(b, data)

	return b
}

//SliceIntRemoveRange 切片删除(范围),包含两端
func SliceIntRemoveRange(s []int, i, j int) []int {
	return append(s[:i], s[j+1:]...)
}

//SliceIntRemoveByIndex 根据下标切片删除
func SliceIntRemoveByIndex(s []int, idx int) []int {
	if idx < 0 || idx > len(s)-1 {
		return s
	}
	return append(s[:idx], s[idx+1:]...)
}

//SliceIntInsert 将元素x插入切片s的索引i处
func SliceIntInsert(s []int, x, i int) []int {
	if i < 0 || i > len(s)-1 {
		return s
	}
	return append(s[:i], append([]int{x}, s[i:]...)...)
}

//SliceToString slice转string
//默认使用逗号(,)为分割符
func SliceToString(data any, sep ...string) string {
	if reflect.TypeOf(data).Kind() != reflect.Slice {
		return ""
	}

	var result string
	if res, err := json.Marshal(data); err == nil && len(res) > 2 {
		result = string(res)[1 : len(res)-1]
		if len(sep) > 0 {
			result = strings.Replace(result, ",", sep[0], -1)
		}
	}

	//return
	return result
}
