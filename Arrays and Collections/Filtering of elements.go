package main

import "fmt"

func filter(arr []int, f func(int) bool) []int {
	result := make([]int, 0)

	for _, v := range arr {
		if f(v) {
			result = append(result, v)
		}
	}
	return result

}

func main() {
	var numbers = []int{1, 2, 3, 4, 5}
	var oddItems = filter(numbers, func(v int) bool {
		return v%2 == 1
	})
	fmt.Println(oddItems)
}
