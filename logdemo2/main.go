package main

import (
	log "github.com/zhuge20100104/laonanhai/logdemo2/logger"
)

var logger log.Logger

// 后台记日志程序
func main() {
	logger = log.NewFileLogger("debug", "test.log", ".")
	defer logger.Close()

	for {
		logger.Debug("这是一条Debug日志")
		name := "祝二"
		logger.Error("%v 是个好捧埂", name)
		logger.Fatal("这是一条Fatal级别的日志")
	}
}
