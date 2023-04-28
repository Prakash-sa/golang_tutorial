package main

import (
	"fmt"
	"strconv"
	"time"
)

func add(a, b int) int {
	time.Sleep(3000)
	return a + b
}

func startGoroutine() {
	result := add(3, 5)
	fmt.Println("result: " + strconv.Itoa(result))
}

func main() {
	go startGoroutine()
	fmt.Println("main thread")
	var input string
	fmt.Scanln(&input)
}
