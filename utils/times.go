package utils

import (
	"time"
	"math"
)
//获取现在的Unix时间戳
func NowTimeStamp()  int64 {
	nowtime := time.Now().Unix()
	return nowtime
}
//获取时间戳(毫秒)
func NowTimeStampMs() int64 {
	nowtime 	:= 		time.Now()
	ms 		   	:= 	   	nowtime.UnixNano() / 1e6
	return ms
}
//获取时间戳(纳秒)
func NowTimeStampNs() int64 {
	nowtime 	:= 		time.Now()
	ms 		   	:= 	   	nowtime.UnixNano()
	return ms
}
//日期格式转换
func DateFormat(times time.Time,types int) string{
	if times.IsZero(){
		return ""
	}
	var format string
	if types <= 1{
		format = "2006-01-02 15:04:05"
	}else if types == 2{
		format = "2006-01-02 15:04"
	}else if types == 3{
		format = "2006-01-02 15"
	}else if types == 4{
		format = "2006-01-02"
	}else if types == 5{
		format = "2006-01"
	}else if types == 6{
		format = "2006"
	}else{
		format = "2006-01-02 15:04:05"
	}
	dateFormat := time.Unix(times.Unix(), 0).Format(format)
	return dateFormat
}
//获取本月的开始时间戳
func GetNowMonthFirst() int64{
	now 							:=	time.Now()
	currentYear, currentMonth, _ 	:= 	now.Date()
	currentLocation 				:= 	now.Location()
	firstOfMonth 					:= 	time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	firstOfMonthUnix				:= 	firstOfMonth.Unix()
	return firstOfMonthUnix
}
//获取本月的结束时间戳
func GetNowMonthEnd() int64{
	now 							:=	time.Now()
	currentYear, currentMonth, _ 	:= 	now.Date()
	currentLocation 				:= 	now.Location()
	firstOfMonth 					:= 	time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth 					:= firstOfMonth.AddDate(0, 1, -1)
	lastOfMonthUnix					:= 	lastOfMonth.Unix()+86399
	return lastOfMonthUnix
}
//返回现在时间
func NowTimeUTC()  time.Time {
	nowTime := time.Unix(NowTimeStamp(),0).UTC()
	return nowTime
}
//两个时间差相隔多少天
func DaysApart(startTime,endTime int64) int64{
	if startTime <= 0 || endTime <= 0{
		return 0
	}
	if startTime >= endTime {
		return 0
	}
	differ := math.Ceil(StringFloat64(Int64String((endTime - startTime) / 86400)))
	return StringInt64(Float64String(differ))
}