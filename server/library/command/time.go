package command

import (
	"time"
)

func init() {
	//设置time时区
	time.Local, _ = time.LoadLocation("GMT") //标准时间
	//time.Local = time.FixedZone("EST", -5*3600)
}

// GetNowDate 年-月-日
func NowDate() string {
	tz, _ := time.LoadLocation("Local") //设置时区
	return time.Unix(time.Now().In(tz).Unix(), 0).In(tz).Format(time.DateOnly)
}

// GetNowDateHour 年-月-日 时:分:钞
func NowDateHour() string {
	tz, _ := time.LoadLocation("Local") //设置时区
	return time.Unix(time.Now().In(tz).Unix(), 0).In(tz).Format(time.DateTime)
}

// ParseTimeFromString 根据字符串(年-月-日 时:分:钞)解析对应时间戳
func ParseTimeFromString(data string) (int, error) {
	tz, _ := time.LoadLocation("Local") //设置时区
	t, err := time.Parse(time.DateTime, data)
	if err != nil {
		return 0, err
	}
	return int(t.In(tz).Unix()), nil
}

// TransDateByTime 年-月-日
func TransDateByTime(date int64) string {
	tz, _ := time.LoadLocation("Local") //设置时区
	return time.Unix(date, 0).In(tz).Format(time.DateOnly)
}

// TransDateByTime 年-月-日 时:分:钞
func TransDateHourByTime(itime int64) string {
	tz, _ := time.LoadLocation("Local") //设置时区
	return time.Unix(itime, 0).In(tz).Format(time.DateTime)
}

// WhichTimestamp 获取指定时间戳零点时间戳
func WhichTimestamp(itime int64) int {
	tz, _ := time.LoadLocation("Local") //设置时区
	t := time.Unix(itime, 0)
	return int(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, tz).In(tz).Unix())
}

// TodayTimestamp 获取今日零点时间戳
func TodayTimestamp() int {
	timeZone, _ := time.LoadLocation("Local") //设置时区
	t := time.Now().In(timeZone)
	return int(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).In(timeZone).Unix())
}

// TomorrowTimestamp 获取明天零点时间戳
func TomorrowTimestamp() int {
	tz, _ := time.LoadLocation("Local") //设置时区
	t := time.Now().In(tz)
	return int(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).In(tz).Add(time.Second * 86400).In(tz).Unix())
}
