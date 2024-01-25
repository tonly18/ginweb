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
	// 工作目录
	SERVER_WORK_PATH string = strings.ToLower(os.Getenv("GIN_SERVER_RUN_PATH"))
	// 运行环境: local、dev、test、prod
	SERVER_RUN_ENV string = strings.ToLower(os.Getenv("GIN_SERVER_RUN_ENV"))
)
