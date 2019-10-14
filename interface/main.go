package main

import (
	"fmt"
)

type geometry interface {
	area() float32
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func main() {
	r := rect{
		width:  2.2,
		height: 3.4,
	}
	fmt.Println(r.area())
}
