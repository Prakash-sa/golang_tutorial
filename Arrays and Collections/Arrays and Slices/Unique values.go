package main

import "fmt"

func unique(input []int) []int {
	result := make([]int, 0, len(input))

	values := make(map[int]bool)

	for _, val := range input {
		if _, ok := values[val]; !ok {
			values[val] = true
			result = append(result, val)
		}
	}
	return result
}

func main() {
	var numbers = []int{1, 3, 2, 1, 3}
	var unique = unique(numbers)

	fmt.Println(unique)

}
