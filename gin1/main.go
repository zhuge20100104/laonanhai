package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexHandler Index页面的处理函数
func IndexHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "这是Index页面",
	})
}

func main() {
	router := gin.Default()
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Hello, 沙河!",
		})
	})

	router.GET("/index", IndexHandler)
	router.Run(":8001")
}
