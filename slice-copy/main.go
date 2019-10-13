package main

import (
	"fmt"
)

func main() {
	// 1.创建空slice
	slice1 := make([]string, 3)
	// 2.添加元素到slice
	slice1[0] = "a"
	slice1[1] = "b"
	slice1[2] = "c"
	// 3.copy
	slice2 := make([]string, len(slice1))
	copy(slice2, slice1)
	fmt.Println(slice2)
}
