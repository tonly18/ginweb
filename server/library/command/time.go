package command

import (
	"time"
)

func init() {
	//设置time时区
	time.Local, _ = time.LoadLocation("GMT") //标准时间
	//time.Local = time.FixedZone("EST", -5*3600)
}

//GetNowDate 年-月-日
func NowDate() string {
	tz, _ := time.LoadLocation("Local") //设置时区
	return time.Unix(time.Now().In(tz).Unix(), 0).In(tz).Format("2006-01-02")
}

//GetNowDateHour 年-月-日 时:分:钞
func NowDateHour() string {
	tz, _ := time.LoadLocation("Local") //设置时区
	return time.Unix(time.Now().In(tz).Unix(), 0).In(tz).Format("2006-01-02 15:04:05")
}

//ParseTimeFromString 根据字符串(年-月-日 时:分:钞)解析对应时间戳
func ParseTimeFromString(data string) (int, error) {
	tz, _ := time.LoadLocation("Local") //设置时区
	t, err := time.Parse("2006-01-02 15:04:05", data)
	if err != nil {
		return 0, err
	}
	return int(t.In(tz).Unix()), nil
}

//TransDateByTime 年-月-日
func TransDateByTime(data int64) string {
	tz, _ := time.LoadLocation("Local") //设置时区
	return time.Unix(data, 0).In(tz).Format("2006-01-02")
}

//TransDateByTime 年-月-日 时:分:钞
func TransDateHourByTime(itime int64) string {
	tz, _ := time.LoadLocation("Local") //设置时区
	return time.Unix(itime, 0).In(tz).Format("2006-01-02 15:04:05")
}

//WhichTimestamp 获取指定时间戳零点时间戳
func WhichTimestamp(itime int64) int {
	tz, _ := time.LoadLocation("Local") //设置时区
	t := time.Unix(itime, 0)
	return int(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, tz).In(tz).Unix())
}

//TodayHMSTimestamp 获取今日时分秒时间戳
func TodayHMSTimestamp(hour, min, sec int) int {
	timeZone, _ := time.LoadLocation("Local") //设置时区
	t := time.Now().In(timeZone)
	return int(time.Date(t.Year(), t.Month(), t.Day(), hour, min, sec, 0, t.Location()).In(timeZone).Unix())
}

//TodayTimestamp 获取今日零点时间戳
func TodayTimestamp() int {
	timeZone, _ := time.LoadLocation("Local") //设置时区
	t := time.Now().In(timeZone)
	return int(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).In(timeZone).Unix())
}

//TomorrowTimestamp 获取明天零点时间戳
func TomorrowTimestamp() int {
	tz, _ := time.LoadLocation("Local") //设置时区
	t := time.Now().In(tz)
	return int(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).In(tz).Add(time.Second * 86400).In(tz).Unix())
}

// CurrentWeekSunday 本周周日时间戳(零点)
// 一周的第一天为周一,最后一天为周日
func CurrentWeekSunday() int {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	week := now.Weekday()
	offDays := 7 - week // 距离周日相差几天
	if week == time.Sunday {
		offDays = 0
	}
	Sunday := today.Add(time.Duration(offDays) * 86400 * time.Second).In(time.Local)

	return int(Sunday.Unix())
}
