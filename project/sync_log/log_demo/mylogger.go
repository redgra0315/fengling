package mylogger

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type Logger interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

const (
	UNKNOWN LogLevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "falat":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

func GetLogSrting(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INTO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}
func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Print("runtime,Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return funcName, fileName, lineNo
}

// NewLog ...
func Newlog(levelStr string) ConcoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConcoleLogger{
		Level: level,
	}
}
func (f *FileLogger) enable(logger LogLevel) bool {
	return f.Level <= logger
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err：%v\n", err)
		return false
	}
	// 如果当前文件的大小,大于等于日志文件的值，就应该返回True
	return fileInfo.Size() > f.maxFileSize
}

func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	//f.fileObj.Close()
	// 备份日志
	nowSrt := time.Now().Format("20060102150405")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%s\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())
	NewLogName := fmt.Sprintf("%s%s-%v-%s.log", f.filePath, f.fileName, f.Level, nowSrt)
	file.Close()
	os.Rename(logName, NewLogName)
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err!%v\n", err)
		return nil, err
	}
	return fileObj, nil
}

func (f *FileLogger) WriteLogBackground() {
	for {
		if f.checkSize(f.fileObj) {
			//	切割文件
			// 关闭日志文件
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		select {
		case logTmp := <-f.logChan:
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", logTmp.timestamp, GetLogSrting(logTmp.Level),
				logTmp.funcName, logTmp.filename, logTmp.lineOn, logTmp.msg)
			//now.Format("2006-01-02 15:04:05")
			fmt.Fprintf(f.fileObj, logInfo)
			if logTmp.Level >= ERROR {
				if f.checkSize(f.fileObjErr) {
					newFileErr, err := f.splitFile(f.fileObjErr)
					if err != nil {
						return
					}
					f.fileObj = newFileErr
				}
				fmt.Fprintf(f.fileObjErr, logInfo)
			}
		default:
			// 取不到休息500毫秒
			time.Sleep(time.Millisecond * 500)
		}
	}

}
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)

		// 先把日志发送通道

		// 造一个logmsg 对象
		logTmp := &logMsg{
			Level:     lv,
			msg:       msg,
			funcName:  funcName,
			filename:  fileName,
			timestamp: now.Format("2006-01-02 15:04:05"),
			lineOn:    lineNo,
		}
		select {

		case f.logChan <- logTmp:
		default:
			//	 丢弃日志

		}

	}
}
