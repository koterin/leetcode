package main

import (
	"fmt"
	"math"
)

type coord struct {
	x int
	y int
}

func main() {
	var (
		n int
		x int
		y int
	)

	fmt.Scanln(&n)

	fmt.Println("number of cities:", n)

	coords := make([]coord, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&x, &y)
		coords[i] = coord{x: x, y: y}
	}

	var maxDist int

	fmt.Scanln(&maxDist)

	var (
		cityFrom int
		cityTo   int
	)

	fmt.Scanln(&cityFrom, &cityTo)

	fmt.Println("number of cities:", n)
	fmt.Println("cityFrom: ", cityFrom, ", cityTo: ", cityTo)
	fmt.Println("distance between them: ", countDistanceBetweenCities(coords[cityFrom], coords[cityTo]))
}

func countDistanceBetweenCities(c1, c2 coord) int {
	x := c2.x - c1.x
	y := c2.y - c1.y

	sum := math.Abs(float64(x)) + math.Abs(float64(y))

	return int(sum)
}
