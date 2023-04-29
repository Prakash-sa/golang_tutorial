package main

import (
	"fmt"
	"time"
)

func fibonacci(x int) int {
	if x <= 1 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func memoize(fun func(int) int) func(int) int {
	var memo = map[int]int{}
	return func(x int) int {
		if r, exists := memo[x]; exists {
			return r
		}
		var r = fun(x)
		memo[x] = r
		return r
	}
}

func main() {
	var memFibonacci = memoize(fibonacci)
	for i := 1; i <= 2; i++ {
		start := time.Now()
		f37 := memFibonacci(37)
		delta := time.Now().Sub(start)
		nanosecond := delta.Nanoseconds()

		fmt.Println(fmt.Sprintf("%d: f37 is %d", i, f37))
		fmt.Println(fmt.Sprintf("%d: nanoseconds is %d", i, nanosecond))
	}

	start := time.Now()
	f38 := memFibonacci(38)
	delta := time.Now().Sub(start)
	nanosecond := delta.Nanoseconds()
	fmt.Println("f38 is", f38)
	fmt.Println("nanoseconds is", nanosecond)

}
