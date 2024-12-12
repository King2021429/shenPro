package utils

import (
	"time"
)

// 这里简单模拟一个节假日配置列表，实际应用中可以从配置文件、数据库等读取更全面准确的数据
var holidays = []string{
	// 元旦
	"2025-01-01",
	"2025-01-02",
	"2025-01-03",
	// 春节
	"2025-02-18",
	"2025-02-19",
	"2025-02-20",
	"2025-02-21",
	"2025-02-22",
	"2025-02-23",
	"2025-02-24",
	// 清明节
	"2025-04-04",
	"2025-04-05",
	"2025-04-06",
	// 劳动节
	"2025-05-01",
	"2025-05-02",
	"2025-05-03",
	// 端午节
	"2025-06-20",
	"2025-06-21",
	"2025-06-22",
	// 中秋节、国庆节
	"2025-10-01",
	"2025-10-02",
	"2025-10-03",
	"2025-10-04",
	"2025-10-05",
	"2025-10-06",
	"2025-10-07",
}

func IsWorkingDay() bool {
	now := time.Now()
	weekday := now.Weekday()
	if weekday >= time.Monday && weekday <= time.Friday {
		dateStr := now.Format("2006-01-01")
		for _, holiday := range holidays {
			if dateStr == holiday {
				return false
			}
		}
		return true
	}
	return false
}
