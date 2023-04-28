package main

import (
	"fmt"
)

func transform(arr []int, f func(int) int) {
	for i := range arr {
		arr[i] = f(arr[i])
	}
}

func main() {
	var numbers = []int{1, 2, 3, 4, 5}
	transform(numbers, func(x int) int {
		return x * 3
	})

	fmt.Println(numbers)
}
