package demo

import (
	"fmt"
	"strconv"
)

//关于go的一些小提示
func Go_tips() {
	a := make([]string, 1000)
	a[1] = "可以i++,不能--i!"
	a[2] = "for循环的这三个部分每个都可以省略"
	a[3] = "在unicode中，一个中文占两个字节，utf-8中一个中文占三个字节，golang默认的编码是utf-8编码，因此默认一个中文占三个字节，但是golang中的字符串底层实际上是一个byte数组"
	for k, v := range a {
		if k == 0 || v == "" {
			continue
		}
		fmt.Println("tip" + strconv.Itoa(k) + " is " + v)
	}
}
