package main

import (
	"fmt"
)

func main() {
	number := 50
	guess := 101

	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("Perfect!")
		}
	}

	var i interface{} = [2]int{}

	switch i.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float")
	case string:
		fmt.Println("string")
	case [2]int:
		fmt.Println("i is [2]int")
	default:
		fmt.Println("i is another number")
	}

}
