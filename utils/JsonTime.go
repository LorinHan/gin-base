package utils

import (
	"time"
)

const localDateTimeFormat string = "2006-01-02 15:04:05"

/**
 * @description: time.Time类型返回给前端的时间不友好，可以用本类型替代为yyyy-mm-dd hh:mm:ss格式
 */
type Time time.Time

func (l Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(localDateTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(l).AppendFormat(b, localDateTimeFormat)
	b = append(b, '"')
	return b, nil
}

func (l *Time) UnmarshalJSON(b []byte) error {
	now, err := time.ParseInLocation(`"`+localDateTimeFormat+`"`, string(b), time.Local)
	*l = Time(now)
	return err
}

func FormatNow() string {
	formatStr := "2006-01-02"
	return time.Now().Format(formatStr)
}


func MonthDays(year int, month int) (days int) {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30
		} else {
			days = 31
		}
	} else {
		if (((year % 4) == 0 && (year % 100) != 0) || (year % 400) == 0) {
			days = 29
		} else {
			days = 28
		}
	}
	return days
}