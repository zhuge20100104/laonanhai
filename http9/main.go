package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
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
	htmlBytes, err := ioutil.ReadFile("./info.html")
	if err != nil {
		fmt.Println("Read HTML file err, err:=", err)
		return
	}

	kuaFunc := func(arg string) (string, error) {
		return arg + "真帅", nil
	}

	t, err := template.New("Info").Funcs(template.FuncMap{"kua": kuaFunc}).Parse(string(htmlBytes))
	if err != nil {
		fmt.Println("Parse HTML file failed, err=", err)
		return
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
