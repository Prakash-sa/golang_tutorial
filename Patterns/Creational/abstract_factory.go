package main

import "fmt"

type ProductA interface {
	TestA()
}

type ProductB interface {
	TestB()
}

type Factory interface {
	CreateA() ProductA
	CreateB() ProductB
}

type ProductA1 struct{}

func (p ProductA1) TestA() {
	fmt.Println("Test A1")
}

type ProductB1 struct{}

func (p ProductB1) TestB() {
	fmt.Println("Test B1")
}

type ProductA2 struct{}

func (p ProductA2) TestA() {
	fmt.Println("Test A2")
}

type ProductB2 struct{}

func (p ProductB2) TestB() {
	fmt.Println("Test b2")
}

type Factory1 struct{}

func (f Factory1) CreateA() ProductA {
	return ProductA1{}
}

func (f Factory1) CreateB() ProductB {
	return ProductB1{}
}

type Factory2 struct{}

func (f Factory2) CreateA() ProductA {
	return ProductA2{}
}

func (f Factory2) CreateB() ProductB {
	return ProductB2{}
}

//client code:

func TestFactory(factory Factory) {
	productA := factory.CreateA()
	ProductB := factory.CreateB()

	productA.TestA()
	ProductB.TestB()
}

func main() {
	TestFactory(Factory1{})
	TestFactory(Factory2{})
}
