package main

import (
	"fmt"
	"math"
)

// Area interface
type Area interface {
	getArea() float64
}

// Triangle interface
type Triangle struct {
	base   float64
	height float64
}

// Square interface
type Square struct {
	sideLength float64
}

func main() {

	tri := Triangle{
		base:   2,
		height: 3,
	}
	sq := Square{
		sideLength: 9,
	}

	printArea(tri)
	printArea(sq)
}

func printArea(a Area) {
	fmt.Println(a.getArea())
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s Square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}
