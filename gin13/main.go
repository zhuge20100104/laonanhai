package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Foo Foo结构体
type Foo struct {
	Foo string `json:"foo" binding:"required"`
}

// Bar Bar结构体
type Bar struct {
	Bar string `json:"bar" binding:"required"`
}

// IndexHandler /index页面的处理函数
func IndexHandler(c *gin.Context) {
	foo := Foo{}
	bar := Bar{}

	if err := c.ShouldBindBodyWith(&foo, binding.JSON); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "你传了一个Foo对象给我",
		})
	} else if err := c.ShouldBindBodyWith(&bar, binding.JSON); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "你传了一个Bar对象给我",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "你传了一个不知道什么对象给我",
		})
	}
}

// 测试ShouldBindBodyWith函数
func main() {
	r := gin.Default()
	r.POST("/index", IndexHandler)
	r.Run(":8001")
}
