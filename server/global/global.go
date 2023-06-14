/**
 * 全局公共变量
 */

package global

import (
	"os"
	"strings"
)

//常量定义
const (
//...
)

//全局变量
var (
	//工作目录
	SERVER_WORK_PATH_ENV, _ = os.Getwd()

	//GINMODE run environment: 'release' or 'debug'
	GINMODE string = os.Getenv("GIN_MODE")
	//运行环境: local、dev、test、prod
	SERVER_RUN_ENV string = strings.ToLower(os.Getenv("GIN_SERVER_RUN_ENV"))
	//配置文件路径
	SERVER_CONFIG_PATH string = os.Getenv("GIN_SERVER_CONFIG_PATH_ENV")
)

//init
func init() {
	//gin运行模式
	if GINMODE == "" {
		GINMODE = "debug"
	}
	//server运行环境
	if SERVER_RUN_ENV == "" {
		SERVER_RUN_ENV = "local"
	}
	//配置文件path
	if SERVER_CONFIG_PATH == "" {
		SERVER_CONFIG_PATH = SERVER_WORK_PATH_ENV + "/conf"
	}
}
