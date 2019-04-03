package demo

import "fmt"

//神奇的魔法-替换变量 alias 元组赋值
func Replace_magic() {
	i := "old_str"
	j := "new_str"
	k := "third_str"
	i, j, k = k, i, j
	fmt.Println("this is a easy way to replace " + i + " " + j + " " + k)

}

//求最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

//计算斐波纳契数列（Fibonacci）的第N个数
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
