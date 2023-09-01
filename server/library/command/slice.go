package command

import (
	"bytes"
	"fmt"
	"math/rand"
	"server/library/generic"
)

//SliceTrans slice数据类型之间的相互转换
func SliceTrans[I generic.Number, R generic.Number](s []I) []R {
	slice := make([]R, 0, len(s))
	for _, v := range s {
		slice = append(slice, R(v))
	}

	//return
	return slice
}

//SliceShuffle 随机打乱[]int
func SliceShuffle[T generic.NumberString](s []T) []T {
	rand.Seed(GenRandomSeed())
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})

	return s
}

//SliceGetRandomElement 切片中随机N个元素
func SliceGetRandomElement[T generic.NumberString](s []T, count int) []T {
	length := len(s)
	if length == 0 || count < 1 || length < count {
		return s
	}

	rand.Seed(GenRandomSeed())
	rand.Shuffle(length, func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})

	ret := make([]T, 0, count)
	for i := 0; i < count; i++ {
		length = len(s)
		rand.Seed(GenRandomSeed())
		ind := rand.Intn(length)
		ret = append(ret, s[ind])
		s = SliceRemoveElement(s, s[ind]) //从切片中删除已随机的元素
	}

	//return
	return ret
}

//SliceIntRemoveElement 切片中删除指定的元素
func SliceRemoveElement[T generic.NumberString](s []T, ele T) []T {
	j := 0
	for _, v := range s {
		if v != ele {
			s[j] = v
			j++
		}
	}
	return s[:j]
}

//SliceCopy 切片复制
func SliceCopy[T generic.NumberString](s []T) []T {
	b := make([]T, len(s))
	copy(b, s)

	return b
}

//SliceRemoveRange 切片删除(开始位置,结束位置),包含两端
func SliceRemoveRange[T generic.NumberString](s []T, i, j int) []T {
	return append(s[:i], s[j+1:]...)
}

//SliceRemoveByIndex 根据下标切片删除
func SliceRemoveByIndex[T generic.NumberString](s []T, idx int) []T {
	if idx < 0 || idx > len(s)-1 {
		return s
	}
	return append(s[:idx], s[idx+1:]...)
}

//SliceHeaderPush 从头部将元素x插入切片s
func SliceHeaderPush[T generic.NumberString](s []T, x T) []T {
	data := make([]T, 0, len(s)+1)
	data = append(data, x)
	return append(data, s...)
}

//SliceTailPush 从头部将元素x插入切片s
func SliceTailPush[T generic.NumberString](s []T, x T) []T {
	return append(s, x)
}

//SliceInsert 将元素x插入切片s的i(索引)处
func SliceInsert[T generic.NumberString](s []T, x T, i int) []T {
	if i < 0 || i > len(s)-1 {
		return s
	}
	return append(s[:i], append([]T{x}, s[i:]...)...)
}

//SliceHeaderPop 从头部弹出元素x
func SliceHeaderPop[T generic.NumberString](s []T) (T, []T) {
	var t T
	if len(s) == 0 {
		return t, nil
	}
	return s[0], s[1:]
}

//SliceTailPop 从尾部弹出元素x
func SliceTailPop[T generic.NumberString](s []T) (T, []T) {
	var t T
	if len(s) == 0 {
		return t, nil
	}
	return s[len(s)-1], s[:len(s)-1]
}

//SliceToString slice转string 默认使用逗号(,)为分割符
func SliceToString[T generic.NumberString](data []T, sep string) string {
	length := len(data)
	if length == 0 {
		return ""
	}

	maxIndex := length - 1
	buffer := bytes.NewBuffer(make([]byte, 0, len(data)*2))
	for k, v := range data {
		if k == maxIndex {
			buffer.Write([]byte(fmt.Sprintf(`%v`, v)))
		} else {
			buffer.Write([]byte(fmt.Sprintf(`%v%v`, v, sep)))
		}
	}

	//return
	return buffer.String()
}

//SliceGetMinElement 获取切片中最小值
func SliceGetMinElement[T generic.Number](s []T) (int, T) {
	accurateBit := 1000
	index, lowestValue := 0, s[0]
	for k, v := range s {
		if int(v*T(accurateBit)) < int(lowestValue*T(accurateBit)) {
			index, lowestValue = k, v
		}
	}

	return index, lowestValue
}

//SliceSum 获取切片中所有元素之和
func SliceSum[T generic.Number](s []T) int {
	var sum int
	for _, v := range s {
		sum += int(v)
	}
	return sum
}

//SliceSum 就地删除重复元素(元素可比较)
func SliceRemoveRepeat(s []int) []int {
	j := 0
	//sort.Ints(s)
	//for i := 1; i < len(s); i++ {
	//	if s[j] == s[i] {
	//		continue
	//	}
	//	j++
	//
	//	//需要保存原始数据时
	//	//in[i], in[j] = in[j], in[i]
	//	//只需要保存需要的数据时
	//	s[j] = s[i]
	//}
	//
	return s[:j+1]
}
