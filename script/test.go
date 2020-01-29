package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, world!")
	test()
	fmt.Println("测试")
	time.Sleep(5 * time.Second)
	fmt.Println("结束")
}

func test() {
	time.AfterFunc(3 * time.Second, func() {
		fmt.Println("after 3s")
	})
}