package main

import "fmt"

func printParity(x int) {
	if x%2 == 0 {
		fmt.Printf("%v is even\n", x)
	} else {
		fmt.Printf("%v is odd\n", x)
	}
}

func printParity1(x int) {
	if x%2 == 0 {
		fmt.Printf("%v is even\n", x)
		return
	}
	fmt.Printf("%v is odd\n", x)
}

func printParity2(x int) {
	if r := x % 2; r == 0 {
		fmt.Printf("%v is even\n", x)
		return
	}
	// r is out of scope here
	fmt.Printf("%v is odd\n", x)
}
