package main

import "fmt"

func main() {
	var n1 = []int{}
	var n2 = make([]int, 0)

	var n3 = []int{1, 2, 3}
	var s1 = []string{"1", "2", "3"}
	a1 := [5]float64{2, 3, 5, 7, 11}

	var n4 = [][]int{{1, 2}, {3, 4}}
	n4[1][1] = 7

	var a2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	a2[1][1] = 7

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(s1)
	fmt.Println(a1)
	fmt.Println(n4)
	fmt.Println(a2)
}
