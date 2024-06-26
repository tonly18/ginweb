package config

import (
	"fmt"
	"github.com/spf13/viper"
	"server/global"
)

// 获取配置文件并解析到指定的struck
func loadConfigFile(fname string, configStruck any) error {
	viper.AddConfigPath(fmt.Sprintf("%v/conf", global.SERVER_WORK_PATH))
	viper.SetConfigName(getConfigFileName(fname))
	viper.SetConfigType("toml")

	//read
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Fatal Error Config File, File Name: %s, Error: %s", fname, err)
	}

	//配置文件监控
	//viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("viper Config file changed......", e.Name)
	//})

	//parse
	return viper.Unmarshal(configStruck)
}

// 获取配置文件
func getConfigFileName(fname string) string {
	return fmt.Sprintf(`%v_%v`, fname, global.SERVER_RUN_ENV)
}
