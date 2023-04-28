package main

import "fmt"

type simpleException struct {
	errText string
}

func (e simpleException) Error() string {
	return e.errText
}

func throwSimpleException() error {
	return simpleException{"simple Exception"}
}

func main() {
	err := throwSimpleException()

	if err != nil {
		fmt.Println(err)
	}
}
