package dao

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"hash/crc32"
	"reflect"
	"strconv"
	"unsafe"
)

var (
	ErrorNoRows = errors.New("sql: no rows in result set(dao)")
)

func b2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func s2Bytes(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}

func hashValue(value any) uint64 {
	switch val := value.(type) {
	case int:
		return uint64(val)
	case uint64:
		return uint64(val)
	case int64:
		return uint64(val)
	case string:
		if v, err := strconv.ParseUint(val, 10, 64); err != nil {
			return uint64(crc32.ChecksumIEEE(s2Bytes(val)))
		} else {
			return uint64(v)
		}
	case []byte:
		return uint64(crc32.ChecksumIEEE(val))
	}

	return cast.ToUint64(fmt.Sprintf(`%s`, value))
}

//获取表名
func getTableName(key any, table string) string {
	hash := hashValue(key) % uint64(2)
	tblSuffix := fmt.Sprintf(`%04d`, hash)

	return fmt.Sprintf(`%v_%v`, table, tblSuffix)
}
