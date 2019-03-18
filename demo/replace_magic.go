package demo

import "fmt"

func Replace_magic() {
	i := "old_str"
	j := "new_str"
	i, j = j, i
	fmt.Println("this is a easy way to replace " + i + " " + j)
}
