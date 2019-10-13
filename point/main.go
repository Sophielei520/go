package main

import (
	"fmt"
)

func main() {
	i := 1
	fmt.Println("i = ", i)
	// 1.val
	zeroVal(i)
	fmt.Println("i = ", i)
	// 2.ptr
	zeroPtr(&i)
	fmt.Println("i = ", i)
}

func zeroVal(ival int) {
	ival = 0
}

func zeroPtr(iptr *int) {
	*iptr = 0
}
