package logger

import (
	"path"
	"runtime"
	"strings"
)

// utils包一般用于存放一些公用的工具函数

//获取调用者信息的函数
func GetCallerInfo(skip int) (fileName string, line int, funcName string) {
	pc, fileName, line, ok := runtime.Caller(skip)
	// 获取Caller信息失败，直接return
	if !ok {
		return
	}
	// 从文件全路径中剥离出文件名
	fileName = path.Base(fileName)
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	return
}

// 根据传进来的level获取对应的level字符串
func getLevelStr(level Level) string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarningLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

// 根据用户传进来的Level字符串，返回对应的Level级别
func parseLogLevel(levelStr string) Level {
	levelStr = strings.ToLower(levelStr)
	switch levelStr {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarningLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	default:
		return DebugLevel
	}
}
