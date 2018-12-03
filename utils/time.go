package utils

import "time"

var formatDateF = "2006-01-02"
var formatDateT = "15:04:05"

func dateFormat(format string) string {
	return time.Now().Format(format)
}

func DateF() string {
	return dateFormat(formatDateF)
}

func DateT() string {
	return dateFormat(formatDateT)
}

func DateUTC() string {
	return DateF() + "T" + DateT() + "Z"
}
