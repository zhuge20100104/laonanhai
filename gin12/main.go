package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User 用户对象
type User struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// LoginHandler /login接口处理函数
func LoginHandler(c *gin.Context) {
	// 处理GET请求，返回HTML
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	}
	// 处理POST请求，Bind JSON
	user := User{}
	// 绑定用户信息失败
	if err := c.ShouldBind(&user); err != nil {
		errMsg := fmt.Sprintf("获取用户信息失败, %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  errMsg,
		})
		return
	}
	// 绑定用户信息成功
	c.JSON(http.StatusOK, gin.H{
		"code":     0,
		"username": user.UserName,
		"password": user.Password,
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Any("/login", LoginHandler)
	r.Run(":8001")
}
