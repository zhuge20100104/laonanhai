package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// XMLHandler XML处理函数
func XMLHandler(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{
		"msg": "XML",
	})
}

// YAMLHandler YAML处理函数
func YAMLHandler(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{
		"msg":    "YAML",
		"status": http.StatusOK,
	})
}

func main() {
	engine := gin.Default()
	engine.GET("/xml", XMLHandler)
	engine.GET("/yaml", YAMLHandler)
	engine.Run(":8001")
}
