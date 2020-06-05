package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// ReadFileByBuf 通过缓冲区的形式读取文件的函数
func ReadFileByBuf(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Errorf("打开文件失败, err=%v", err))
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 10*1024)
	buf := make([]byte, 10*1024)
	res := ""
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(fmt.Errorf("读取文件错误, err=%v", err))
		}
		res += string(buf[:n])
	}
	return res
}
