package demo

import "fmt"

//内建关键字new的使用
func New_key_word_to_use() {
	//它只接受一个参数，这个参数是一个类型，返回一个指向该类型内存地址的指针。请注意它同时把分配的内存置为类型的零值。
	i := new(string)
	*i = "key word new's return value is function point!"
	fmt.Println(*i)
}
