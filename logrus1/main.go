package main

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	// 设置logrus打印JSON格式的日志
	log.SetFormatter(&logrus.JSONFormatter{})
	log.WithFields(logrus.Fields{
		"name": "祝二",
		"age":  18,
	}).Warn("这是一条Warning级别的日志!")
	log.Info("这是一条info级别的日志")
}
