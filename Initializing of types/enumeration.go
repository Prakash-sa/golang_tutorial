package main

import "fmt"

const (
	Platinum = iota
	Gold
	Silver
)

const (
	Mercury = iota + 1
	Venus
	Earth
)

func main() {
	var gold = Gold
	var venus = Venus
	fmt.Println(gold)
	fmt.Println(venus)
}
