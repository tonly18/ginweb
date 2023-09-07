package dao

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cast"
	"hash/crc32"
	"reflect"
	"strconv"
	"unsafe"
)

//字节切片转字符串
func b2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//字符串转字节切片
func s2Bytes(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}

//生成对应hash值
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

func genEntity(length int) []any {
	entity := make([]any, 0, length)
	for i := 0; i < length; i++ {
		entity = append(entity, new(sql.RawBytes))
	}

	return entity
}

func genRecord(data []any, fields []string) map[string]any {
	record := make(map[string]any, len(fields))
	for k, v := range data {
		record[fields[k]] = b2String(*v.(*sql.RawBytes))
	}

	return record
}
