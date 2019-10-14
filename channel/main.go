package main

import (
	"fmt"
)

func main() {
	//channel的正常使用
	fmt.Println("Main Start")
	ch1 := make(chan int)
	go func() {
		ch1 <- 10
		fmt.Println("Write data into channel")
	}()

	i := <-ch1
	fmt.Print("Read data from channel", i)
	fmt.Println("Main exit")
}
