package main

import "fmt"

type queue []int

func (q queue) Push(v int) queue {
	return append(q, v)
}

func (q queue) Shift() (queue, int) {
	return q[1:], q[0]
}

func main() {
	intQueue := make(queue, 0)
	intQueue = intQueue.Push(1)
	intQueue = intQueue.Push(3)
	intQueue = intQueue.Push(5)

	intQueue, first := intQueue.Shift()
	//first is 1

	intQueue, second := intQueue.Shift()
	//second is 3

	intQueue, third := intQueue.Shift()
	//third is 5

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)

}
