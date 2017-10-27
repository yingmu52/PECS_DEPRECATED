package benchmark

import "math/rand"

type square struct {
	center
	width float64
}

type retangle struct {
	left, right, bottom, top float64
}

func unitSquareContainer() retangle {
	return retangle{-0.5, 0.5, -0.5, 0.5}
}
func (r retangle) randCircles(n int) circles {
	cs := circles{}
	radius := float64((r.right - r.left) / float64(n))
	for i := 0; i < n; i++ {
		randCircle := circle{
			center{
				randomFloat(r.left+radius, r.right-radius),
				randomFloat(r.bottom+radius, r.top-radius),
			},
			radius,
		}
		cs = append(cs, randCircle)
	}
	return cs
}

func randomFloat(min float64, max float64) float64 {
	return rand.Float64()*(max-min) + min
}
