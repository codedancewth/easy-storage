package main

import (
	"fmt"
)

func main() {
	// 新建一个redis的链接
	client := NewReClient()

	// 设置一个redis的set值
	ok, _ := client.Set("wth", "1", 30)
	fmt.Println(fmt.Sprintf("set value %+v", ok))

	value := ""
	// 获取一个redis的值
	ok2, _ := client.Get("wth", value)
	fmt.Println(fmt.Sprintf("get value %+v", ok2))
}
