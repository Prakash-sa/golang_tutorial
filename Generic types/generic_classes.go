package main

import "fmt"

type number interface {
	int8 | float32
}

type size[T number] struct {
	width, height T
}

func (s size[T]) AsText() string {
	return fmt.Sprintf("[%v; %v]", s.width, s.height)
}

func main() {

	var sizeInt = size[int8]{5, 8}
	var textInt = sizeInt.AsText()

	var sizeFloat = size[float32]{3.7, 1.58}

	var textFloat = sizeFloat.AsText()

	fmt.Println("textInt is", textInt)
	fmt.Println("textFloat is", textFloat)
}
