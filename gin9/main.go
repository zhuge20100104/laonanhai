package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexHandler 主页面GET处理函数
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

// UploadHandler 上传文件处理函数
func UploadHandler(c *gin.Context) {
	fileObj, err := c.FormFile("filename")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}

	fName := fileObj.Filename
	saveFileName := fmt.Sprintf("./%v", fName)
	err = c.SaveUploadedFile(fileObj, saveFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "上传文件成功",
	})

}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/upload", IndexHandler)
	r.POST("/upload", UploadHandler)
	r.Run(":8001")
}
