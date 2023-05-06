package utils

import (
	"server/core/logger"
	"server/model/dao"
)

//Finish 程序退出时执行一些清理工作
func Finish() {
	logger.FinishClear()
	dao.FinishClear()
}
