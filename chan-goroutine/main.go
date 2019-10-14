package main

import "fmt"

// 1.将1-10放入通道
func counts(natural chan int) {
	count := 10
	for index := 0; index < count; index++ {
		natural <- index
	}
	close(natural)
}

// 2.将通道中的数据传入平方数通道
func squarer(squares, natural chan int) {
	for v := range natural {
		squares <- v * v
	}
	close(squares)
}

// 3.打印输出
func printer(squares chan int) {
	for v := range squares {
		fmt.Println(v)
	}
}

func main() {
	natural := make(chan int)
	square := make(chan int)

	go counts(natural)
	go squarer(square, natural)
	printer(square)
}
