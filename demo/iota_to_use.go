package demo

import (
	"fmt"
	"strconv"
)

func Iota_to_use() {
	const (
		one = iota + 1
		two
		three
		_
		five
	)
	fmt.Println("this is first usage about iota " + strconv.Itoa(one) + strconv.Itoa(two) + strconv.Itoa(three) + strconv.Itoa(five))

}
