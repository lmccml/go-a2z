package demo

import "fmt"

func Point_demo() {
	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"
	var z, w int
	fmt.Println(&z == &z, &z == &w, &z == nil) // "true false false"
	//在Go语言中，返回函数中局部变量的地址也是安全的。例如下面的代码，调用f函数时创建局部变量v，在局部变量
	// 地址被返回之后依然有效，因为指针p依然引用这个变量。
	fmt.Println(f() == f()) // "false"
}
func f() *int {
	v := 1
	return &v
}
