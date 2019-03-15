package demo

import (
	"fmt"
)

func Go_tips() {
	tip_maps := make(map[string]string)
	tip_maps["tip_1 "] = "可以i++,不能--i!"

	for k, v := range tip_maps {
		fmt.Println(k + " is " + v)
	}
}
