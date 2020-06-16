package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShopIndexHandler /shopping/index页面的处理函数
func ShopIndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "/shopping/index",
	})
}

// ShopHomeHandler /shopping/home页面的处理函数
func ShopHomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "/shopping/home",
	})
}

// LiveIndexHandler /live/index页面的处理函数
func LiveIndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "/live/index",
	})
}

// LiveHomeHandler /live/home页面的处理函数
func LiveHomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "/live/home",
	})
}

func main() {
	r := gin.Default()
	shoppingG := r.Group("/shopping")
	{
		shoppingG.GET("/index", ShopIndexHandler)
		shoppingG.GET("/home", ShopHomeHandler)
	}

	liveG := r.Group("/live")
	{
		liveG.GET("/index", LiveIndexHandler)
		liveG.GET("/home", LiveHomeHandler)
	}
	r.Run(":8001")
}
