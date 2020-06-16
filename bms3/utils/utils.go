package utils

import "fmt"

// handleErrorFunc 错误处理函数闭包
// shouldPanic 是否引发恐慌错误
// err 传入的错误信息
// msg 传入的msg信息
func handleErrorFunc(shouldPanic bool) func(error, string) {
	return func(err error, msg string) {
		if err != nil {
			fmt.Printf("[%s] --- %v\n", msg, err)
		}

		if err != nil && shouldPanic {
			panic(err)
		}
	}
}

// PanicErrorHand 处理错误时同时panic
var PanicErrorHand = handleErrorFunc(true)

// DefErrorHand 处理错误时不panic
var DefErrorHand = handleErrorFunc(false)
