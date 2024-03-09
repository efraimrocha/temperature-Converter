package main

import (
	"fmt"
)

// Converte Celsius para Fahrenheit
func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

// Converte Celsius para Kelvin
func celsiusToKelvin(c float64) float64 {
	return c + 273.15
}

// Converte Fahrenheit para Celsius
func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

// Converte Kelvin para Celsius
func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func main() {

	var temp float64
	var unit string
	fmt.Print("Enter temperature: ")
	fmt.Scan(&temp)
	fmt.Print("Enter temperature unit (C, F, K): ")
	fmt.Scan(&unit)

	switch unit {
	case "c":
		var f = celsiusToFahrenheit(temp)
		var k = celsiusToKelvin(temp)
		fmt.Printf("%g°C, equivale a %gºF e %gK", temp, f, k)
	case "f":
		var c = fahrenheitToCelsius(temp)
		var k = celsiusToKelvin(c)
		fmt.Printf("%g°F, equivale a %gºC e %gK", temp, c, k)
	case "k":
		var c = kelvinToCelsius(temp)
		var f = celsiusToFahrenheit(c)
		fmt.Printf("%g°K, equivale a %gºC e %gF", temp, c, f)
	default:
		fmt.Println("Temperatura ou unidade inválida.")
		return
	}

}
