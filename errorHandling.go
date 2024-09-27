package main

import (
	"errors"
)

func errorHandling(a, b float64) (result float64, err error) {
	if b == 0 {
		err = errors.New("division by zero is not possible")
		return
	}
	result = a / b
	return // Automatically returns 'result' and 'err' (which is nil here)
}
