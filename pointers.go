package main

import "fmt"

func declare() {
	var ptr *string
	greeting := "Hello"
	ptr = &greeting

	fmt.Println("Address:", ptr)
	fmt.Println("Value:", *ptr)
}

func declare1() {
	x := 3
	y := &x

	fmt.Println(y)
	fmt.Println(*y)
}

func negate(x *int) {
	neg := -*x
	*x = neg
}
