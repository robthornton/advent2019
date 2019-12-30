package main

import "math"

func fuelForMass(mass int64) int64 {
	return int64(math.Floor(float64(mass/3)) - 2)
}

func fuelForFuel(mass int64) int64 {
	fuel := fuelForMass(mass)
	if fuel < 0 {
		return 0
	}
	return fuel + fuelForFuel(fuel)
}

func totalFuelForModule(mass int64) int64 {
	fuel := fuelForMass(mass)
	return fuel + fuelForFuel(fuel)
}
