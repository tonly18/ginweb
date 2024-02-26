package command

import (
	"time"
)

func init() {
	//设置time时区
	//time.Local, _ = time.LoadLocation("GMT") //标准时间
	//time.Local = time.FixedZone("GMT", 8*3600)//GMT+8小时
	time.Local, _ = time.LoadLocation("Asia/Shanghai") //上海
}

// 当前时间戳
func NowTime() int64 {
	tz, _ := time.LoadLocation("Local") //设置时区
	return time.Now().In(tz).Unix()
}

// 当前年-月-日
func NowDate() string {
	tz, _ := time.LoadLocation("Local") //设置时区
	return time.Unix(time.Now().In(tz).Unix(), 0).In(tz).Format(time.DateOnly)
}

// 当前年-月-日 时:分:钞
func NowDateHour() string {
	tz, _ := time.LoadLocation("Local") //设置时区
	return time.Unix(time.Now().In(tz).Unix(), 0).In(tz).Format(time.DateTime)
}

// 字符串 => 时间戳
func ParseTimeFromString(data string, mark int8) time.Time {
	tz, _ := time.LoadLocation("Local") //设置时区
	if mark == 1 {
		t, _ := time.Parse(time.DateTime, data)
		return t.In(tz)
	}
	if mark == 2 {
		t, _ := time.Parse("2006/01/02 15:04:05", data)
		return t.In(tz)
	}

	return time.Time{}.In(tz)
}

// 时间戳 => 年-月-日
func TransDateByTime(date int64) time.Time {
	tz, _ := time.LoadLocation("Local") //设置时区
	return time.Unix(date, 0).In(tz)
}

// 时间戳 => 年-月-日 时:分:钞
func TransDateHourByTime(itime int64) time.Time {
	tz, _ := time.LoadLocation("Local") //设置时区
	return time.Unix(itime, 0).In(tz)
}

// 获取指定时间戳零点时间
func WhichTimestamp(itime int64) time.Time {
	tz, _ := time.LoadLocation("Local") //设置时区
	t := time.Unix(itime, 0)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, tz).In(tz)
}

// 获取今日零点时间
func YesterdayTimestamp() time.Time {
	tz, _ := time.LoadLocation("Local") //设置时区
	t := time.Now().In(tz)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).In(tz).Add(time.Second * -86400).In(tz)
}

// 获取今日零点时间
func TodayTimestamp() time.Time {
	tz, _ := time.LoadLocation("Local") //设置时区
	t := time.Now().In(tz)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).In(tz)
}

// 获取明天零点时间
func TomorrowTimestamp() time.Time {
	tz, _ := time.LoadLocation("Local") //设置时区
	t := time.Now().In(tz)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).In(tz).Add(time.Second * 86400).In(tz)
}

// 获取指定月份第一天、最后一天时间
// 相当: 2024-01-01 00:00:00 - 2024-01-31 23:59:59
// @params:
//
//	month int	-1上月|0本月|1下月 以此类推
//
// @return:
//
//	time.Time
//	time.Time
func GetMonthFirstAndLastDay(month int) (time.Time, time.Time) {
	tz, _ := time.LoadLocation("Local") //设置时区
	y, m, _ := time.Now().Date()
	m = time.Month(int(m) + month)
	firstDay := time.Date(y, m, 1, 0, 0, 0, 0, tz)
	lastDay := time.Unix(firstDay.AddDate(0, 1, -1).Unix(), 0)

	return firstDay, lastDay
}

// 获取指定周的第一天(周一)、最后一天(周日)时间
// 第一天为周一,最后一天为周日
// @params:
//
//	week int	-1上周|0本周|1下周 以此类推
//
// @return:
//
//	time.Time
//	time.Time
func GetWeekFirstAndLastDay(week int) (time.Time, time.Time) {
	tz, _ := time.LoadLocation("Local") //设置时区

	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	y, m, d := now.Date()
	thisWeek := time.Date(y, m, d, 0, 0, 0, 0, tz)
	startTime := thisWeek.AddDate(0, 0, offset+7*week)
	endTime := thisWeek.AddDate(0, 0, offset+6+7*week)

	return startTime, endTime
}
