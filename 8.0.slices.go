package main

import "fmt"

// slices are wrapper references on top of arrays

var newSlice = []int{4, 9, 2, 5}
var newSliceString = []string{"one", "two"}

var sliceViaMake = make([]int, 3, 10)

func sliceFromArray() {
	array := [5]int{1, 2, 5, 10, 3}
	slice := array[0:3]

	fmt.Print(slice) // [1 2 5]
}

func underlyingArrayIsChanged() {
	a := sliceViaMake[0:2]
	a[0] = 1

	fmt.Println(a)            // [1 0]
	fmt.Println(sliceViaMake) // [1 0 0]
}

func appendExample() {
	fmt.Println(len(newSlice)) // 4
	fmt.Println(cap(newSlice)) // 4

	appended := append(newSlice, 2) // append() returns new slice

	fmt.Println(appended)      // [4 9 2 5 2]
	fmt.Println(len(appended)) // 5
	fmt.Println(cap(appended)) // 8

	fmt.Println(newSlice)      // [4 9 2 5]
	fmt.Println(len(newSlice)) // 4
	fmt.Println(cap(newSlice)) // 4
}

func main() {
	sliceFromArray()
}
