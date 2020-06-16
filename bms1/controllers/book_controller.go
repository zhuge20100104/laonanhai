package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhuge20100104/laonanhai/bms1/db"
	m "github.com/zhuge20100104/laonanhai/bms1/models"
)

// BookController 书籍处理控制器类
type BookController struct{}

// NewBookController 新建BookController对象并注册路由
func NewBookController(r *gin.Engine) *BookController {
	controller := new(BookController)
	r.GET("/book/list", controller.ListBook)
	r.GET("/book/new", controller.NewBook)
	r.POST("/book/new", controller.CreateBook)
	return controller
}

// ListBook 列出书籍列表GET请求
func (b *BookController) ListBook(c *gin.Context) {
	bookList, err := db.Db.QueryAllBook()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}

	// 查询成功，返回HTML页面

	c.HTML(http.StatusOK, "book_list.html", gin.H{
		"code": 0,
		"data": bookList,
	})
}

// NewBook 新建书本的GET请求
func (b *BookController) NewBook(c *gin.Context) {
	c.HTML(http.StatusOK, "new_book.html", nil)
}

// CreateBook 新建书本的POST请求
func (b *BookController) CreateBook(c *gin.Context) {
	titleStr := c.PostForm("title")
	priceStr := c.PostForm("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	// 验证价格是float64
	if err != nil {
		msg := "无效的价格参数"
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}
	book := &m.Book{
		Title: titleStr,
		Price: price,
	}

	err = db.Db.InsertBook(book)

	// 插入数据失败
	if err != nil {
		msg := "插入数据失败，请重试!"
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}
	// 插入成功，跳转到书籍列表页
	c.Redirect(http.StatusFound, "/book/list")
}
