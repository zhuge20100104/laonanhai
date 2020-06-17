package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User 用户信息结构体变量
type User struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// LoginHandler 登陆处理函数
func LoginHandler(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	}

	var u User
	// 获取用户信息失败
	if err := c.ShouldBind(&u); err != nil {
		fmt.Println("用户名和密码不能为空!")
		errMsg := fmt.Sprintf("用户名或密码不能为空 %v", err)
		c.HTML(http.StatusOK, "login.html", gin.H{
			"err": errMsg,
		})
		return
	}

	// 用户名和密码正确
	if u.UserName == "祝二" && u.Password == "123" {
		// 设置cookie，跳转到index页面
		c.SetCookie("username", u.UserName, 20, "/", "127.0.0.1", false, true)
		c.Redirect(http.StatusFound, "/index")
		return
	}
	// 用户名和密码错误
	c.HTML(http.StatusOK, "login.html", gin.H{
		"err": "用户名或密码错误",
	})
}

// IndexHandler index页面处理函数
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// HomeHandler Home页面处理函数
func HomeHandler(c *gin.Context) {
	username, err := c.Cookie("username")
	// 若未找到cookie，跳转到 login页面
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"username": username,
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Any("/login", LoginHandler)
	r.GET("/index", IndexHandler)
	r.GET("/home", HomeHandler)
	r.Run(":8001")
}
