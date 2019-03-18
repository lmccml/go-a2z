package demo

import (
	"fmt"
	"strconv"
)

func Rune_to_use() {
	//byte类型的底层类型是int8类型
	str := "hello 中国"
	//rune类型的底层类型是int32类型
	rune_str := []rune(str)
	fmt.Println("the length hello 中国 by byte array is " + strconv.Itoa(len(str)))
	//byte类型实际上是一个int8类型，int8适合表达ascii编码的字符，而int32可以表达更多的数，可以更容易的处理unicode字符
	fmt.Println("the length hello 中国 by rune is " + strconv.Itoa(len(rune_str)))
}
