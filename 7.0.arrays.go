package main

import "fmt"

var emptyArray = [2]int{}

var threeLengthArray = [3]int{1, 2, 5}

var compilerInferLength = [...]string{"one", "two"}

func addCity(city string, cities [2]string) {
	cities[1] = city
}

func addCityPointer(city string, cities *[2]string) {
	cities[1] = city
}

func main() {
	cities := [2]string{"London"}
	copyCities := cities
	copyPointers := cities

	cities[1] = "New York"
	addCity("Miami", copyCities)
	addCityPointer("San Francisco", &copyPointers)

	fmt.Println("cities: ", cities)
	fmt.Println("copy: ", copyCities)
	fmt.Println("copyPointers: ", copyPointers)
	fmt.Println("emptyArray: ", emptyArray)
	fmt.Println("emptyArray length: ", len(emptyArray))
}
