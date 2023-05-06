package config

//bag init value
const bagValue = `[{"id":100,"count":10,"time":0}]`
const userValue = `{"avatar":0,"name":"","level":1,"exp":0,"gold":0,"server_id":0}`

//Global
var globalDefaultValue = map[string]string{
	"bag":  bagValue,
	"user": userValue,
}

//获取初始值
func GetDefaultDBValue(tbl string) string {
	if data, ok := globalDefaultValue[tbl]; ok {
		return data
	}

	return ""
}

//获取所有初始值
func GetDefaultDBValueAll(tbl string) map[string]string {
	return globalDefaultValue
}
