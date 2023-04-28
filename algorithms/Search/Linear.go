package main

import (
	"fmt"
	"time"
)

func linearSearch(arr []int, x int) int {
	i := 0
	count := len(arr)
	for i < count {
		if arr[i] == x {
			return i
		}
		i++
	}
	return -1
}

func main() {
	items := []int{2, 3, 5, 7, 11, 13, 17}
	fmt.Println(linearSearch(items, 1))
	fmt.Println(linearSearch(items, 3))

	// simplified speed test

	items = make([]int, 10000000)
	for i := 0; i < len(items); i++ {
		items[i] = i
	}
	count := 100
	start := time.Now()
	for i := 0; i < count; i++ {
		linearSearch(items, 777777)
	}

	delta := time.Now().Sub(start)
	nanoseconds := delta.Nanoseconds()

	fmt.Println(nanoseconds)
}
