package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("square root of negative number")
	}
	return math.Sqrt(x), nil
}

func sqrtSum(x, y float64) (float64, error) {
	sqrtX, err := sqrt(x)
	if err != nil {
		return sqrtX, errors.New("Sqrt(x) error: " + err.Error())
	}

	sqrtY, err := sqrt(y)
	if err != nil {
		return sqrtY, errors.New("Sqrt(y) error: " + err.Error())
	}
	return sqrtX + sqrtY, nil
}

func main() {
	sqrtSum, err := sqrtSum(16.0, -25.0)

	fmt.Println(sqrtSum)
	fmt.Println(err)
}
