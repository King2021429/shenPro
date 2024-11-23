package utils

import "time"

func IsWeekday(date time.Time) bool {
	weekday := date.Weekday()
	return weekday != time.Saturday && weekday != time.Sunday
}
