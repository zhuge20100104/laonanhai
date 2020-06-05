package main

import (
	"fmt"
	"net/http"
	"text/template"
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

	userMap := map[int]User{
		1: User{"祝二", "123", 28},
		2: User{"王五", "123", 18},
		3: User{"莉莉", "123", 20},
	}

	t.Execute(w, userMap)
}

func main() {
	http.HandleFunc("/info", InfoHandler)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		panic(fmt.Errorf("启动HTTPServer失败, err=%v", err))
	}
}
