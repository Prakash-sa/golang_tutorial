package main

import (
	"fmt"
	"strconv"
)

func main() {

	i := 42
	fmt.Println("%v", i)

	var j string
	j = strconv.Itoa(i)
	fmt.Println("%v, %T\n", j, j)

}
