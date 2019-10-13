package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	// 1.排序前
	fmt.Println(strs)
	// 2.排序
	sort.Strings(strs)
	// 3.排序后
	fmt.Println(strs)
}
