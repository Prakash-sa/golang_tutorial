package main

import (
	"fmt"
	"sort"
)

func main() {
	var numbers = []int{11, 2, 5, 7, 3}
	sort.Ints(numbers)
	//{2,3,5,7,11}
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	//{11,7,5,3,2}

	fmt.Println(numbers)

}
