package main

import (
	"fmt"
	"time"
)

func main() {
	//获取时间戳
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)

	//格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(tm.Format("2006-01-02 15:04"))
	fmt.Println(tm.Format("02/01/2006 15:04:05 PM"))

	//从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串
	loc, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Println(err)
	}
	tm2, _ := time.ParseInLocation("2006-01-02 15:04", "2020-10-13 13:55", loc)
	tm3 := tm2.Unix()
	fmt.Println(tm3)
	tm4 := time.Unix(tm3, 0)
	fmt.Println(tm4.Format("2006-01-02 15:04"))
}
