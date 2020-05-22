package logger

import (
	"fmt"
	"os"
	"time"
)

type FileLogger struct {
	level       int
	logFilePath string
	logFileName string
	logFile     *os.File
}

func NewFileLogger(level int, logFilePath, logFileName string) *FileLogger {
	fl := &FileLogger{
		level:       level,
		logFilePath: logFilePath,
		logFileName: logFileName,
	}

	fl.InitFileLogger()
	return fl
}

func (fl *FileLogger) InitFileLogger() {
	filePath := fmt.Sprintf("%s/%s", fl.logFilePath, fl.logFileName)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("打开文件%v失败, err=%v\n", filePath, err.Error()))
	}
	fl.logFile = file
}

func (fl *FileLogger) log(format string, args ...interface{}) {
	fileName, funcName, line := getCallerInfo()
	nowStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	format = fmt.Sprintf("%s [%s] [%s:%s] [%d] %s", nowStr, getLevelStr(fl.level),
		fileName, funcName, line, format)
	fmt.Fprintf(fl.logFile, format, args...)
	fmt.Fprintln(fl.logFile)
}

func (fl *FileLogger) Debug(format string, args ...interface{}) {
	// 如果当前级别大于 DEBUG,只写DEBUG以上的日志
	if fl.level > DEBUG {
		return
	}
	fl.log(format, args)
}

func (fl *FileLogger) Trace(format string, args ...interface{}) {
	// 如果当前级别大于 TRACE,只写TRACE以上的日志
	if fl.level > TRACE {
		return
	}
	fl.log(format, args)
}

func (fl *FileLogger) Info(format string, args ...interface{}) {
	// 如果当前级别大于 INFO,只写INFO以上的日志
	if fl.level > INFO {
		return
	}
	fl.log(format, args)
}

func (fl *FileLogger) Warn(format string, args ...interface{}) {
	// 如果当前级别大于 WARN,只写WARN以上的日志
	if fl.level > WARN {
		return
	}
	fl.log(format, args)
}

func (fl *FileLogger) Error(format string, args ...interface{}) {
	// 如果当前级别大于 ERROR,只写ERROR以上的日志
	if fl.level > ERROR {
		return
	}
	fl.log(format, args)
}

func (fl *FileLogger) Critical(format string, args ...interface{}) {
	// 如果当前级别大于 CRITICAL,只写CRITICAL以上的日志
	if fl.level > CRITICAL {
		return
	}
	fl.log(format, args)
}
