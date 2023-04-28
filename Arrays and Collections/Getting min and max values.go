package main

import (
	"fmt"
	"math"
)

func min(arr []int) int {
	var result = math.MaxInt64
	for _, v := range arr {
		if v < result {
			result = v
		}
	}
	return result
}

func max(arr []int) int {
	var result = math.MinInt64
	for _, v := range arr {
		if result < v {
			result = v
		}
	}
	return result
}

func main() {
	var numbers = []int{2, 3, 5, 7, 11}

	var min = min(numbers)
	var max = max(numbers)

	fmt.Println("min is", min)
	fmt.Println("max is", max)

}
