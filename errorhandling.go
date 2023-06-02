package main

import "fmt"

func area(x, y int) (*int, error) {
	if x == 0 || y == 0 {
		return nil, fmt.Errorf("zero area: [%v, %v]", x, y)
	}
	area := x * y
	return &area, nil
}
