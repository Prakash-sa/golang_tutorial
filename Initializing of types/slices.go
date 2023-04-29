package main

import "fmt"

type employee struct {
	FirstName string
	LastName  string
}

func main() {
	var primeNumber = []int{2, 3, 5, 7, 11, 13, 17, 19}
	var numbers = [5]int{1, 2, 3, 4, 5}
	var gameList = []string{"soccer", "hockey", "basketball"}

	var employee = []employee{
		employee{"Anton", "Pavlov"},
		employee{"Elena", "Kirienko"},
	}

	fmt.Println(primeNumber)
	fmt.Println(numbers)
	fmt.Println(gameList)
	fmt.Println(employee)
}
