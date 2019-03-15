package demo

import (
	"fmt"
	"strconv"
)

//关键词make的使用示例
func Make_key_word_to_use() {
	//1.用于slice
	slice_demo := make([]string, 3)
	slice_demo[1] = "this is the slice_demo slice second contents"
	fmt.Println(slice_demo[1])
	//2.用于map
	map_demo := make(map[string]int, 5)
	map_demo["wuhan"] = 1
	map_demo["jinhua"] = 2
	for k, v := range map_demo {
		fmt.Println("the city " + k + " in your heart number is " + strconv.Itoa(v))
	}
	//3.用于chan
	chan_demo := make(chan int, 1)
	fmt.Println("channel demo is running!")
	chan_demo <- 1
	fmt.Println("channel demo is sending data!")
	receiver := <-chan_demo
	fmt.Println("channel demo's data is " + strconv.Itoa(receiver))
	fmt.Println("channel demo is closed!")
}
