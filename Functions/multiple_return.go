package main

import "fmt"

func getFirstLast(ar []int) (int, int) {
	var first = -1
	var last = -1

	if len(ar) > 0 {
		first = ar[0]
		last = ar[len(ar)-1]
	}
	return first, last
}

func main() {
	var arr = []int{2, 3, 5}
	first, last := getFirstLast(arr)
	fmt.Println(first)
	fmt.Println(last)
}
