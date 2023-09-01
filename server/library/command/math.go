package command

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
)

//Round 四舍五入,ROUND_HALF_UP 模式实现
//返回将 val 根据指定精度 precision(十进制小数点后数字的数目)进行四舍五入的结果.precision 也可以是负数或零.
func Round(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}

//GetRandom 随即数
//在一个区间内求随机数
//不含上下限 (min, max)
func GetRandom(min, max int) int {
	if min > max {
		min, max = max, min
	}
	min = min + 1
	rand.Seed(GenRandomSeed())
	return rand.Intn(max-min) + min
}

//GetRandomWithAll 随即数
//在一个区间内求随机数
//包含上下限 [min, max]
func GetRandomWithBorder(min, max int) int64 {
	if min > max {
		min, max = max, min
	}
	rand.Seed(GenRandomSeed())
	return int64(rand.Intn(max-min+1) + min)
}

//GenRandomSeed 生成随机数
func GenRandomSeed() int64 {
	max := big.NewInt(math.MaxInt64)
	randomNum, _ := crand.Int(crand.Reader, max)

	return randomNum.Int64()
}
