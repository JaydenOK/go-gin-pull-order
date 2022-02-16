package service

import (
	"log"
	"strconv"
	"time"
)

//获取当前时间戳
func Time() string {
	return time.Now().String()
}

//获取当前日期
func Date() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func StringToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func IntToString(value int) string {
	return strconv.Itoa(value)
}
