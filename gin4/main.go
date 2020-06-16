package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User 返回给前端的User对象
type User struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

// HelloHandler 处理Hello路由的函数
func HelloHandler(c *gin.Context) {
	user := User{
		Name: "张三",
		Pwd:  "12345",
	}
	c.JSON(http.StatusOK, &user)
}

func main() {
	engine := gin.Default()
	engine.GET("/hello", HelloHandler)
	engine.Run(":8001")
}
