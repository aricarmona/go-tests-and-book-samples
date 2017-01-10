// Definition of packages in GO
package main

// Pattern for import modules, like Python
// This module is used for format and definition of strings
import "fmt"

// Main function in GO, like C
func main() {
	// Double
	var input float64
	fmt.Println("Write the value to double: ")
	fmt.Scanf("%f: ", &input)
	fmt.Println("The double is:", double(input))

	// Fahrenheit
	var fahrenheit float64
	fmt.Println("Write the value in fahrenheit: ")
	fmt.Scanf("%f: ", &fahrenheit)
	fmt.Println("The celsius value is:", fahrenheit2celsius(fahrenheit))

	// Meters
	var foot float64
	fmt.Println("Write the value in foot: ")
	fmt.Scanf("%f: ", &foot)
	fmt.Println("The meters value is:", foot2meter(foot))
}

func double(input float64) float64 {
	return input * 2
}

func fahrenheit2celsius(fahrenheit float64) float64 {
	return ( fahrenheit - 32 ) * 5 / 9
}

func foot2meter(foot float64) float64 {
	const fixedMeter float64 = 0.3048
	return foot * fixedMeter
}