package main

import (
	"fmt"
	"time"
)

func doSort(items []int, fst int, lst int) {
	if fst >= lst {
		return
	}

	i := fst
	j := lst
	x := items[(fst+lst)/2]

	for i < j {
		for items[i] < x {
			i++
		}

		for items[j] > x {
			j--
		}

		if i <= j {
			var tmp = items[i]
			items[i] = items[j]
			items[j] = tmp
			i++
			j--
		}
	}
	doSort(items, fst, j)
	doSort(items, i, lst)
}

func quickSort(arr []int) []int {
	length := len(arr)
	items := make([]int, length)
	copy(items, arr)
	doSort(items, 0, length-1)
	return items
}

func main() {
	items := []int{4, 1, 5, 3, 2}

	sortItems := quickSort(items)

	//simplified speed test
	items = make([]int, 200)
	for i := 0; i < len(items); i++ {
		items[i] = i
	}
	var tmp = items[5]
	items[5] = items[6]
	items[6] = tmp
	count := 10000
	start := time.Now()
	for i := 0; i < count; i++ {
		quickSort(items)
	}

	delta := time.Now().Sub(start)
	nanoseconds := delta.Nanoseconds()

	fmt.Println(sortItems)
	fmt.Println(nanoseconds)

}
