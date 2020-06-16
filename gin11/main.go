package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostsHandler 博客按年月日获取的处理函数
func PostsHandler(c *gin.Context) {
	year := c.Param("year")
	mon := c.Param("month")
	day := c.Param("day")

	c.JSON(http.StatusOK, gin.H{
		"year":  year,
		"month": mon,
		"day":   day,
	})
}

// 演示 URL 参数获取
func main() {
	r := gin.Default()
	r.GET("/posts/:year/:month/:day", PostsHandler)
	r.Run(":8001")
}
