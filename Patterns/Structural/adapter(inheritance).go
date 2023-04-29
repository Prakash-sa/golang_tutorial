package main

import (
	"fmt"
	"strings"
)

type StringList struct {
	rows []string
}

func (sl StringList) getString() string {
	return strings.Join(sl.rows, "\n")
}

func (sl *StringList) add(value string) {
	sl.rows = append(sl.rows, value)
}

type TextAdapter struct {
	StringList
}

func (ta TextAdapter) getText() string {
	return ta.getString()
}

func getTextAdapter() TextAdapter {
	adapter := TextAdapter{}
	adapter.add("Line 1")
	adapter.add("line 2")
	return adapter
}

func main() {
	adapter := getTextAdapter()
	text := adapter.getText()

	fmt.Println(text)
}
