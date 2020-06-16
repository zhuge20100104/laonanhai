package db

import (
	// 加载MYSQL数据库驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	m "github.com/zhuge20100104/laonanhai/bms1/models"
	"github.com/zhuge20100104/laonanhai/bms1/utils"
)

var (
	// Db 全局的数据库对象
	Db *DB
)

// DB DB对象结构体
type DB struct {
	DB *sqlx.DB
}

func init() {
	Db = new(DB)
	dsn := "root:root@tcp(127.0.0.1:3306)/books"
	db, err := sqlx.Open("mysql", dsn)
	// 连接数据库失败，直接Panic
	utils.PanicErrorHand(err, "sqlx.Open Connect DB")
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	Db.DB = db
}

// QueryAllBook 查询所有书籍的函数
func (d *DB) QueryAllBook() (bookList []*m.Book, err error) {
	sqlStr := "select id, title, price from book;"
	err = d.DB.Select(&bookList, sqlStr)
	if err != nil {
		utils.DefErrorHand(err, "d.DB.Select")
		return
	}
	return
}

// InsertBook 插入书籍
func (d *DB) InsertBook(book *m.Book) (err error) {
	sqlStr := "insert into book(title, price) values(?,?)"
	_, err = d.DB.Exec(sqlStr, book.Title, book.Price)
	if err != nil {
		utils.DefErrorHand(err, "db.Exec Insert")
		return
	}
	return
}

// CloseDB 关闭数据库连接的函数
func (d *DB) CloseDB() {
	d.DB.Close()
}
