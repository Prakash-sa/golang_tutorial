package main

import "fmt"

type vehicle interface {
	test()
}

type car struct{}

func (c car) test() {
	fmt.Println("Test car")
}

type truck struct{}

func (t truck) test() {
	fmt.Println("Test truck")
}

func main() {

	var list = []vehicle{
		car{},
		truck{},
	}

	for _, curVehicle := range list {
		curVehicle.test()
	}

}
