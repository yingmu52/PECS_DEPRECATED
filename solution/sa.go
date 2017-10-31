package solution

import (
	"PECS/problem"
	"fmt"
	"math"
)

// SimmulatedAnnealing implemnts https://en.wikipedia.org/wiki/Simulated_annealing
func SimmulatedAnnealing(n int) problem.Circles {
	mMax := 500
	tolFun := float64(1e-4)
	output, _ := problem.RandomUnitCircles(n)
	f := output.Objective()

	for m := 0; m < mMax; m++ {
		t := float64(m / mMax)
		for k := 0; k < mMax; k++ {
			tmpOutput := neighbour(output, t)
			tmpF := tmpOutput.Objective()
			df := tmpF - f
			//use Metropolis condition to accept or reject the test point as current point.
			if df < 0 || problem.RandomFloat(0.0, 1.0) < math.Exp(-t*df/(math.Abs(f))/tolFun) {
				output = tmpOutput
				f = tmpF
			}
		}
	}
	return output
}

func neighbour(cs problem.Circles, t float64) problem.Circles { // current neighbour under current tempreture
	neighbour := problem.Circles{}
	bound := cs.Container()
	fmt.Println(bound.ToSquare().Width)
	randCircles, _ := problem.RandomUnitCircles(len(cs))
	mu := math.Pow(10.0, t*100.0)
	for i, c := range randCircles {
		c.Center.X += inverseMu(cs[i].Center.X, mu)
		c.Center.Y += inverseMu(cs[i].Center.Y, mu)
		c = c.ForceInbound(bound)
		neighbour = append(neighbour, c)
	}
	return cs
}

func inverseMu(i, mu float64) float64 {
	sign := 0.0
	if i != 0 {
		sign = i / math.Abs(i)
	}
	return (math.Pow(1+mu, math.Abs(i)) - 1) / mu * sign
}
