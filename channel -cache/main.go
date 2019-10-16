package main

import (
	"fmt"
	"time"
)

func main() {
	//有缓存的channel不阻塞的例子
	ch1 := make(chan int, 10)
	//开启一个gorutine向channel里写数据。
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
			fmt.Println("Write data into channel", i)
		}
	}()
	//主gorutine从channel里读取数据
	for i := 0; i < 5; i++ {
		temp := <-ch1
		fmt.Println("Read data from channel", temp)
	}
	time.Sleep(2 * time.Second)
}
