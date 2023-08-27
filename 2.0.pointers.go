package main

import "fmt"

// * dereferencing operator
// & address operator

var ptr *string

func dereference() {
	greeting := "Hello"
	ptr = &greeting

	fmt.Println("Address:", ptr)
	fmt.Println("Value:", *ptr)
}

func declare1() {
	x := 3
	y := &x

	fmt.Println(y)  // 0xc00001a0f0
	fmt.Println(*y) // 3
}

func negateSimple(x int) {
	neg := -x
	x = neg
}

func negatePointer(x *int) {
	neg := -*x
	*x = neg
}

// variables are passed by value except pointers
func negateExample() {
	x := 3
	fmt.Println(x) // 3
	negatePointer(&x)
	fmt.Println(x) // -3

	y := 4
	fmt.Println(y) // 4
	negateSimple(y)
	fmt.Println(y) // 4
}
