package main

import (
	//console "github.com/fengling/project/log_demo/project/log_demo"
	//"github.com/fengling/project/log_demo/project/log_demo"
	"github.com/fengling/project/log_demo/project/log_demo"
	"time"
)

// 测试自己写的日之苦
func main() {
	log := mylogger.Newlog("info")
	for {

		log.Debug("这是Debug日志")
		log.Info("这是Info日志")
		log.Error("你错了")
		time.Sleep(time.Second)
	}
}

//import (
//	"net/http"
//
//	"github.com/labstack/echo"
//)
//
//func main() {
//	e := echo.New()
//	e.GET("/", func(c echo.Context) error {
//		return c.String(http.StatusOK, "Hello, World!")
//	})
//	e.Logger.Fatal(e.Start(":1323"))
//}
