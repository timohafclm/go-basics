package main

import "fmt"

// var declare examples

var greeting string           // ""
var defaultValueInteger int   // 0
var defaultValueFloat float64 // 0.0

var hello = "Hello"

var ExportedName string

func main() {
	// only function scope
	shortDeclaration := "sf"

	fmt.Println(shortDeclaration)
}
