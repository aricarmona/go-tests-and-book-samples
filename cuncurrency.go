// Definition of packages in GO
package main

import (
	"fmt"
	"time"
	"math/rand"
)

// Main function in GO, like C
func main() {

	for i := 0; i < 10; i++ {
		go test(i)
	}

	var input string
	fmt.Scanln(&input)
}

func test(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}