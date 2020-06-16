package main

import (
	"net/http"
	"os"

	"github.com/zhuge20100104/laonanhai/logrus2/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// IndexHandler /index接口处理函数
func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Hello world",
	})
}

// InitLog 初始化log对象的函数
func InitLog() {
	log.Formatter = &logrus.JSONFormatter{}
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	utils.PanicErrorHand(err, "os.Open InitLog")
	log.Out = file
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Out
	log.Level = logrus.InfoLevel
}

func main() {
	InitLog()
	r := gin.Default()
	r.GET("/index", IndexHandler)
	r.Run(":8001")
}
