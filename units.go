package main

import (
	"fmt"
)

type UnitCategory string

const (
	Length      UnitCategory = "Length"      // base = m
	Mass        UnitCategory = "Mass"        // base = kg
	Temperature UnitCategory = "Temperature" // base = K
)

type Unit struct {
	Name     string
	Category UnitCategory
	ToBase   func(float64) float64
	FromBase func(float64) float64
	Symbol   string
}

var units = map[string]Unit{
	// Length
	"meter": { // basic
		Category: Length,
		ToBase:   func(v float64) float64 { return v },
		FromBase: func(v float64) float64 { return v },
		Symbol:   "m"},
	"millimeter": {
		Category: Length,
		ToBase:   func(v float64) float64 { return v * 0.001 },
		FromBase: func(v float64) float64 { return v / 0.001 },
		Symbol:   "mm"},
	"centimeter": {
		Category: Length,
		ToBase:   func(v float64) float64 { return v * 0.01 },
		FromBase: func(v float64) float64 { return v / 0.01 },
		Symbol:   "cm"},
	"kilometer": {
		Category: Length,
		ToBase:   func(v float64) float64 { return v * 1000 },
		FromBase: func(v float64) float64 { return v / 1000 },
		Symbol:   "km"},
	"inch": {
		Category: Length,
		ToBase:   func(v float64) float64 { return v * 0.0254 },
		FromBase: func(v float64) float64 { return v / 0.0254 },
		Symbol:   "in"},
	"foot": {
		Category: Length,
		ToBase:   func(v float64) float64 { return v * 0.3048 },
		FromBase: func(v float64) float64 { return v / 0.3048 },
		Symbol:   "ft"},
	"yard": {
		Category: Length,
		ToBase:   func(v float64) float64 { return v * 0.9144 },
		FromBase: func(v float64) float64 { return v / 0.9144 },
		Symbol:   "yd"},
	"mile": {
		Category: Length,
		ToBase:   func(v float64) float64 { return v * 1609.34 },
		FromBase: func(v float64) float64 { return v / 1609.34 },
		Symbol:   "mi"},
	// Mass
	"kilogram": { // basic
		Category: Mass,
		ToBase:   func(v float64) float64 { return v },
		FromBase: func(v float64) float64 { return v },
		Symbol:   "kg"},
	"gram": {
		Category: Mass,
		ToBase:   func(v float64) float64 { return v * 0.001 },
		FromBase: func(v float64) float64 { return v / 0.001 },
		Symbol:   "g"},
	"milligram": {
		Category: Mass,
		ToBase:   func(v float64) float64 { return v * 0.000_001 },
		FromBase: func(v float64) float64 { return v / 0.000_001 },
		Symbol:   "mg"},
	"ounce": {
		Category: Mass,
		ToBase:   func(v float64) float64 { return v / 35.274 },
		FromBase: func(v float64) float64 { return v * 35.274 },
		Symbol:   "oz"},
	"pound": {
		Category: Mass,
		ToBase:   func(v float64) float64 { return v * 0.4536 },
		FromBase: func(v float64) float64 { return v / 0.4536 },
		Symbol:   "lb"},

	// Temperature
	"Kelvin": { // basic
		Category: Temperature,
		ToBase:   func(v float64) float64 { return v },
		FromBase: func(v float64) float64 { return v },
		Symbol:   "K"},
	"Celsius": {
		Category: Temperature,
		ToBase:   func(v float64) float64 { return v + 273.15 },
		FromBase: func(v float64) float64 { return v - 273.15 },
		Symbol:   "°C"},
	"Fahrenheit": {
		Category: Temperature,
		ToBase:   func(v float64) float64 { return (v-32)*5/9 + 273.15 },
		FromBase: func(v float64) float64 { return (v-273.15)*9/5 + 32 },
		Symbol:   "°F"},
}

func getNames(category UnitCategory) []string {
	var result []string
	for key, unit := range units {
		if unit.Category == category {
			result = append(result, key)
		}
	}
	return result
}

func getSymbol(name string) string {
	unit := units[name]
	return unit.Symbol
}

func Convert(value float64, unitFrom, unitTo string) (float64, error) {
	from, ok1 := units[unitFrom]
	to, ok2 := units[unitTo]
	if !ok1 || !ok2 {
		return 0, fmt.Errorf("unknown unit(s)")
	}
	if from.Category != to.Category {
		return 0, fmt.Errorf("incompatible categories: %s != %s", from.Category, to.Category)
	}

	baseValue := from.ToBase(value)
	result := to.FromBase(baseValue)
	return result, nil
}
