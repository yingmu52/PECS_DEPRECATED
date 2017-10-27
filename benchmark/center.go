package benchmark

import (
	"math"
)

type center struct {
	x, y float64
}

func (c center) distance(t center) float64 {
	return math.Sqrt(math.Pow(c.x-t.x, 2) + math.Pow(c.y-t.y, 2))
}
