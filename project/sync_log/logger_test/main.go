package main

import (
	"github.com/fengling/project/log_demo/project/sync_log/log_demo"
	"time"
)

// 声明一个全局的接口变量
var log mylogger.Logger

// 测试自己写的日之苦
func main() {
	log = mylogger.NewFileLogger("debug", "./", "logged", 10*1024)

	for {
		//id := 10000
		id1 := 10001
		name := "陕西"
		log.Error("你错了 addr:%v", name)
		log.Debug("这是Debug日志")
		log.Info("这是Info日志,id1: %d", id1)
		//log.Error("你错了 addr:%v", name)
		time.Sleep(time.Second)
	}
}
