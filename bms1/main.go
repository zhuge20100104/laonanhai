package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuge20100104/laonanhai/bms1/controllers"
	"github.com/zhuge20100104/laonanhai/bms1/db"
)

func main() {
	// 退出主函数时关闭数据库连接
	defer db.Db.CloseDB()
	r := gin.Default()
	// 加载HTML模板目录
	r.LoadHTMLGlob("templates/*")
	InitRoutes(r)
	r.Run(":8001")
}

// InitRoutes 初始化路由组的函数
func InitRoutes(r *gin.Engine) {
	controllers.NewBookController(r)
}
