package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexHandler 主页处理函数
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"msg": "嘿嘿嘿",
	})
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.Static("/static", "./statics")
	engine.GET("/index", IndexHandler)
	engine.Run(":8001")
}
