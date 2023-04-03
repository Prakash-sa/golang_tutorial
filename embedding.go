package main

import (
	"fmt"
	"reflect"
)

type Animal1 struct {
	Name   string `required max:100`
	Origin string
}
type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false

	a := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(b)

	fmt.Println(a)

	t := reflect.TypeOf(Animal1{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
