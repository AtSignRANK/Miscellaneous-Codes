package main

import (
	"fmt"
	"math/rand"
)

var points = 0
var points_in_circle = 0

func point() {
	points++
	var x float64 = rand.Float64()
	var y float64 = rand.Float64()

	if x*x+y*y <= 1 {
		points_in_circle++
	}
}

func main() {
	for {
		point()
		fmt.Printf("Pi is: %.10f\n", 4*float32(points_in_circle)/float32(points))
	}
}
