package main

import (
	"fmt"
	"time"
)

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
