package demo

import "fmt"

//逗号的妙用
func Comma_magic() {
	a := 519
	b := 1
	//译注：函数的右小括弧也可以另起一行缩进，同时为了防止编译器在行尾自动插入分号而导致的编译错误，可以在
	//末尾的参数变量后面显式插入逗号。像下面这样：
	c := test_add(
		a,
		b, // 最后插入的逗号不会导致编译错误，这是Go编译器的一个特性
	)
	fmt.Println(c)
}
func test_add(a int, b int) int {
	return a + b
}
