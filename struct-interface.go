// Definition of packages in GO
package main

import (
	"fmt"
	"math"
)

// Main function in GO, like C
func main() {
	// Circle
	c := Circle{x: 1, y: 2, r: 3}
	fmt.Println(c)
	fmt.Println(c.x)
	fmt.Println(c.area())

	// Rectangle
	r := Rectangle{1,2,5,4}
	fmt.Println(r)
	fmt.Println(r.area())

	// Listing
	l := Listing{}
	fmt.Println(l)
}

type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1 float64
	y1 float64
	x2 float64
	y2 float64
}

func (r *Rectangle) area() float64 {
	return ( r.x2 - r.x1 ) * ( r.y2 - r.y1 )
}

type Address struct {
	state string
	city string
	neighborhood string
	street string
	number string
	zipCode string
}

type Features struct {
	Barbecue bool
	Garden bool
	Pool bool
}

type Property struct {
	Address Address
	Features Features
	TotalArea int
	UtilArea int
	Floor int
	Active bool
}

type Development struct {
	Address Address
	Features Features
	TotalArea int
	Active bool
}

type Listing struct {
	Property Property
	Price float64
}

type Name struct {
	firstName string
	middleName string
	lastName string
	nickName string
}

type Person struct {
	Name Name
	Address Address
	Age int
	Gender string
}