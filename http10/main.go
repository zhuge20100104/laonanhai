package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// InfoHandler Info接口的回调处理函数
func InfoHandler(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("info.html", "ul.html")
	if err != nil {
		fmt.Printf("Parse HTML file failed, err=%v\n", err)
		return
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/info", InfoHandler)
	http.ListenAndServe(":8001", nil)
}
