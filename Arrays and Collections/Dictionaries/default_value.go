package main

import (
	"fmt"
	"strconv"
)

func main() {
	var dic1 = map[int]string{1: "one"}
	var strValue = dic1[2]
	// ""

	var dic2 = map[string]int{"one": 1}

	var intValue = dic2["two"]
	//0

	fmt.Println("\"" + strValue + "\"")
	fmt.Println("\"" + strconv.Itoa(intValue) + "\"")
}
