package main

import "fmt"

func main() {
	// Variables
	var hello string = "Hello, World"
	fmt.Println(hello)

	var first string = "First"
	var second string = "Second"
	first += " - " + second
	fmt.Println(first)

	var (
		a int = 10
		b = 20
		c = 30
	)
	fmt.Println(a, b, c)

	const name = "Alberto Roberto"
	fmt.Println(name)

	const (
		firstName string = "Alberto"
		lastName = "Roberto"
	)
	fmt.Println(firstName, lastName)
}
