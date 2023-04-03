package main

import (
	"fmt"
)

func main() {
	// var students [3]string
	// fmt.Printf("%v\n", students)
	// students[0] = "Prakash"
	// fmt.Printf("%v", students)

	//slice

	// a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	// b := a[3:]
	// fmt.Println(a)
	// fmt.Println(b)

	// a:=make([]int,3,100)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n",len(a))

	a := []int{}
	a = append(a, 1)
	fmt.Println(a)
}
