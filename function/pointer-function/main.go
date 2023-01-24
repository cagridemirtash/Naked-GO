package main

import "fmt"

type Coordinates struct {
	x, y float64
}

func main() {
	coordinate := Coordinates{5, 2}
	// Go interpreter set to coordinate like this -> (&coordinate).Scale(5)
	coordinate.Scale(5)
	fmt.Println("x:", coordinate.x, "y:", coordinate.y)
}

// Scale Pointer Receivers
func (coordinate *Coordinates) Scale(tempF float64) {
	coordinate.x = coordinate.x * tempF
	coordinate.y = coordinate.y * tempF
}
