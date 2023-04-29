package main

import "fmt"

type Strategy interface {
	DoOperation(a int, b int) int
}

type AddStrategy struct{}

func (s AddStrategy) DoOperation(a int, b int) int {
	return a + b
}

type SubstractStrategy struct{}

func (s SubstractStrategy) DoOperation(a int, b int) int {
	return a - b
}

type Calc struct {
	_strategy Strategy
}

// execute the current state
func (c Calc) Execute(a int, b int) int {
	if c._strategy == nil {
		return 0
	}
	return c._strategy.DoOperation(a, b)
}

// changes the current state
func (c *Calc) SetStrategy(strategy Strategy) {
	c._strategy = strategy
}

func main() {

	calc := Calc{}
	result1 := calc.Execute(5, 3)

	calc.SetStrategy(AddStrategy{})
	result2 := calc.Execute(5, 3)

	calc.SetStrategy(SubstractStrategy{})
	result3 := calc.Execute(5, 3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

}
