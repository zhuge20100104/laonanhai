package models

// Book 书籍对象结构体
type Book struct {
	ID    int64   `db:"id" form:"id"`
	Title string  `db:"title" form:"title"`
	Price float64 `db:"price" form:"price"`
}
