package main

import (
	"fmt"
	"time"
)

//写一个计时器程序，将 running() 函数并发执行，每隔5秒打印一次计数器，
//而 main 的 goroutine 则等待用户输入，两个行为可以同时进行.
func runnning() {
	var times int
	//构建一个无限循环
	for {
		times++
		fmt.Println("tick", times)
		//延时5秒
		time.Sleep(5 * time.Second)
	}

}

func main() {
	//并发执行程序
	go runnning()
	var input string
	fmt.Scanln(&input)
}
