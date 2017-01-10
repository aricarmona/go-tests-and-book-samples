package main

import "fmt"

func main() {
	fmt.Println("TYPES\n\n")
	fmt.Println("1 + 1 = ", 1 + 1)
	fmt.Println("1 + 1 = ", 1.0 + 1.0)
	fmt.Println( "\n\n")

	// Strings
	fmt.Println( len("Hello") )
	fmt.Println("Hello, " + "World")
	fmt.Println("Hello"[0])
	fmt.Println("\n\n")

	// Boolean
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!false)
}