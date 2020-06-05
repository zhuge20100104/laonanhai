package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	level       Level
	FileName    string
	FilePath    string
	File        *os.File      //记录普通日志的字段
	ErrFile     *os.File      //记录错误日志的字段
	maxSize     int64         //日志文件按此大小进行切分
	logDataChan chan *LogData // LogData通道，用于后台记日志
}

func NewFileLogger(levelStr, fileName, filePath string) *FileLogger {
	level := parseLogLevel(levelStr)
	fl := &FileLogger{
		level:       level,
		FileName:    fileName,
		FilePath:    filePath,
		maxSize:     10 * 1024 * 1024,           // 10M切分
		logDataChan: make(chan *LogData, 50000), // 50000日志缓冲区
	}
	// 打开标准日志文件和错误日志文件
	fl.initFile()
	return fl
}

func (f *FileLogger) initFile() {
	logName := path.Join(f.FilePath, f.FileName)
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败, %v\n", logName, err))
	}
	f.File = fileObj
	errLogName := fmt.Sprintf("%s.err", logName)
	errFileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(fmt.Errorf("打开错误日志文件%s失败, %v\n", logName, err))
	}
	f.ErrFile = errFileObj
	go f.writeLogBackground() // 启动后台写日志Go routine
}

// 后台写日志的程序
func (f *FileLogger) writeLogBackground() {
	for logData := range f.logDataChan {

		logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s", logData.TimeStr, logData.FileName,
			logData.LineNo, logData.FuncName,
			logData.LogLevel, logData.Message)

		if f.checkSplit(f.File) {
			f.File = f.splitLogFile(f.File)
		}
		fmt.Fprintln(f.File, logMsg)

		// 当前级别大于error级别，记录到error日志文件
		level := parseLogLevel(logData.LogLevel)
		if level >= ErrorLevel {
			if f.checkSplit(f.ErrFile) {
				f.ErrFile = f.splitLogFile(f.ErrFile)
			}
			fmt.Fprintln(f.ErrFile, logMsg)
		}
	}
}

// 检查当前文件是否需要split
func (f *FileLogger) checkSplit(file *os.File) bool {
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	return fileSize >= f.maxSize
}

// 切分文件的函数
func (f *FileLogger) splitLogFile(file *os.File) *os.File {
	fileName := file.Name()
	backupName := fmt.Sprintf("%v_%v.back", fileName, time.Now().Unix())
	file.Close()
	// 备份原来的文件
	os.Rename(fileName, backupName)
	// 新建一个文件
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(fmt.Errorf("打开日志文件失败"))
	}
	return fileObj
}

func (f *FileLogger) log(level Level, format string, args ...interface{}) {
	// 设置的级别大于当前级别，不输出
	if f.level > level {
		return
	}

	msg := fmt.Sprintf(format, args...)
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, line, funcName := GetCallerInfo(3)
	levelStr := getLevelStr(level)

	logData := &LogData{
		Message:  msg,
		LogLevel: levelStr,
		LineNo:   line,
		TimeStr:  nowStr,
		FuncName: funcName,
		FileName: fileName,
	}

	// 这样写通道满了以后就会阻塞，不是特别好，添加一个select语句
	// f.logDataChan <- logData

	select {
	case f.logDataChan <- logData: // 能够写入日志数据就直接写入
	default:
		<-f.logDataChan          //通道满了以后，取出最早的一行记录
		f.logDataChan <- logData //写入最新记录
	}
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DebugLevel, format, args...)
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)
}

func (f *FileLogger) Warn(format string, args ...interface{}) {
	f.log(WarningLevel, format, args...)
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.log(FatalLevel, format, args...)
}

func (f *FileLogger) Close() {
	f.File.Close()
	f.ErrFile.Close()
	close(f.logDataChan)
}
