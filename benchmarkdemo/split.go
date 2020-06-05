package split

import (
	"strings"
)

/**
* 定义一个分割字符串的函数
* 输入 a:b:c 和 : 输出 [a,b,c]
 */
func Split(str, sep string) []string {
	// 预先分配好内存，不用动态分配，查看优化效果
	count := strings.Count(str, sep)
	result := make([]string, 0, count+1)
	idx := strings.Index(str, sep)

	for idx >= 0 {
		result = append(result, str[:idx])
		str = str[idx+len(sep):]
		idx = strings.Index(str, sep)
	}
	result = append(result, str)
	return result
}

/**
* 定义一个分割字符串的函数
* 输入 a:b:c 和 : 输出 [a,b,c]
 */
func Split1(str, sep string) (result []string) {
	idx := strings.Index(str, sep)

	for idx >= 0 {
		result = append(result, str[:idx])
		str = str[idx+len(sep):]
		idx = strings.Index(str, sep)
	}
	result = append(result, str)
	return
}
