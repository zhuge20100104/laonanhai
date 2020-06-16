package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// FormHandler Form处理函数
func FormHandler(c *gin.Context) {
	name := c.DefaultPostForm("name", "张三")
	city := c.DefaultPostForm("city", "雄安")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"city": city,
	})
}

// PathHandler Path处理函数
func PathHandler(c *gin.Context) {
	action := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"action": action,
	})
}

func main() {
	engine := gin.Default()
	engine.POST("/form", FormHandler)
	engine.GET("/book/:action", PathHandler)
	engine.Run(":8001")
}
