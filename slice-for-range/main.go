package main

import (
	"fmt"
)

func main() {
	slice1 := []string{"a", "b", "c"}
	for _, names := range slice1 {
		fmt.Println(names)
	}
}
