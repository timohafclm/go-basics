package main

import "fmt"

var map1 = map[string]int{
	"ten": 10,
}

func initMap() {
	m := make(map[string]int)
	m["one"] = 1
	fmt.Println(m)
	delete(m, "one")
}

func missingKey() {
	value, ok := map1["two"]
	fmt.Println("is it exist:", ok)
	fmt.Println("value:", value)
}

func main() {
	missingKey()
}
