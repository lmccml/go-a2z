package main

import (
	"fmt"
	"go-test/demo"
	"os"
)

func main() {
	//爬虫小栗子
	//demo.Get_nba_live_info();
	//关键词make的使用示例
	demo.Make_key_word_to_use()
	//关键词new的使用示例
	demo.New_key_word_to_use()
	//打印命令行参数
	fmt.Println("go run main.go parameter----打印命令行参数(空格分开):" + os.Args[0])
	//获取go小技巧
	demo.Go_tips()
	//iota的妙用
	demo.Iota_to_use()
	//rune和byte表示字符串
	demo.Rune_to_use()
	//快速替换变量
	demo.Replace_magic()
	//指针示例
	demo.Point_demo()

}
