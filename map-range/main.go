package main

import (
	"fmt"
)

func main() {
	map1 := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range map1 {
		fmt.Printf("%s->%s\n", k, v)
	}
}
