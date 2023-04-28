package main

import "fmt"

func main() {
	var numbers = []int{2, 3, 5, 7, 11}

	var first2 = numbers[:2]

	var last3 = numbers[2:]

	fmt.Println("first 2", first2)
	fmt.Println("last 3", last3)
}
