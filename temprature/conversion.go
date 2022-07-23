package main

import "errors"

type Conversion interface {
	Convert(number float64, unit string) (float64, error)
	GetUnit() string
}

type Celsius struct {
	Unit string
}

func (c Celsius) Convert(value float64, unit string) (float64, error) {
	if unit == "f" || unit == "fahrenheit" {
		return (value * (9.0 / 5.0)) + 32, nil
	} else if unit == "k" || unit == "kelvin" {
		return value + 273.15, nil
	}
	return 0, errors.New("invalid conversion")
}

func (c Celsius) GetUnit() string {
	return c.Unit
}

type Fahrenheit struct {
	Unit string
}

func (f Fahrenheit) Convert(value float64, unit string) (float64, error) {
	if unit == "c" || unit == "celsius" {
		return (value - 32) * (5.0 * 9.0), nil
	} else if unit == "k" || unit == "kelvin" {
		return (value-32)*(5.0/9.0) + 273.15, nil
	}
	return 0, errors.New("invalid conversion")
}

func (f Fahrenheit) GetUnit() string {
	return f.Unit
}

type Kelvin struct {
	Unit string
}

func (k Kelvin) Convert(value float64, unit string) (float64, error) {
	if unit == "c" || unit == "celsius" {
		return value - 273.15, nil
	} else if unit == "f" || unit == "fahrenheit" {
		return 1.8*(value-273.15) + 32, nil
	}
	return 0, errors.New("invalid conversion")
}

func (k Kelvin) GetUnit() string {
	return k.Unit
}
