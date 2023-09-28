package dao

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/tonly18/xsql"
	"hash/crc32"
	"server/config"
	"strconv"
	"unsafe"
)

// redis配置
type RedisConfig struct {
	Host     string
	Password string
}

// DB配置
var dbConfig *xsql.Config

// redis配置
var rdConfig *RedisConfig

func init() {
	// DB配置
	dbConfig = &xsql.Config{
		Host:     config.Config.Mysql.Host,
		Port:     cast.ToInt(config.Config.Mysql.Port),
		UserName: config.Config.Mysql.Username,
		Password: config.Config.Mysql.Password,
		DBName:   "test",
	}
	// redis配置
	rdConfig = &RedisConfig{
		Host:     config.Config.Redis.Host,
		Password: config.Config.Redis.Password,
	}
}

// bytesToString []byte转string
func bytesToString(b []byte) string {
	return unsafe.String(&b[0], len(b))
}

// stringToBytes string 转[]byte
func stringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// 生成对应hash值
func hashValue(value any) uint64 {
	switch val := value.(type) {
	case int:
		return uint64(val)
	case uint64:
		return val
	case int64:
		return uint64(val)
	case string:
		if v, err := strconv.ParseUint(val, 10, 64); err != nil {
			return uint64(crc32.ChecksumIEEE(stringToBytes(val)))
		} else {
			return v
		}
	case []byte:
		return uint64(crc32.ChecksumIEEE(val))
	}

	return cast.ToUint64(fmt.Sprintf(`%s`, value))
}

// 获取表名
func getTableName(key any, table string, tblNum ...int) string {
	num := 33 //默认分33张表
	if len(tblNum) > 0 {
		num = tblNum[0]
	}
	hash := hashValue(key) % uint64(num)
	tblSuffix := fmt.Sprintf(`%04d`, hash)

	return fmt.Sprintf(`%v_%v`, table, tblSuffix)
}
