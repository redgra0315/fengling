package main

import (
	"fmt"
	"time"
)

// 时区问题

func f1() {
	now := time.Now() // 本地时间
	fmt.Println(now)
	// 明天的这个时间
	time.Parse("2006-01-02 15:04:05", "2021-03-13 16:49:23")
	// 按照东八区的时间和格式
	// 根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
		return
	}
	// 按照指定时区解析时间
	timeobj, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-03-13 16:49:23", loc)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(timeobj)

	// 时间对象相减
	td := timeobj.Sub(now)
	fmt.Println(td)
}
func main() {
	f1()
	fmt.Println("---------")
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//	时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	// time.Unix()
	ret := time.Unix(1615622634, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Month())
	fmt.Println(ret.Date())
	fmt.Println(ret.Add(24 * time.Hour))

	//	 定时器
	//timer := time.Tick(time.Second)
	//for t := range timer {
	//	fmt.Println(t) //1秒中执行一次
	//
	//}

	// 时间格式化
	fmt.Println(now.Format("2006-01-02 15:04:05 06"))

	// 精确到毫秒  2006-01-02  15:04:05.000   几个零 代表都面显示毫秒位数
	fmt.Println(now.Format("2006-01-02 15:04:05.0000"))

	// 按照对应的格式解析字符串类型的时间
	timeobj, _ := time.Parse("2006-01-02", "2021-03-14")
	fmt.Println(timeobj)
	fmt.Println(timeobj.Unix())

	// 时间相减
	nextYear := timeobj.UTC()
	fmt.Println(nextYear.Sub(timeobj))
	time.Sleep(1 * time.Second)
}
