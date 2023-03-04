package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64 // satuan dalam liter
}

func (c *Car) EstimateDistance() float64 {
	fuelEfficiency := 1.5                                      // satuan liter/mill
	fuelInMill := c.FuelIn * 1000                              // konversi ke satuan milliliter
	distanceInKm := fuelInMill / (fuelEfficiency * 1000) * 100 // satuan jarak dalam kilometer
	return distanceInKm
}

func main() {
	myCar := Car{"Sedan", 20.0}
	estimatedDistance := myCar.EstimateDistance()
	fmt.Printf("Estimasi jarak yang bisa ditempuh oleh mobil %v dengan bahan bakar %v liter adalah %.2f km\n", myCar.Type, myCar.FuelIn, estimatedDistance)
}
