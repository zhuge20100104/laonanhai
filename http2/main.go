package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// IndexHandler Index页面的处理函数
func IndexHandler(w http.ResponseWriter, req *http.Request) {
	// 如果是GET请求，直接返回 index页面
	if req.Method == "GET" {
		data, err := ioutil.ReadFile("./index.html")
		if err != nil {
			panic(fmt.Errorf("读取本地 HTML文件失败, err=%v", err))
		}
		w.Write(data)
		// 如果是POST请求，在这里处理POST请求的逻辑
	} else {
		err := req.ParseForm()
		if err != nil {
			panic(fmt.Errorf("解析Form参数失败, err = %v", err))
		}
		userName := req.Form.Get("username")
		pwd := req.Form.Get("pwd")
		ret := fmt.Sprintf(`用户名: %v
密码: %v`, userName, pwd)

		w.Write([]byte(ret))
	}

}

func main() {
	http.HandleFunc("/index", IndexHandler)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		panic(fmt.Errorf("启动HTTP Server失败, err= %v", err))
	}
}
