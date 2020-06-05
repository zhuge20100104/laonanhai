package fib

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2) // 递归调用，获取Fibnacci数列的值
}
