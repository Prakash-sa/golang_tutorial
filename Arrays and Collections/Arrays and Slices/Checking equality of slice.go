package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {

	var n1 = []byte{1, 2, 3}
	var n2 = []byte{1, 2, 3}
	var n3 = []byte{3, 2, 1}

	var equal1 = bytes.Compare(n1, n2) == 0

	var equal2 = reflect.DeepEqual(n1, n2)

	var equal3 = bytes.Compare(n1, n3)

	fmt.Println(equal1)
	fmt.Println(equal2)
	fmt.Println(equal3)
}
