package config

//Redis Key 前缀
const REDIS_PREFIX = "yusheng_"

//Redis Key
const (
	//帐号信息
	REDIS_KEY_ACCOUNT = REDIS_PREFIX + "account:"
	//角色信息
	REDIS_KEY_PLAYER = REDIS_PREFIX + "player:"
	//货币
	REDIS_KEY_CURRENCY = REDIS_PREFIX + "currency:"
	//背包
	REDIS_KEY_BAG = REDIS_PREFIX + "bag:"
)
