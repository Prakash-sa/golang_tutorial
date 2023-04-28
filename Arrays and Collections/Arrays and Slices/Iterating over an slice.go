package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numbers = []int{2, 3, 5, 7, 11, 13, 17}
	var str = ""
	for _, number := range numbers {
		if len(str) > 0 {
			str += "; "
		}
		str += strconv.Itoa(number)
	}
	fmt.Println(str)
}
