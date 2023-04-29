package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	s2 := "Hello World!"

	for _, v := range s2 {
		fmt.Printf(string(v))
	}
}
