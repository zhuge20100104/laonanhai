package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhuge20100104/laonanhai/cookiemiddleware/utils"
)

// ToPath 全局TOPATH对象
var ToPath string

// User 用户信息结构体
type User struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// CookieMiddleware Cookie版本的login中间件
func CookieMiddleware(c *gin.Context) {
	username, err := c.Cookie("username")

	// 获取username信息失败
	if err != nil {
		// 设置下一个跳转的页面，好让下次直接跳转到页面
		toPath := c.FullPath()
		utils.Data.Set("next", toPath)
		c.Set("next", toPath)
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.Set("username", username)
	c.Next()
}

// LoginHandler /login登陆处理函数
func LoginHandler(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	}

	u := User{}
	if err := c.ShouldBind(&u); err != nil {
		errMsg := fmt.Sprintf("用户名和密码不能为空%v", err)
		c.HTML(http.StatusOK, "login.html", gin.H{
			"err": errMsg,
		})
		return
	}

	// 用户名和密码正确，跳转到指定页面
	if u.UserName == "祝二" && u.Password == "123" {
		toPath := utils.Data.GetString("next")

		if toPath == "" {
			toPath = "/index"
		}

		fmt.Println("toPath:", toPath)
		c.SetCookie("username", u.UserName, 20, "/", "127.0.0.1", false, true)
		c.Redirect(http.StatusFound, toPath)
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{
		"err": "用户名和密码错误",
	})
}

// IndexHandler /index页面的处理函数
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// VipHandler /vip页面的处理函数
func VipHandler(c *gin.Context) {
	val, exists := c.Get("username")
	if !exists {
		fmt.Println("您还没有登陆!!")
		return
	}
	username, ok := val.(string)
	if !ok {
		fmt.Println("您还没有登陆!!")
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "vip.html", gin.H{
		"username": username,
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Any("/login", LoginHandler)

	r.Use(CookieMiddleware)
	r.GET("/index", IndexHandler)
	r.GET("/vip", VipHandler)
	r.Run(":8001")
}
