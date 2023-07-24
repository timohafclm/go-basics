package main

import "fmt"

func main() {
	swap, i, i2 := subAndSwap(1, 2)

	fmt.Println(swap, i, i2)
	deferExample()
}

// func add(x int) {} - Go doesn't support method overloading

func add(x, y int) int {
	return x + y
}

func subAndSwap(x int, y int) (int, int, int) {
	sub := func(a, b int) int {
		return a - b
	}
	return sub(x, y), y, x
}

// named return values are only for short functions because they harm readability
func rectangle(x, y int) (area int, circumf int) {
	area = x * y
	circumf = 2 * (x + y)
	return
}

func ignoreReturnValue() {
	area, _ := rectangle(1, 2)
	fmt.Println(area)
}

// defer functions are used for cleanup operations
func deferExample() {
	msg := "Hello"
	defer printMsg("Defer-0", msg)
	defer printMsg("Defer-1", msg)

	msg = "Bye"
	fmt.Println("deferExample has exited")
	// output:
	// deferExample has exited
	// [Defer-1]: Hello
	// [Defer-0]: Hello (it was evaluated immediately)
}

func printMsg(id string, s string) {
	fmt.Printf("[%v]: %v\n", id, s)
}
