package main

import (
	"fmt"
	"sort"
)

func main() {

	var numbers = []int{2, 3, 1}
	var firstOddIndex = sort.Search(len(numbers), func(i int) bool { return numbers[i]%2 == 1 })
	fmt.Println(firstOddIndex)
}
