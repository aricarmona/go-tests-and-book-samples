// Definition of packages in GO
package main

import (
	"fmt"
	"math"
)

// Main function in GO, like C
func main() {
	// Average
	list := []float64{11,2,3,4,5,65,2}
	fmt.Printf("Value: %.2f\n", average(list))

	// Check and multiple returns
	sum, err := sumPositiveValues(5.0, 4.0)
	if !err {
		fmt.Println("Error", sum)
	} else {
		fmt.Println("OK", sum)
	}

	// Multiple parameters
	sum2, err := sumPositiveIntegers(11,2,3,4,5,65,2,3,2)
	if !err {
		fmt.Println("Error", sum2)
	} else {
		fmt.Println("OK", sum2)
	}

	// Recursive functions
	fmt.Println("Factorial: ", factorial(12))

	// Calc half number
	num3 := 4.0
	halfNum, even := half(num3)
	fmt.Println(num3, halfNum, even)

	// The greater values
	fmt.Printf("The greater value is %d\n", findTheGreaterNumber(1,4,5,101,3,6,7,8,77,6,5,90,3,54))

	// Get some odd values
	fmt.Println(makeOddGenerator(100))
}

func average(vals []float64) float64 {
	total := 0.0

	for _, val := range vals {
		total += val
	}

	return total / float64(len(vals))
}

func sumPositiveValues(a float64, b float64) (float64, bool) {
	sum := 0.0

	if a < 0 || b < 0 {
		return sum, false
	}

	sum += a + b

	return sum, true
}

func sumPositiveIntegers(list...int) (int, bool) {
	sum := 0

	for _, val := range list {
		if val < 0 {
			return 0, false
		}

		sum += val
	}

	return sum, true
}

func factorial(x uint64) uint64 {
	if x == 0 {
		return 1
	}

	return x * factorial( x - 1 )
}

func half(num float64) (float64, bool) {
	return num / 2, isEven(num)
}

func isEven(num float64) bool {
	return math.Mod(num, 2) == 0
}

func findTheGreaterNumber(values...int) int {
	greater := 0

	for _, value := range values {
		if value > greater {
			greater = value
		}
	}

	return greater
}

func makeOddGenerator(quantity int) []int {
	var list []int

	for i := 1; len(list) < quantity; i++ {
		if math.Mod(float64(i), 2) != 0 {
			list = append(list, i)
		}
	}

	return list
}