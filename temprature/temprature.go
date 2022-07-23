package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	startConversion()
}

type Temperature struct {
	units []Conversion
}

// Add all units to slice
func (t *Temperature) addTemperatureUnit(n Conversion) {
	t.units = append(t.units, n)
}

// Conversion occurs in this method
func (t *Temperature) conversion(value float64, unit, goalUnit string) (float64, error) {
	var result float64
	// Check contains or not
	for _, f := range t.units {
		if strings.ToLower(unit) == strings.ToLower(f.GetUnit()) {
			result, _ = f.Convert(value, strings.ToLower(goalUnit))
			return result, nil
		}
	}
	// If not contains
	return 0, errors.New("invalid conversion")
}

func startConversion() {
	// Create empty struct and add all units
	temperatureConverter := Temperature{}
	temperatureConverter.addTemperatureUnit(Celsius{"C"})
	temperatureConverter.addTemperatureUnit(Fahrenheit{"F"})
	temperatureConverter.addTemperatureUnit(Kelvin{"K"})

	fmt.Println("\nTemprature converter started.")
	for {
		var valueWithUnit string
		var unit string
		var goalUnit string

		fmt.Println("> Input the actual temprature with unit (like 15C) or enter x to exit: ")
		// Get value as text
		_, err := fmt.Scanln(&valueWithUnit)

		// Error or exit condition
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		} else if strings.ToLower(valueWithUnit) == "x" {
			fmt.Println("Goodbye!")
			os.Exit(0)
		}

		// Get value
		value, err := strconv.ParseFloat(valueWithUnit[:len(valueWithUnit)-1], 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// Get unit
		unit = valueWithUnit[len(valueWithUnit)-1:]

		fmt.Println("> Enter unit to transform:")
		_, err = fmt.Scanln(&goalUnit)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		// Get unit which to be converted
		goalValue, err := temperatureConverter.conversion(value, unit, goalUnit)
		if err != nil {
			fmt.Println(err)
		} else {
			// Print result
			fmt.Printf("%.2f%s is equal to %.2f%s\n\n", value, strings.ToUpper(unit), goalValue, strings.ToUpper(goalUnit))
		}
	}
}
