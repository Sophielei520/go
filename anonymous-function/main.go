package main

import (
	"fmt"
	"time"
)

//在 main() 函数中创建一个匿名函数并为匿名函数启动 goroutine。
//匿名函数没有参数。代码将并行执行定时打印计数的效果。
func main() {
	go func() {
		var times int
		for {
			times++
			fmt.Println("tick", times)
			time.Sleep(time.Second)
		}
	}()

	var input string
	fmt.Scanln(&input)
	//所有 goroutine 在 main() 函数结束时会一同结束。

}
