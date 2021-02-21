package mylogger

import "fmt"

// 往终端写日志

// Logger 日志结构
type Loger struct {
}

// NewLog
func Newlog() Loger {
	return Loger{}
}

func (l Loger) Debug(msg string) {
	fmt.Fprintln(msg)
}
func (l Loger) Info(msg string) {
	fmt.Fprintln(msg)
}
func (l Loger) Warning(msg string) {
	fmt.Fprintln(msg)
}
func (l Loger) Error(msg string) {
	fmt.Fprintln(msg)
}
func (l Loger) Fatal(msg string) {
	fmt.Fprintln(msg)
}
