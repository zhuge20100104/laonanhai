package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhuge20100104/laonanhai/bms3/db"
	m "github.com/zhuge20100104/laonanhai/bms3/models"
	"github.com/zhuge20100104/laonanhai/bms3/utils"
)

// BookController 书籍处理控制器类
type BookController struct{}

// NewBookController 新建BookController对象并注册路由
func NewBookController(r *gin.Engine) *BookController {
	controller := new(BookController)
	r.GET("/book/list", controller.ListBook)
	r.GET("/book/new", controller.NewBook)
	r.POST("/book/new", controller.CreateBook)
	r.GET("/book/delete", controller.DeleteBook)
	r.Any("/book/edit", controller.EditBook)
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

// DeleteBook 删除书籍的函数
func (b *BookController) DeleteBook(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.DefErrorHand(err, "strconv.Atoi ID")
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "不合规的ID字符串",
		})
		return
	}

	err = db.Db.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "删除书籍信息失败",
		})
		return
	}

	// 删除成功，跳转到bookList
	c.Redirect(http.StatusFound, "/book/list")
}

func (b *BookController) editGET(c *gin.Context) {
	// 拿到前端传过来的ID
	idStr := c.Query("id")
	if len(idStr) == 0 {
		utils.DefErrorHand(errors.New("id字符串长度不能为0"), "c.Query Edit Book")
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "传入的id不能为空",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.DefErrorHand(err, "strconv.Atoi Edit Book")
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "Id只能为整数",
		})
		return
	}

	// 从数据库中查询出对应数据
	book, err := db.Db.QueryBookByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "不存在的书籍ID",
		})
		return
	}
	c.HTML(http.StatusOK, "edit_book.html", book)
}

func (b *BookController) editPost(c *gin.Context) {
	// 获取提交的书籍信息
	book := m.Book{}
	if err := c.ShouldBind(&book); err != nil {
		utils.DefErrorHand(err, "c.ShouldBind editPost")
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "获取书籍信息失败",
		})
		return
	}

	// 更新书籍信息到数据库
	if err := db.Db.UpdateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  "更新数据库失败",
		})
		return
	}

	// 更新成功，跳转到 bookList页面
	c.Redirect(http.StatusFound, "/book/list")
}

// EditBook 编辑书籍信息的处理函数
func (b *BookController) EditBook(c *gin.Context) {
	// 处理GET请求
	if c.Request.Method == "GET" {
		b.editGET(c)
		// 返回给前端进行渲染
		return
	}

	b.editPost(c)
}
