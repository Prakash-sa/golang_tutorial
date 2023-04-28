package main

import "fmt"

func main() {
	var firstname = "1"
	var numberlist = ""

	switch firstname {
	case "1":
		numberlist = "1"
		fallthrough
	case "2":
		numberlist += "-2"
		fallthrough
	case "3":
		numberlist += "-3"
	}

	fmt.Println(numberlist)
}
