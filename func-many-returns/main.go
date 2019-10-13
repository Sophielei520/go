package main

import (
	"fmt"
)

func main() {
	a, b := manyReturns()
	fmt.Println(a)
	fmt.Println(b)
}

func manyReturns() (int, int) {
	return 3, 5
}
