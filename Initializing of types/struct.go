package main

import "fmt"

type size struct {
	width, height int
}

type point struct {
	top, left int
}

type rectangle struct {
	size  size
	point point
}

func main() {
	var rect = rectangle{
		size{10, 10},
		point{5, 5},
	}
	fmt.Println(rect)
}
