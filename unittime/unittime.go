package unittime

import (
	"fmt"
	"strconv"
	"time"
)

//检查刷新时间
func CheckTimeUpdate(updateTime int64, timeType string) bool {
	switch timeType {
	case "F":
		break //不刷新
	case "M":
		return NowMonthGreaterSomeTime(updateTime) //月刷新
	case "D":
		return NowDayGreaterSomeTime(updateTime) //日刷新
	case "W":
		return NowWeekGreaterSomeTime(updateTime) //周刷新
	}
	return false
}

//是否隔天
func NowDayGreaterSomeTime(someTime int64) bool {
	nowDate := time.Now()
	someDate := time.Unix(int64(someTime), 0)
	nowDateInt, _ := strconv.Atoi(nowDate.Format("20060102"))
	someDateint, _ := strconv.Atoi(someDate.Format("20060102"))
	if nowDateInt != someDateint {
		return true
	}
	return false
}

//是否隔月
func NowMonthGreaterSomeTime(someTime int64) bool {
	nowDate := time.Now()
	someDate := time.Unix(int64(someTime), 0)
	nowDateInt, _ := strconv.Atoi(nowDate.Format("20060102"))
	someDateint, _ := strconv.Atoi(someDate.Format("20060102"))
	if nowDateInt != someDateint {
		return true
	}
	return false
}

//是否隔周
func NowWeekGreaterSomeTime(someTime int64) bool {
	nowDate := time.Now()
	someDate := time.Unix(int64(someTime), 0)
	nowYear, nowWeek := nowDate.ISOWeek()
	someYear, someWeek := someDate.ISOWeek()
	//跨年或者垮周都算
	if nowYear > someYear {
		return true
	} else if nowWeek > someWeek {
		return true
	}
	return false
}

//当日0点时间
func DateZeroTime(someTime int64) int64 {
	someDate := time.Unix(int64(someTime), 0)
	timeStr := someDate.Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	timeNumber := t.Unix() - 3600*8
	return timeNumber
}

//下周0点时间
func NextWeekOne() int64 {
	nowDate := time.Now()
	day := nowDate.Weekday().String()

	switch day {
	case "Sunday":
		return 0
	case "Monday":
		return 6
	case "Tuesday":
		return 5
	case "Wednesday":
		return 4
	case "Thursday":
		return 3
	case "Friday":
		return 2
	case "Saturday":
		return 1
	}
	return 0
}

func FormatToTimestamp(formatDate string) time.Time {
	loc, _ := time.LoadLocation("Local")                              //重要：获取时区
	theTime, _ := time.ParseInLocation("2006-01-02", formatDate, loc) //使用模板在对应时区转化为time.time类型
	return theTime
}

func HourTimestamp() int64 {
	now := time.Now()
	timestamp := now.Unix() - int64(now.Second()) - int64((60 * now.Minute()))
	fmt.Println(timestamp, time.Unix(timestamp, 0), now.Unix())
	return timestamp
}
