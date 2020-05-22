package logger

import (
	"path"
	"runtime"
)

const (
	DEBUG = iota
	TRACE
	INFO
	WARN
	ERROR
	CRITICAL
)

// 获取log日志级别字符串
func getLevelStr(level int) string {
	switch level {
	case DEBUG:
		{
			return "DEBUG"
		}
	case TRACE:
		{
			return "TRACE"
		}
	case INFO:
		{
			return "INFO"
		}
	case WARN:
		{
			return "WARN"
		}
	case ERROR:
		{
			return "ERROR"
		}
	case CRITICAL:
		{
			return "CRITICAL"
		}
	default:
		{
			return "DEBUG"
		}
	}
}

func getCallerInfo() (fileName, funcName string, line int) {
	pc, fileName, line, ok := runtime.Caller(3)

	// 获取当前调用堆栈信息失败，直接返回默认值
	if !ok {
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	fileName = path.Base(fileName)
	return
}
