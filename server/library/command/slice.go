package command

import (
	"bytes"
	"cmp"
	"fmt"
	"math/rand"
	"server/library/generic"
	"slices"
	"sort"
)

// SliceTrans slice数据类型之间的相互转换
func SliceTrans[E generic.Number, T generic.Number](s []E) []T {
	slice := make([]T, 0, len(s))
	for _, v := range s {
		slice = append(slice, T(v))
	}

	//return
	return slice
}

// SliceShuffle 随机打乱[]int
func SliceShuffle[T comparable](s []T) []T {
	rnd := rand.New(rand.NewSource(GenRandom()))
	rnd.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})

	return s
}

// SliceRandomElement 切片中随机N个元素
func SliceRandomElement[T comparable](s []T, count int) []T {
	length := len(s)
	if length == 0 || count < 1 || length < count {
		return s
	}

	rnd := rand.New(rand.NewSource(GenRandom()))
	rnd.Shuffle(length, func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})

	ret := make([]T, 0, count)
	for i := 0; i < count; i++ {
		length = len(s)
		rnd = rand.New(rand.NewSource(GenRandom()))
		ind := rnd.Intn(length)
		ret = append(ret, s[ind])
		s = SliceRemoveElement(s, s[ind]) //从切片中删除已随机的元素
	}

	//return
	return ret
}

// SliceIntRemoveElement 切片中删除指定的元素
func SliceRemoveElement[T comparable](s []T, ele T) []T {
	j := 0
	for _, v := range s {
		if v != ele {
			s[j] = v
			j++
		}
	}
	return s[:j]
}

// SliceRemoveRange 切片删除(开始位置,结束位置),包含两端
func SliceRemoveRange[T comparable](s []T, i, j int) []T {
	return append(s[:i], s[j+2:]...)
}

// SliceRemoveByIndex 根据下标切片删除
func SliceRemoveByIndex[T comparable](s []T, idx int) []T {
	if idx < 0 || idx > len(s)-1 {
		return s
	}
	return append(s[:idx], s[idx+1:]...)
}

// SliceRemoveRepeat 就地删除重复元素(元素可比较)
func SliceRemoveRepeat[T cmp.Ordered](s []T) []T {
	if len(s) < 2 {
		return s
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	x := 0
	for i := 1; i < len(s); i++ {
		if s[x] == s[i] {
			continue
		}
		x++
		s[x] = s[i]
	}

	return s[:x+1]
}

// SliceCopy 切片复制
func SliceCopy[T comparable](s []T) []T {
	b := make([]T, 0, len(s))
	copy(b, s)

	return b
}

// SliceHeaderPush 从头部将元素x插入切片s
func SliceHeaderPush[T comparable](s []T, x T) []T {
	data := make([]T, 0, len(s)+1)
	data = append(data, x)
	return append(data, s...)
}

// SliceTailPush 从头部将元素x插入切片s
func SliceTailPush[T comparable](s []T, x T) []T {
	return append(s, x)
}

// SliceInsert 将元素x插入切片s的i(索引)处
func SliceInsert[T comparable](s []T, x T, i int) []T {
	if i < 0 || i > len(s)-1 {
		return s
	}
	return append(s[:i], append([]T{x}, s[i:]...)...)
}

// SliceHeaderPop 从头部弹出元素x
func SliceHeaderPop[E comparable](s []E) (E, []E) {
	if len(s) == 0 {
		return *new(E), nil
	}
	return s[0], s[1:]
}

// SliceTailPop 从尾部弹出元素x
func SliceTailPop[T comparable](s []T) (T, []T) {
	var t T
	if len(s) == 0 {
		return t, nil
	}
	return s[len(s)-1], s[:len(s)-1]
}

// SliceToString slice转string 默认使用逗号(,)为分割符
func SliceToString[T comparable](data []T, sep string) string {
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

// SliceMax 获取切片中最大值
func SliceMax[T cmp.Ordered](s []T) T {
	return slices.Max(s)
}

// SliceMin 获取切片中最小值
func SliceMin[T cmp.Ordered](s []T) T {
	return slices.Min(s)
}

// SliceSum 获取切片中所有元素之和
func SliceSum[T cmp.Ordered](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

// SliceContains 是否包含指定的值
func SliceContains[E comparable](s []E, v E) bool {
	return slices.Contains(s, v)
}
