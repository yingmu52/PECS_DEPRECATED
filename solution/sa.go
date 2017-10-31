package solution

import (
	"PECS/problem"
	"math"
)

// SimmulatedAnnealing implemnts https://en.wikipedia.org/wiki/Simulated_annealing
func SimmulatedAnnealing(cs problem.Circles) problem.Circles {
	mMax := 500
	output := cs
	current := output
	f := objective(output)
	fCurrent := f

	for m := 0; m < mMax; m++ {
		t := float64(m / mMax)
		for k := 0; k < mMax; k++ {
			tmpOutput, tmpF := neighbour(current, t)
			df := tmpF - fCurrent
			//use Metropolis condition to accept or reject the test point as current point.
			if df < 0 || problem.RandomFloat(0.0, 1.0) < randomEscapeProbability(df, t, f) {
				current = tmpOutput
				fCurrent = tmpF
			}
			if fCurrent < f {
				output = current
				f = fCurrent
			}
		}
	}
	return output
}

func neighbour(cs problem.Circles, t float64) (problem.Circles, float64) { // current neighbour under current tempreture
	neighbour := problem.Circles{}
	bound := cs.Container()
	randCircles, _ := problem.RandomUnitCircles(len(cs))
	mu := math.Pow(10.0, t*100.0)
	for i, c := range randCircles {
		c.Center.X += inverseMu(cs[i].Center.X, mu)
		c.Center.Y += inverseMu(cs[i].Center.Y, mu)
		c = c.ForceInbound(bound)
		neighbour = append(neighbour, c)
	}
	return cs, objective(cs)
}

func inverseMu(i, mu float64) float64 {
	sign := 0.0
	if i != 0 {
		sign = i / math.Abs(i)
	}
	return (math.Pow(1+mu, math.Abs(i)) - 1) / mu * sign
}

func randomEscapeProbability(df, t, f float64) float64 {
	tolFun := float64(1e-4)
	esp := float64(1e-7)
	return math.Exp(-t * df / (math.Abs(f) + esp) / tolFun)
}

func objective(cs problem.Circles) float64 {
	s := cs.Container().ToSquare().Width
	return math.Pow(s, 2) + math.Pow(s, 3)*cs.Energy()
}
