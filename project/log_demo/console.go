package mylogger

import (
	"fmt"
	"time"
)

type LogLevel uint16

// Logger 日志结构
type ConcoleLogger struct {
	Level LogLevel
}

func NewConsoleLogger(levelStr string) ConcoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConcoleLogger{
		Level: level,
	}
}
func (c *ConcoleLogger) enable(logger LogLevel) bool {
	return c.Level <= logger
}

func (c *ConcoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		funcName, fileName, lineNo := getInfo(3)

		now := time.Now()
		//now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), GetLogSrting(lv), funcName, fileName, lineNo, msg)

	}
}
func (c *ConcoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c *ConcoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

func (c *ConcoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

func (c *ConcoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

func (c *ConcoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
