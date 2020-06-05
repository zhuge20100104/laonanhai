package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// User 用户信息结构体
type User struct {
	UserName string
	Password string
	Age      int
}

// InfoHandler info请求的处理函数
func InfoHandler(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./info.html")
	if err != nil {
		panic(fmt.Errorf("Parse html file failed, err=%v", err))
	}

	user := User{
		UserName: "祝二",
		Password: "12345",
		Age:      28,
	}

	t.Execute(w, user)
}

func main() {
	http.HandleFunc("/info", InfoHandler)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		panic(fmt.Errorf("启动HTTPServer失败, err=%v", err))
	}
}
