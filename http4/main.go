package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// InfoHandler info请求的处理函数
func InfoHandler(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./info.html")
	if err != nil {
		panic(fmt.Errorf("Parse html file failed, err=%v", err))
	}
	t.Execute(w, "哈利波特历险记")
}

func main() {
	http.HandleFunc("/info", InfoHandler)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		panic(fmt.Errorf("启动HTTPServer失败, err=%v", err))
	}
}
