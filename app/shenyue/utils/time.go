package utils

import (
	"time"
)

// 这里简单模拟一个节假日配置列表，实际应用中可以从配置文件、数据库等读取更全面准确的数据
var holidays = []string{
	"2025-04-04", // 清明节
	"2025-04-05", // 清明节
	"2025-04-06", // 清明节
	"2025-05-01", // 劳动节
	"2025-05-02", // 劳动节
	"2025-05-03", // 劳动节
	"2025-05-04", // 劳动节
	"2025-05-05", // 劳动节
	"2025-05-31", // 端午节
	"2025-06-01", // 端午节
	"2025-06-02", // 端午节
	"2025-10-01", // 国庆节
	"2025-10-02", // 国庆节
	"2025-10-03", // 国庆节
	"2025-10-04", // 国庆节
	"2025-10-05", // 国庆节
	"2025-10-06", // 国庆节
	"2025-10-07", // 国庆节
	"2025-10-08", // 国庆节
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
