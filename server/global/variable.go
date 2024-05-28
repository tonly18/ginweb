/**
 * 全局公共变量
 */

package global

import (
	"github.com/joho/godotenv"
	"os"
)

// 全局变量
var (
	// 工作目录
	SERVER_WORK_PATH string
	// 运行环境: local、dev、test、prod
	SERVER_RUN_ENV string
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	SERVER_RUN_ENV = os.Getenv("SERVER_RUN_ENV")
	SERVER_WORK_PATH = os.Getenv("SERVER_RUN_PATH")
}
