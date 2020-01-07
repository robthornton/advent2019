package fuel

import "math"

// ForMass returns the fuel for the given mass.
func ForMass(mass int64) int64 {
	return int64(math.Floor(float64(mass/3)) - 2)
}

// ForFuel returns the fuel required for fuel.
func ForFuel(mass int64) int64 {
	fuel := ForMass(mass)
	if fuel < 0 {
		return 0
	}
	return fuel + ForFuel(fuel)
}

// TotalForModule calculates the total fuel for a given module's mass.
func TotalForModule(mass int64) int64 {
	fuel := ForMass(mass)
	return fuel + ForFuel(fuel)
}
