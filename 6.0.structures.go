package main

import "fmt"

var (
	beachVacationThreshold float64 = 22
	skiVacationThreshold   float64 = -2
)

type CityTemp interface {
	Name() string
	TempC() float64
	TempF() float64
	SetTempC(t float64)

	BeachVacationReady() bool
	SkiVacationReady() bool
}

type city struct {
	name        string
	tempC       float64
	hasBeach    bool
	hasMountain bool
}

// func (c city) name {} - field and method names can't clash

func NewCityTemp(name string, tempC float64) CityTemp {
	return &city{name: name, tempC: tempC}
}

func NewCityTempVerbose(name string, tempC float64, hasBeach bool, hasMountain bool) CityTemp {
	return &city{
		name:        name,
		tempC:       tempC,
		hasBeach:    hasBeach,
		hasMountain: hasMountain}
}

func newCity(name string, tempC float64) *city {
	return &city{name: name, tempC: tempC}
}

func (c city) Name() string { // methods are special functions with special receiver arguments (c city)
	return c.name
}

func (c city) TempC() float64 {
	return c.tempC
}

func (c city) TempF() float64 {
	return (c.tempC * 9 / 5) + 32
}

func TempFEquivalent(c city) float64 { // it's equivalent function of the TempF method, just passing by city instance
	return (c.tempC * 9 / 5) + 32
}

func (c *city) SetTempC(temp float64) { // use pointers to persist changes outside of function scope
	c.tempC = temp
}

func (c city) BeachVacationReady() bool {
	return c.tempC > beachVacationThreshold && c.hasBeach
}

func (c city) SkiVacationReady() bool {
	return c.tempC <= skiVacationThreshold && c.hasMountain
}

func (c city) nonInterfaceMethodSetTempC(temp float64) {
	c.tempC = temp
}

func passInterface(c CityTemp, tempC float64) {
	c.SetTempC(tempC)
}

func main() {
	var c = city{name: "Limassol", tempC: 35}
	c.tempC = 20
	fmt.Println(c) // {Limassol 20 false false}, 20 is only in main() scope

	var c1 = newCity("Minsk", 3)
	c1.tempC = 10
	c1.nonInterfaceMethodSetTempC(15)
	fmt.Println(c1) // {Minsk 10 false false}

	var c2 = NewCityTemp("Damask", 30)
	c2.SetTempC(31)
	fmt.Println(c2) // &{Damask 31 false false}

	var c3 = NewCityTemp("LA", 20)
	passInterface(c3, 30)
	fmt.Println(c3) // &{LA 30 false false}, interfaces are built-in pointer types
}
