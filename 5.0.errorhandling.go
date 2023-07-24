package main

import (
	"errors"
	"fmt"
)

// by convention error value is the last returned value
func area(x, y int) (*int, error) {
	if x == 0 {
		return nil, errors.New("zero area")
	}
	if y == 0 {
		return nil, fmt.Errorf("zero area: [%v, %v]", x, y)
	}
	area := x * y
	return &area, nil
}

func main() {
	x := 2
	y := 0
	area, err := area(x, y)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("area(%v,%v) = %v;\n", x, y, area)
}
