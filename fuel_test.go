package main

import "testing"

func TestFuelForMass(t *testing.T) {
	fuelValues := []struct {
		mass     int64
		expected int64
	}{
		{mass: 12, expected: 2},
		{mass: 14, expected: 2},
		{mass: 1969, expected: 654},
		{mass: 100756, expected: 33583},
	}

	for _, fv := range fuelValues {
		got := fuelForMass(fv.mass)

		if got != fv.expected {
			t.Logf("expected %d, got %d", fv.expected, got)
			t.Fail()
		}
	}
}

func TestFuelForFuel(t *testing.T) {
	fuelValues := []struct {
		fuel     int64
		expected int64
	}{
		{fuel: -1, expected: 0},
		{fuel: 2, expected: 0},
		{fuel: 654, expected: 216 + 70 + 21 + 5},
		{fuel: 33583, expected: 11192 + 3728 + 1240 + 411 + 135 + 43 + 12 + 2},
	}

	for _, fv := range fuelValues {
		got := fuelForFuel(fv.fuel)

		if got != fv.expected {
			t.Logf("expected %d, got %d", fv.expected, got)
			t.Fail()
		}
	}
}

func TestTotalFuelForModule(t *testing.T) {
	fuelValues := []struct {
		mass     int64
		expected int64
	}{
		{mass: 12, expected: 2},
		{mass: 1969, expected: 966},
		{mass: 100756, expected: 50346},
	}

	for _, fv := range fuelValues {
		got := totalFuelForModule(fv.mass)

		if got != fv.expected {
			t.Logf("expected %d, got %d", fv.expected, got)
			t.Fail()
		}
	}
}
