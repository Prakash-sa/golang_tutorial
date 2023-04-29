package main

import "fmt"

type named interface {
	getName() string
}

// sity is
type Sity struct {
	Name string
}

func (s Sity) getName() string {
	return s.Name
}

// star is
type Star struct {
	Name string
}

func (s Star) getName() string {
	return s.Name
}

func main() {
	var rows = []named{
		Sity{"Rome"},
		Star{"Sirius"},
	}

	list := ""

	for _, row := range rows {
		if len(list) > 0 {
			list += ", "
		}
		list += row.getName()
	}
	fmt.Println(list)
}
