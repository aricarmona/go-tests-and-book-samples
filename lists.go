// Definition of packages in GO
package main

// Pattern for import modules, like Python
// This module is used for format and definition of strings
import "fmt"

// Main function in GO, like C
func main() {
	workingWithArrays()
	fmt.Println("\n\n")

	workingWithArraysAdvanced()
	fmt.Println("\n\n")

	workingWithSlices()
	fmt.Println("\n\n")

	workingWithMaps()
}

func workingWithArrays() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	// Percorrendo os itens de outra maneira
	for _, value := range x {
		fmt.Println(value)
	}
}

func workingWithArraysAdvanced() {
	x := [5]float64{ 12, 123, 453, 456, 567 }

	// Percorrendo os itens de outra maneira
	for _, value := range x {
		fmt.Println(value)
	}
}

func workingWithSlices() {
	x := make([]int, 10)

	// Percorrendo os itens de outra maneira
	for i := 0; i < 10; i++ {
		x[i] = i
	}

	fmt.Println(x)
}

func workingWithMaps() {
	// Create a map
	values := make(map[string]int)

	// Define some values to map
	values["first"] = 1
	values["second"] = 2
	fmt.Println(values)

	// Check if a specific value exists
	name, ok := values["first"]
	fmt.Println(name, ok)

	// Remove an element from map
	delete(values, "first")
	fmt.Println(values)

	// Check again if a specific value exists
	// In Go there is a way for detect if an elements exists in Map. We can do that thowht getting value and second value
	// The second value detect if the index was found in Map. It will return a boolen value
	name2, ok2 := values["first"]
	fmt.Println(name2, ok2)

	// Other way for define big maps
	elements := map[string]string {
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"C": "Carbon",
	}

	fmt.Println( elements )

	if el, ok := elements["Li"]; ok {
		fmt.Println(el)
	}

	if el, ok := elements["B"]; ok {
		fmt.Println(el)
	}
}

func check() bool {
	return true
}
