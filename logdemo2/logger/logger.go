package logger

type Level uint16

const (
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

// Logger接口类，用于规范 FileLogger和ConsoleLogger的行为
type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}

// Log数据类
type LogData struct {
	Message  string
	LogLevel string
	LineNo   int
	TimeStr  string
	FuncName string
	FileName string
}
