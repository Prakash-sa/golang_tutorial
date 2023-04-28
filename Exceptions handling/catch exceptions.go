package main

import (
	"errors"
	"fmt"
)

func throwWhenNullOrEmpty(arr []int) error {
	if arr == nil {
		return errors.New("Array is nil")
	}
	if len(arr) == 0 {
		return errors.New("Array is empty")
	}
	return nil
}

//--All exception--

// func main() {
// 	err := throwWhenNullOrEmpty([]int{})

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

//--specific--

func main() {
	err := throwWhenNullOrEmpty([]int{})

	if err.Error() == "Array is nil" {
		fmt.Println("Error: array is specified")
	} else if err.Error() == "Array is empty" {
		fmt.Println("Error: array is empty")
	}
}
