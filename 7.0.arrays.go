package main

import "fmt"

var emptyArray = [2]int{} // [0 0] - zero values were initialized

var threeLengthArray = [3]int{1, 2, 5}

var compilerInferLength = [...]string{"one", "two"}

func arrayLength() int {
	return len(threeLengthArray)
}

func addCity(city string, cities [2]string) {
	cities[1] = city
}

func addCityPointer(city string, cities *[2]string) {
	cities[1] = city
}

// arrays are copied when we assign them to variables or passed to functions
func arraysAreValueTypes() {
	cities := [2]string{"London"}
	copyCities := cities
	copyPointers := cities

	cities[1] = "New York"
	addCity("Miami", copyCities)
	addCityPointer("Paris", &copyPointers)

	fmt.Println("cities: ", cities)             // [London, New York]
	fmt.Println("copy: ", copyCities)           // [London]
	fmt.Println("copyPointers: ", copyPointers) // [London, Paris]
}

func main() {
	arraysAreValueTypes()
}
