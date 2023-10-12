package command

import (
	crand "crypto/rand"
	"math"
	"math/big"
)

// Round 四舍五入,ROUND_HALF_UP 模式实现
// 返回将 val 根据指定精度 precision(十进制小数点后数字的数目)进行四舍五入的结果.precision 也可以是负数或零.
func Rounding(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}

// GetRandom 随即数
// 在一个区间内求随机数, 不含上下限(min, max)
func Random(min, max int) int {
	if min > max {
		min, max = max, min
	}
	min = min + 1
	rnd := GenRand()
	return rnd.Intn(max-min) + min
}

// GenRandomWithSides 随即数
// 在一个区间内求随机数, 包含上下限[min, max]
func GenRandomWithSides(min, max int) int64 {
	if min > max {
		min, max = max, min
	}
	rnd := GenRand()
	return int64(rnd.Intn(max-min+1) + min)
}

// GenRandom 生成随机数
func GenRandom() int64 {
	randomNum, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	return randomNum.Int64()
}
