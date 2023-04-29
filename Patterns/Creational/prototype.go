package main

import "fmt"

type Shape interface {
	Clone() Shape
}

type Square struct {
	LineCount int
}

func (s Square) Clone() Shape {
	return Square{s.LineCount}
}

//Client

type ShapeMaker struct {
	Shape Shape
}

func (sm ShapeMaker) MakeShape() Shape {
	return sm.Shape.Clone()
}

func main() {
	square := Square{4}
	maker := ShapeMaker{square}

	square1 := maker.MakeShape()
	square2 := maker.MakeShape()

	fmt.Println(square1)
	fmt.Println(square2)
}
