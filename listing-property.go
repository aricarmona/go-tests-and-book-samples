// Definition of packages in GO
package main

import "fmt"

// Import listing package. There are some struct in there
import l "listing"

// Main function in GO, like C
func main() {
	// Listing
	p := l.Property{}

	// Populate with new values
	p.Populate()

	// Print object with all values
	fmt.Println(p)
}
