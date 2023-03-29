package main

import "fmt"

/**
NOTE: interfaces in go work in same way as any other language.
Mainly used to have api contracts just like other langs.

NOTE: In go the implementation is a bit different.
You define interface, and a struct.
then struct will implement the interface, if the you have defined the method
with the struct.
*/

type Bike interface {
	GetMileage() string
	GetSpecialFeatures() string
}

type BikeAttributes struct {
	Cost     int
	Name     string
	TopSpeed int
	Mileage  int
	Features string
}

func (g BikeAttributes) GetMileage() string {
	return fmt.Sprintf("%s, mileage is %d", g.Name, g.Mileage)
}

func (b BikeAttributes) GetSpecialFeatures() string {
	return fmt.Sprintf("%s, special features are: %s", b.Name, b.Features)
}

func main() {
	var gixerSF Bike
	gixerSF = BikeAttributes{
		Cost:     170000,
		Name:     "Tuktuk - Gixer",
		TopSpeed: 125,
		Mileage:  35,
		Features: "Sexy on road, smooth engine.",
	}

	fmt.Println(gixerSF.GetSpecialFeatures())

	var continentalGT Bike
	continentalGT = BikeAttributes{
		Cost:     400000,
		Name:     "Continental GT",
		TopSpeed: 140,
		Mileage:  30,
		Features: "0-100 in 5 seconda, goes like a beast",
	}

	fmt.Println(continentalGT.GetSpecialFeatures())

}
