package main

import "fmt"

func merge(ms ...map[int]string) map[int]string {
	res := map[int]string{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

func main() {
	var d1 = map[int]string{1: "one"}
	var d2 = map[int]string{2: "two"}
	var d3 = map[int]string{3: "three"}

	var dAll = merge(d1, d2, d3)

	fmt.Println(dAll)
}
