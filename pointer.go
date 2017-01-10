// Definition of packages in GO
package main

import "fmt"

// Main function in GO, like C
func main() {

	var1 := 0
	return1 := doSomethingWithCopy(var1)
	fmt.Printf("Copy Value - Value: %d, Return: %d.\n", var1, return1)

	var2 := 0
	return2 := doSomethingWithPointer(&var2)
	fmt.Printf("Pointer - Value: %d, Return: %d.\n", var2, return2)


}

func doSomethingWithCopy(val int) int {
	val = 10
	return val
}

func doSomethingWithPointer(val *int) *int {
	*val = 10
	return val
}