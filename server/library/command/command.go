package command

import (
	"hash/crc32"
	"strconv"
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

//GenTraceID
//生成链路跟踪码
func GenTraceID() string {
	traceId := GenRandomSeed()
	return strconv.Itoa(int(traceId))
}
