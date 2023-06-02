package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func subAndSwap(x int, y int) (int, int, int) {
	sub := func(a, b int) int {
		return a - b
	}
	return sub(x, y), y, x
}

func rectangle(x, y int) (area int, circumf int) {
	area = x * y
	circumf = 2 * (x + y)
	return
}

func ignoreReturnValue() {
	area, _ := rectangle(1, 2)
	fmt.Println(area)
}
