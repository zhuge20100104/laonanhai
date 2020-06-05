package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

// InfoHandler info请求的处理函数
func InfoHandler(w http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadFile("./info.html")
	if err != nil {
		panic(fmt.Errorf("读取本地文件失败 %v", err))
	}
	dataStr := string(data)
	num := rand.Intn(10)
	if num > 5 {
		dataStr = strings.Replace(dataStr, "{{ooxx}}", "对子哈特一枚", -1)
	} else {
		dataStr = strings.Replace(dataStr, "{{ooxx}}", "好奇铂金装", -1)
	}
	w.Write([]byte(dataStr))
}

func main() {
	http.HandleFunc("/info", InfoHandler)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		panic(fmt.Errorf("启动HTTPServer失败, err=%v", err))
	}
}
