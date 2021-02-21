package mylogger

import (
	"fmt"
	"time"
)

// 往终端写日志
//type LogLevel uint16
//
//const (
//	UNKNOWN LogLevel = iota
//	DEBUG
//	INFO
//	WARNING
//	ERROR
//	FALAT
//)
//// Logger 日志结构
//type Logger struct {
//	Level LogLevel
//}
//
//func parseLogLevel(s string) (LogLevel, error) {
//	s = strings.ToLower(s)
//	switch s {
//	case "debug":
//		return DEBUG,nil
//	case "info":
//		return INFO,nil
//	case "warning":
//		return WARNING,nil
//	case "error":
//		return ERROR,nil
//	case "falat":
//		return FALAT,nil
//	default:
//		err := errors.New("无效的日志级别")
//		return UNKNOWN,err
//	}
//}
// NewLog
func Newlog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func log(lv LogLevel, msg string) {
	funcName, fileName, lineNo := getInfo(2)
	now := time.Now()
	//now.Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), GetLogSrting(lv), funcName, fileName, lineNo, msg)
}
func (l Logger) enable(logger LogLevel) bool {
	return l.Level <= logger
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		log(DEBUG, msg)
	}
}
func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		log(INFO, msg)
	}
}
func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		log(WARNING, msg)
	}
}
func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		log(ERROR, msg)
	}
}
func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		log(FATAL, msg)
	}
}
