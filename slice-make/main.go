package main

import (
	"fmt"
)

func main() {
	// 1.创建空slice
	slice1 := make([]string, 3)
	fmt.Println(slice1)
	// 2.添加元素到slice
	slice1[0] = "a"
	slice1[1] = "b"
	slice1[2] = "c"
	fmt.Println(slice1)
}
