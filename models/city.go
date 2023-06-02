package models

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

func NewCity(name string, tempC float64) CityTemp {
	return &city{name: name, tempC: tempC}
}

func NewCityVerbose(name string, tempC float64, hasBeach bool, hasMountain bool) CityTemp {
	return &city{
		name:        name,
		tempC:       tempC,
		hasBeach:    hasBeach,
		hasMountain: hasMountain}
}

func (c city) Name() string {
	return c.name
}

func (c city) TempC() float64 {
	return c.tempC
}

func (c city) TempF() float64 {
	return (c.tempC * 9 / 5) + 32
}

func (c *city) SetTempC(temp float64) {
	c.tempC = temp
}

func (c city) BeachVacationReady() bool {
	return c.tempC > beachVacationThreshold && c.hasBeach
}

func (c city) SkiVacationReady() bool {
	return c.tempC <= skiVacationThreshold && c.hasMountain
}

func setWhenUsePointer() {
	c := NewCity("London", 15.5)
	c.SetTempC(45)
	fmt.Println(c)
}
