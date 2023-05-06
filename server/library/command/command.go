package command

import (
	"github.com/spf13/cast"
	"hash/crc32"
	"math/rand"
	"strconv"
	"time"
)

//HashValue 计算hash值
func HashValue(value any) uint64 {
	switch val := value.(type) {
	case int:
		return uint64(val)
	case uint64:
		return uint64(val)
	case int64:
		return uint64(val)
	case string:
		if v, err := strconv.ParseUint(val, 10, 64); err != nil {
			return uint64(crc32.ChecksumIEEE(Slice(val)))
		} else {
			return uint64(v)
		}
	case []byte:
		return uint64(crc32.ChecksumIEEE(val))
	default:
		//return 0
	}

	return 0
}

//GenRandom 随即数
//在一个区间内求随机数,不含上下限 (min, max)
func GenRandom(min, max int) int {
	min++
	rand.Seed(ShuffleUnixNano())
	return rand.Intn(max-min) + min
}

//GenRandomContainBoth 随即数
//在一个区间内求随机数,含上下限 [min, max]
func GenRandomContainBoth(min, max int) int {
	rand.Seed(ShuffleUnixNano())
	return rand.Intn(max-min+1) + min
}

//GenTraceID
//生成链路跟踪码
func GenTraceID() string {
	traceId := ShuffleUnixNano()
	return strconv.Itoa(int(traceId))
}

//ShuffleUnixNano
//乱序UnixNano
func ShuffleUnixNano() int64 {
	nano := []byte(cast.ToString(time.Now().UnixNano()))
	rand.Shuffle(len(nano), func(i, j int) {
		nano[i], nano[j] = nano[j], nano[i]
	})
	traceId, _ := strconv.ParseInt(string(nano), 10, 64)

	return traceId
}
