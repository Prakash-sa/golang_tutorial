package main

import "fmt"

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func main() {
	intStack := make(stack, 0)
	intStack = intStack.Push(1)
	intStack = intStack.Push(5)
	intStack = intStack.Push(3)

	intStack, first := intStack.Pop()
	fmt.Println(first)
	//3

	intStack, second := intStack.Pop()
	fmt.Println(second)
	//5

	intStack, third := intStack.Pop()
	fmt.Println(third)
	//1
}
