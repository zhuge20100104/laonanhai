package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginHandler 登陆处理函数
func LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"msg": "哈哈哈",
	})
}

// IndexHandler 主页处理函数
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"msg": "嘿嘿嘿",
	})
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/login", LoginHandler)
	engine.GET("/index", IndexHandler)
	engine.Run(":8001")
}
