package logger

import (
	"fmt"
	"os"
	"time"
)

// 往终端输出日志的类
type ConsoleLogger struct {
	level Level
}

// 终端日志类构造函数
func NewConsoleLogger(levelStr string) *ConsoleLogger {
	logLevel := parseLogLevel(levelStr)
	cl := &ConsoleLogger{
		level: logLevel,
	}
	return cl
}

func (c *ConsoleLogger) log(level Level, format string, args ...interface{}) {
	// 设置的级别大于当前级别，不输出
	if c.level > level {
		return
	}

	msg := fmt.Sprintf(format, args...)
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, line, funcName := GetCallerInfo(3)
	levelStr := getLevelStr(level)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s", nowStr, fileName, line, funcName,
		levelStr, msg)
	fmt.Fprintln(os.Stdout, logMsg)
}

func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	c.log(DebugLevel, format, args...)
}

func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	c.log(InfoLevel, format, args...)
}

func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	c.log(WarningLevel, format, args...)
}

func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	c.log(ErrorLevel, format, args...)
}

func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	c.log(FatalLevel, format, args...)
}

func (c *ConsoleLogger) Close() {

}
