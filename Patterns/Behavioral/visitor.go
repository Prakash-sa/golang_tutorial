package main

import (
	"fmt"
	"strconv"
)

type Element interface {
	Accept(v CarVisitor)
}

type Engine struct{}

func (e Engine) Accept(v CarVisitor) {
	v.visitEngine(e)
}

type Wheel struct {
	Number int
}

func (w Wheel) Accept(v CarVisitor) {
	v.visitWheel(w)
}

type Car struct {
	_items []Element
}

func (c Car) Accept(v CarVisitor) {
	for _, e := range c._items {
		e.Accept(v)
	}
	v.visitCar(c)
}

type CarVisitor interface {
	visitEngine(engine Engine)
	visitWheel(wheel Wheel)
	visitCar(car Car)
}

type TestCarVisitor struct{}

func (v TestCarVisitor) visitEngine(engine Engine) {
	fmt.Println("Test Engine")
}

func (v TestCarVisitor) visitWheel(wheel Wheel) {
	fmt.Println("test wheel #" + strconv.Itoa(wheel.Number))
}

func (v TestCarVisitor) visitCar(car Car) {
	fmt.Println("test car")
}

type RepairCarVisitor struct{}

func (v RepairCarVisitor) visitEngine(engine Engine) {
	fmt.Println("repair engine")
}

func (v RepairCarVisitor) visitWheel(wheel Wheel) {
	fmt.Println("repair wheel #" + strconv.Itoa(wheel.Number))
}

func (v RepairCarVisitor) visitCar(car Car) {
	fmt.Println("repair car")
}

func main() {
	car := Car{[]Element{
		Engine{},
		Wheel{1},
		Wheel{2},
		Wheel{3},
		Wheel{4},
	}}
	v1 := TestCarVisitor{}
	v2 := RepairCarVisitor{}
	car.Accept(v1)
	car.Accept(v2)
}
