package main

import "fmt"

func main() {
	var primeNumbers = []int{2, 5, 7}
	primeNumbers = append(primeNumbers, 11)
	// {2,5,7,11}
	fmt.Println(primeNumbers)

	primeNumbers = append(primeNumbers[0:1], append([]int{3},
		primeNumbers[1:]...)...)
	// {2,3,5,7,11}
	fmt.Println(primeNumbers)

	primeNumbers = append(primeNumbers[0:2], primeNumbers[3:]...)
	// {2,3,7,11}
	fmt.Println(primeNumbers)

	primeNumbers = append(primeNumbers, 13, 17)
	// {2,3,7,11,13,17}
	fmt.Println(primeNumbers)

	primeNumbers = append([]int{2, 3, 5}, primeNumbers[2:]...)
	// {2,3,5,7,11,13,17}
	fmt.Println(primeNumbers)

	primeNumbers = primeNumbers[:0]
	// {}
	fmt.Println(primeNumbers)
}
