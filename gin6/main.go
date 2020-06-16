package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// QueryStringHandler QueryString处理器函数
func QueryStringHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "祝二")
	city := c.DefaultQuery("city", "雄安")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"city": city,
	})
}

func main() {
	engine := gin.Default()
	engine.GET("/querystring", QueryStringHandler)
	engine.Run(":8001")
}
