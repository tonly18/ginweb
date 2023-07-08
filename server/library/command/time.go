package command

import (
	"time"
)

func init() {
	//设置time时区
	time.Local, _ = time.LoadLocation("GMT") //标准时间
	//time.Local = time.FixedZone("EST", -5*3600)
}

//ParseTimeFromString 根据字符串解析时间 必须是 2006/1/2 15:04:05 格式的字符串
//
//@param
//@return
func ParseTimeFromString(itime string) (time.Time, error) {
	parse, err := time.Parse("2006/1/2 15:04:05", itime)
	if err != nil {
		return parse, err
	}
	return parse, nil
}

//GetNowDate 年月日
//
//@param
//@return
func GetNowDate() string {
	timeZone, _ := time.LoadLocation("Local") //设置时区
	return time.Unix(time.Now().In(timeZone).Unix(), 0).In(timeZone).Format("2006-01-02")
}

//TranDateByTime 年月日
//
//@param
//@return
func TranDateByTime(itime int) string {
	timeZone, _ := time.LoadLocation("Local") //设置时区
	return time.Unix(int64(itime), 0).In(timeZone).Format("2006-01-02")
}

//GetNowDateHour 年月日-时分钞
//
//@param
//@return
func GetNowDateHour() string {
	timeZone, _ := time.LoadLocation("Local") //设置时区
	return time.Unix(time.Now().In(timeZone).Unix(), 0).In(timeZone).Format("2006-01-02 15:04:05")
}

//GetTodayTimestamp 获取今日零点时间戳
//
//@param
//@return
func GetTodayTimestamp() int {
	timeZone, _ := time.LoadLocation("Local") //设置时区
	t := time.Now().In(timeZone)
	return int(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).In(timeZone).Unix())
}

//GetWhichTimestamp 获取指定时间戳零点时间戳
//
//@param
//@return
func GetWhichTimestamp(itime int) int {
	timeZone, _ := time.LoadLocation("Local") //设置时区
	t := time.Unix(int64(itime), 0)
	return int(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, timeZone).In(timeZone).Unix())
}

//GetTodayHMSTimestamp 获取今日时分秒时间戳
//
//@param
//@return
func GetTodayHMSTimestamp(hour, min, sec int) uint32 {
	timeZone, _ := time.LoadLocation("Local") //设置时区
	t := time.Now().In(timeZone)
	return uint32(time.Date(t.Year(), t.Month(), t.Day(), hour, min, sec, 0, t.Location()).In(timeZone).Unix())
}

//GetTomorrowHMSTimestamp 获取明日时分秒时间戳
//
//@param
//@return
func GetTomorrowHMSTimestamp(hour, min, sec int) uint32 {
	timeZone, _ := time.LoadLocation("Local") //设置时区
	t := time.Now().In(timeZone)
	return uint32(time.Date(t.Year(), t.Month(), t.Day()+1, hour, min, sec, 0, t.Location()).In(timeZone).Unix())
}

//GetTomorrowTimestamp 获取明天零点时间戳
//
//@param
//@return
func GetTomorrowTimestamp() int {
	timeZone, _ := time.LoadLocation("Local") //设置时区
	t := time.Now().In(timeZone)
	return int(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).In(timeZone).Add(time.Second * 86400).In(timeZone).Unix())
}
