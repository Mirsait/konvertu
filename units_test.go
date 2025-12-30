package main

import (
	"math"
	"testing"
)

func almostEqual4(a, b float64) bool {
	const epsilon = 1e-4
	return math.Abs(a-b) < epsilon
}

func almostEqual3(a, b float64) bool {
	const epsilon = 1e-3
	return math.Abs(a-b) < epsilon
}

func TestConvert(t *testing.T) {
	tests := []struct {
		value    float64
		fromUnit string
		toUnit   string
		expected float64
	}{
		// Celsius ↔ Fahrenheit
		{0, "Celsius", "Fahrenheit", 32},
		{100, "Celsius", "Fahrenheit", 212},
		{32, "Fahrenheit", "Celsius", 0},
		{212, "Fahrenheit", "Celsius", 100},

		// Celsius ↔ Kelvin
		{0, "Celsius", "Kelvin", 273.15},
		{100, "Celsius", "Kelvin", 373.15},
		{273.15, "Kelvin", "Celsius", 0},
		{373.15, "Kelvin", "Celsius", 100},

		// Fahrenheit ↔ Kelvin
		{32, "Fahrenheit", "Kelvin", 273.15},
		{212, "Fahrenheit", "Kelvin", 373.15},
		{273.15, "Kelvin", "Fahrenheit", 32},
		{373.15, "Kelvin", "Fahrenheit", 212},

		// Identity conversions
		{25, "Celsius", "Celsius", 25},
		{77, "Fahrenheit", "Fahrenheit", 77},
		{300, "Kelvin", "Kelvin", 300},
	}

	for _, tt := range tests {
		got, _ := Convert(tt.value, tt.fromUnit, tt.toUnit)
		if !almostEqual4(got, tt.expected) {
			t.Errorf("Convert(%v, %s, %s) = %v; want %v",
				tt.value, tt.fromUnit, tt.toUnit, got, tt.expected)
		}
	}
}

func TestConvertLength(t *testing.T) {
	tests := []struct {
		value    float64
		fromUnit string
		toUnit   string
		expected float64
	}{
		// Meter ↔ Millimeter
		{1, "meter", "millimeter", 1000},
		{1000, "millimeter", "meter", 1},

		// Meter ↔ Centimeter
		{1, "meter", "centimeter", 100},
		{250, "centimeter", "meter", 2.5},

		// Meter ↔ Kilometer
		{1000, "meter", "kilometer", 1},
		{1.2, "kilometer", "meter", 1200},

		// Meter ↔ Inch
		{1, "meter", "inch", 39.3701},
		{12, "inch", "meter", 0.3048},

		// Meter ↔ Foot
		{1, "meter", "foot", 3.2808},
		{3, "foot", "meter", 0.9144},

		// Meter ↔ Yard
		{1, "meter", "yard", 1.0936},
		{2, "yard", "meter", 1.8288},
		{1, "yard", "foot", 3},
		{1, "yard", "inch", 36},

		// Meter ↔ Mile
		{1609.34, "meter", "mile", 1},
		{1, "mile", "meter", 1609.34},

		// Identity conversions
		{123, "meter", "meter", 123},
		{456, "inch", "inch", 456},
	}

	for _, tt := range tests {
		got, _ := Convert(tt.value, tt.fromUnit, tt.toUnit)
		if !almostEqual4(got, tt.expected) {
			t.Errorf("Convert(%v, %s, %s) = %v; want %v",
				tt.value, tt.fromUnit, tt.toUnit, got, tt.expected)
		}
	}
}

func TestConvertMass(t *testing.T) {
	tests := []struct {
		value    float64
		fromUnit string
		toUnit   string
		expected float64
	}{
		// Kilogram ↔ Gram
		{1, "kilogram", "gram", 1000},
		{2500, "gram", "kilogram", 2.5},

		// Kilogram ↔ Milligram
		{1, "kilogram", "milligram", 1_000_000},
		{500_000, "milligram", "kilogram", 0.5},

		// Kilogram ↔ Ounce
		{1, "kilogram", "ounce", 35.274},
		{16, "ounce", "kilogram", 0.454},

		// Kilogram ↔ Pound
		{1, "kilogram", "pound", 2.205},
		{10, "pound", "kilogram", 4.536},

		// Gram ↔ Milligram
		{1, "gram", "milligram", 1000},
		{2000, "milligram", "gram", 2},

		// Ounce ↔ Pound
		{16, "ounce", "pound", 1},
		{2, "pound", "ounce", 32},

		// Identity conversions
		{123, "kilogram", "kilogram", 123},
		{456, "ounce", "ounce", 456},
	}

	for _, tt := range tests {
		got, _ := Convert(tt.value, tt.fromUnit, tt.toUnit)
		if !almostEqual3(got, tt.expected) {
			t.Errorf("Convert(%v, %s, %s) = %v; want %v",
				tt.value, tt.fromUnit, tt.toUnit, got, tt.expected)
		}
	}
}
