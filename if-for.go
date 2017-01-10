// Definition of packages in GO
package main

// Pattern for import modules, like Python
// This module is used for format and definition of strings
import "fmt"

// Main function in GO, like C
func main() {
	showNumbers()
	fmt.Println("\n\n")

	showPrettyNumbers()
	fmt.Println("\n\n")

	showOnlyNumbersDivisibleForThree()
}

func showNumbers() {
	for i:= 1; i <= 10; i++ {
		fmt.Println(i)
		if i % 2 == 0 {
			fmt.Println("par")
		} else {
			fmt.Println("impar")
		}
	}
}

func showPrettyNumbers() {
	for i := 0; i < 11; i++ {
		switch i {
			case 0: fmt.Println("Zero")
			case 1: fmt.Println("One")
			case 2: fmt.Println("Two")
			case 3: fmt.Println("Three")
			default: fmt.Println(i)
		}
	}
}

func showOnlyNumbersDivisibleForThree() {
	for i:= 1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		}
	}

}