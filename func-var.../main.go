package main

import (
	"fmt"
)

func main() {
	// 1.变参函数 1,2
	varChange(1, 2)
	// 2.变参函数 1,2,3
	varChange(1, 2, 3)
}

func varChange(nums ...int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)
}
