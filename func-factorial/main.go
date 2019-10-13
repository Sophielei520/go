package main

import (
	"fmt"
)

func main() {
	fmt.Println(fact(5))
}

func fact(nums int) int {
	if nums == 0 {
		return 1
	}
	return nums * fact(nums-1)
}
