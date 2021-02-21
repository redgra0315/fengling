package main

import (
	mylogger "github.com/fengling/project/log_demo"
	"time"
)

// 测试自己写的日之苦
func main()  {
	log := mylogger.Newlog()
	for {

		log.Deubug("这是Debug日志")
		log.Info("这是Info日志")
		time.Sleep(time.Second)
	}
}
