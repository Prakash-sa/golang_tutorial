package main

import (
	"fmt"
)

func fibbonacci(x int) int {
	if x <= 1 {
		return x
	}
	return fibbonacci(x-1) + fibbonacci(x-2)
}

func main() {
	var f10 = fibbonacci(10)

	fmt.Println(f10)
}
