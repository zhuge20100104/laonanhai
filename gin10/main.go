package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CostMiddleware Cost中间件
func CostMiddleware(c *gin.Context) {
	start := time.Now()
	// 设置其他 Handler都可以用的key
	c.Set("key", "Key [Start]")
	c.Next()
	cost := time.Since(start)
	fmt.Println("Cost", cost)
}

// IndexHandler /shop/index处理函数
func IndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "/shop/index",
	})
}

// HomeHandler /shop/home处理函数
func HomeHandler(c *gin.Context) {
	time.Sleep(1 * time.Second)
	val, _ := c.Get("key")
	valStr := val.(string)

	fmt.Println("key: ", valStr)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "/shop/home",
	})
}

func main() {
	r := gin.Default()
	// 使用Cost计算中间件
	r.Use(CostMiddleware)

	shopG := r.Group("/shop")
	{
		shopG.GET("/index", IndexHandler)
		shopG.GET("/home", HomeHandler)
	}

	r.Run(":8001")
}
