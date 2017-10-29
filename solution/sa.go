package solution

import (
	"PECS/problem"
	"math"
	"math/rand"
)

// SimmulatedAnnealing implemnts https://en.wikipedia.org/wiki/Simulated_annealing
func SimmulatedAnnealing(cs problem.Circles) problem.Circles {
	mMax := 500
	tolFun := float64(1e-4)
	output := make(problem.Circles, len(cs))
	copy(output[:], cs)
	f := output.Objective()

	for m := 0; m < mMax; m++ {
		t := float64(m / mMax)
		for k := 0; k < mMax; k++ {
			randCircles := problem.UnitSquare().RandCircles(len(cs))
			tmpOutput := neighbour(randCircles, t)
			tmpF := tmpOutput.Objective()
			df := tmpF - f
			//use Metropolis condition to accept or reject the test point as current point.
			if df < 0 || rand.Float64() < math.Exp(-t*df/(math.Abs(f))/tolFun) {
				output = tmpOutput
				f = tmpF
			}
			if tmpF < f {
				output = tmpOutput
				f = tmpF
			}
		}
	}
	return output
}

func neighbour(cs problem.Circles, t float64) problem.Circles { // current neighbour under current tempreture
	neighbour := problem.Circles{}
	mu := math.Pow(10.0, t*100.0)
	for _, c := range cs {
		c.Center.X += inverseMu(c.Center.X, mu)
		c.Center.Y += inverseMu(c.Center.Y, mu)
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
