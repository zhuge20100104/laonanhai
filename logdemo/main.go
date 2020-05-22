package main

import "github.com/zhuge20100104/laonanhai/logdemo/logger"

func main() {
	log := logger.NewFileLogger(logger.DEBUG, ".", "log.log")
	log.Debug("%v", "呵呵")
	log.Info("%v", "你妹")
	log.Critical("%v", "可以IPO")
}
