package config

//config.toml 配置文件
type configStruck struct {
	Title    string
	Http     httpConfig
	Log      logConfig
	Mysql    mysqlConfig
	Redis    redisConfig
	Memcache memConfig
}

//http config
type httpConfig struct {
	Host string
	Port string
}

//log file
type logConfig struct {
	LogFilePath string
}

//mysql config
type mysqlConfig struct {
	Host      string
	Port      string
	Username  string
	Passworld string
}

//redis config
type redisConfig struct {
	Host         string
	Password     string
	MinIdleConns int //最小空闲链接数
	PoolSize     int //链接池最大链接数
	Db           int //redis 库
}

//memConfig config
type memConfig struct {
	Host         string
	MaxIdleConns int //最大空闲链接数
	Timeout      int //链接超时:毫秒
}

//Config data
var Config *configStruck = &configStruck{}

//init
func init() {
	if err := getConfig(); err != nil {
		panic(err)
	}
}

//获取config配置文件
func getConfig() error {
	if Config.Title == "" {
		if err := loadConfigFile("config", Config); err != nil {
			return err
		}
	}

	return nil
}
