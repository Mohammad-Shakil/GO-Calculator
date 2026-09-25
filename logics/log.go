package logics

import (
	"errors"
)

func Sum(i1, i2 float64) float64 {
	return i1 + i2
}

func Sub(i1, i2 float64) float64 {
	return i1 - i2
}
func Mul(i1, i2 float64) float64 {
	return i1 * i2
}
func Div(i1, i2 float64) (float64, error) {
	if i2 == 0 {

		return i2, errors.New("\nInvalid operation")
	}
	return i1 / i2, nil
}
