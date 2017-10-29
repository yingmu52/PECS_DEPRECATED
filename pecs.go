package main

import (
	"PECS/problem"
	"PECS/solution"
	"fmt"
)

func main() {
	circles := problem.UnitSquare().RandCircles(10)
	results := solution.SimmulatedAnnealing(circles)
	p(circles)
	p(results)
}

func p(cs problem.Circles) {
	for _, c := range cs {
		fmt.Println(c)
	}
	fmt.Println()
}
