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
	// 3.append添加元素
	slice1 = append(slice1, "d")
	slice1 = append(slice1, "e", "f")
	fmt.Println(slice1)
}
