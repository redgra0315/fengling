package mylogger

import (
	"fmt"
	"os"
	"path"
)

// 往文件里面写日志相关代码

type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件的名称
	fileLevel   string
	fileObj     *os.File //
	fileObjErr  *os.File
	maxFileSize int64
	logChan     chan *logMsg
}

type logMsg struct {
	Level     LogLevel
	msg       string
	funcName  string
	filename  string
	timestamp string

	lineOn int
}

// 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	LogLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		Level:       LogLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *logMsg, 50000),
	}
	err = f1.initFile() //按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return f1
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName+"-"+"info"+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err: %v\n", err)
		return err

	}
	fileObjErr, err := os.OpenFile(fullFileName+"-"+"err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open logerr file failed, err: %v\n", err)
		return err

	}
	// 日志文件都已经打开了
	f.fileObj = fileObj
	f.fileObjErr = fileObjErr

	// 开启后台的ground 写日志
	//for i := 0; i < 5; i++ {
	go f.WriteLogBackground()
	//	}
	return nil

}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)

}
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)

}
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)

}
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)

}
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.fileObjErr.Close()
}
