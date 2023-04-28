package main

import "fmt"

func main() {
	var numbers1 = []int{1, 2, 3, 4, 5}
	var numbers2 = make([]int, len(numbers1))
	copy(numbers2, numbers1)

	fmt.Println(numbers2)
}
