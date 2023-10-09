/**
 * 全局公共变量
 */

package global

import (
	"os"
	"strings"
)

// 全局变量
var (
	// GIN_MODE gin run environment: 'release' or 'debug'
	GIN_MODE string = os.Getenv("GIN_MODE")
	// 工作目录
	SERVER_WORK_PATH string = strings.ToLower(os.Getenv("GIN_SERVER_RUN_PATH"))
	// 运行环境: local、dev、test、prod
	SERVER_RUN_ENV string = strings.ToLower(os.Getenv("GIN_SERVER_RUN_ENV"))
)

// init
func init() {
	//gin运行模式
	if GIN_MODE == "" {
		GIN_MODE = "debug"
	}
	//server运行环境
	if SERVER_RUN_ENV == "" {
		SERVER_RUN_ENV = "local"
	}
}
