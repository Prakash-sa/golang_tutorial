package main

import "fmt"

func indexOf(arr []int, v int) int {
	for i := range arr {
		if arr[i] == v {
			return i
		}
	}
	return -1
}

func main() {
	var numbers = []int{2, 3, 5, 7, 11, 13}

	var contain5 = indexOf(numbers, 5) != -1

	var index5 = indexOf(numbers, 5)
	var contain10 = indexOf(numbers, 10) != -1

	fmt.Println(contain5)
	fmt.Println(index5)
	fmt.Println(contain10)
}
