package logics

import (
	"errors"
)

func Sum(i1, i2 int) int {
	return i1 + i2
}

func Sub(i1, i2 int) int {
	return i1 - i2
}
func Add(i1, i2 int) int {
	return i1 * i2
}
func Div(i1, i2 int) (int, error) {
	if i2 == 0 {

		return i2, errors.New("Invalid syntex")
	}
	return i1 / i2, nil
}
