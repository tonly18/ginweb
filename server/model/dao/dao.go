/*
 * error code: 30000000 ` 30000999
 */

package dao

import (
	"errors"
	"fmt"
	"server/library/command"
)

var (
	ErrorNoRows = errors.New("sql: no rows in result set(dao)")
)

//获取表名
func getTableName(key any, table string) string {
	hash := command.HashValue(key) % uint64(2)
	tblSuffix := fmt.Sprintf(`%04d`, hash)

	return fmt.Sprintf(`%v_%v`, table, tblSuffix)
}
