package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{2, 1, 3}
	// 1.排序前
	fmt.Println(ints)
	// 2.排序
	sort.Ints(ints)
	// 3.排序后
	fmt.Println(ints)
}
