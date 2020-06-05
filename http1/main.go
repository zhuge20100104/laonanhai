package main

import (
	"fmt"
	"net/http"
)

// SayHello 输出纯文本的处理函数
func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello，沙河"))
}

// SayHelloFile 文件版本的sayHello函数
func SayHelloFile(w http.ResponseWriter, r *http.Request) {
	res := ReadFileByBuf("./hello.txt")
	w.Write([]byte(res))
}

// SayHelloHTML HTML版本的sayHello函数
func SayHelloHTML(w http.ResponseWriter, r *http.Request) {
	res := ReadFileByBuf("./hello.html")
	w.Write([]byte(res))
}

func main() {
	http.HandleFunc("/", SayHello)
	http.HandleFunc("/file", SayHelloFile)
	http.HandleFunc("/html", SayHelloHTML)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(fmt.Errorf("启动 HTTP Server失败, err=%v", err))
	}
}
