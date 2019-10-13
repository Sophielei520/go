package main

import (
	"fmt"
)

type Student struct {
	name  string
	age   int
	score float64
}

func main() {
	s := Student{
		name:  "Richard",
		age:   5,
		score: 98.5,
	}
	fmt.Println(s)
}
