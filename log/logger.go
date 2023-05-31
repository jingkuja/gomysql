package log

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

type LogLevel uint16

// 日志级别
const (
	UNKNOW LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNIG
	ERROR
	FATAL
)

// 解析日志
func paraLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {

	case "debug":
		return DEBUG, nil

	case "tarce":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warnig":
		return WARNIG, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOW, err
	}
}

// 获取日志的字符串格式
func getLogStr(level LogLevel) string {

	switch level {

	case DEBUG:
		return "debug"

	case TRACE:
		return "tarce"
	case INFO:
		return "info"
	case WARNIG:
		return "warnig"
	case ERROR:
		return "error"
	case FATAL:
		return "fatal"
	default:
		return "unknow"
	}
}

// 定义日志的结构体
type FileLogger struct {
	Level    LogLevel
	filePath string
	fileName string
	//要打开和写入的文件，一个日志文件一个错误日志文件
	fileObj     *os.File
	errfileObj  *os.File
	maxFileSize int64
}

// 构造函数
func NewFlieLogger(LeveStr, fp, fn string, size int64) *FileLogger {

	level, err := paraLogLevel(LeveStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		Level:       level,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: size,
	}
	err = f1.initFile()
	if err != nil {
		panic(err)
	}
	return f1
}

// 初始化要打开和写入的日志文件的操作
func (f *FileLogger) initFile() error {
	join := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(join, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log fail ,err: %v\n", err)
		return err
	}

	errFileObj, err := os.OpenFile(join+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log fail ,err: %v\n", err)
		return err
	}
	//日志文件都打开
	f.fileObj = fileObj
	f.errfileObj = errFileObj
	return nil

}

// 判断级别
func (l FileLogger) enable(logLevel LogLevel) bool {
	return l.Level > logLevel
}

// 打印日志操作
func (f *FileLogger) Log(leve LogLevel, msg string) {
	now := time.Now()
	if f.enable(leve) {

		fmt.Fprintf(f.fileObj, "[%s] [%s] %s", now.Format("2006-01-02 15:04:05"), getLogStr(leve), msg)
	}

	if leve > ERROR {

		fmt.Fprintf(f.errfileObj, "[%s] [%s] %s", now.Format("2006-01-02 15:04:05"), getLogStr(leve), msg)
	}
}

func (l FileLogger) Debug(msg string, a ...interface{}) {
	msg = fmt.Sprint(msg, a)
	if l.enable(DEBUG) {
		l.Log(DEBUG, msg)
	}

}

func (l FileLogger) Info(msg string, a ...interface{}) {
	msg = fmt.Sprint(msg, a)
	if l.enable(WARNIG) {
		l.Log(WARNIG, msg)
	}

}

func (l FileLogger) Warning(msg string, a ...interface{}) {
	msg = fmt.Sprint(msg, a)
	if l.enable(WARNIG) {
		l.Log(WARNIG, msg)
	}

}

func (l FileLogger) Error(msg string, a ...interface{}) {
	msg = fmt.Sprint(msg, a)
	if l.enable(ERROR) {
		l.Log(ERROR, msg)
	}

}

func (l FileLogger) Fatal(msg string, a ...interface{}) {
	msg = fmt.Sprint(msg, a)
	if l.enable(FATAL) {
		l.Log(FATAL, msg)
	}

}

func (f *FileLogger) Colse() {
	f.fileObj.Close()
	f.errfileObj.Close()
}
