package problem

import (
	"math"
	"math/rand"
)

// Center ...
type Center struct {
	X, Y float64
}

func (c Center) distance(t Center) float64 {
	return math.Sqrt(math.Pow(c.X-t.X, 2) + math.Pow(c.Y-t.Y, 2))
}

// Circle ...
type Circle struct {
	Center
	Radius float64
}

func (c Circle) square() Retangle {
	return Retangle{
		c.Center.X - c.Radius,
		c.Center.X + c.Radius,
		c.Center.Y - c.Radius,
		c.Center.Y + c.Radius,
	}
}

// Circles is an array of circle
type Circles []Circle

func (cs Circles) container() Retangle {
	rect := cs[0].square()
	for _, c := range cs {
		r := c.square()
		if rect.left > r.left {
			rect.left = r.left
		}
		if rect.right < r.right {
			rect.right = r.right
		}
		if rect.top < r.top {
			rect.top = r.top
		}
		if rect.bottom > r.bottom {
			rect.bottom = r.bottom
		}
	}
	return rect
}

type index struct {
	i, j int
}

// Energy measures the overlapping status of circles
func (cs Circles) Energy() float64 {
	e := float64(0)
	n := len(cs)
	memo := map[index]float64{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// avoid checking self
			if i == j {
				continue
			}
			// avoid checking same path twice
			ind := index{i, j}
			_, exist := memo[ind]
			if exist {
				continue
			}
			// calculate energy and cache it to memo
			c := cs[i]
			t := cs[j]
			dist := c.Center.distance(t.Center) - c.Radius - t.Radius
			if dist > 0 {
				continue
			}
			cache := math.Pow(dist, 2)
			memo[ind] = cache
			indReverse := index{j, i}
			memo[indReverse] = cache
			e += cache
		}
	}
	return e
}

// Objective is objective function
func (cs Circles) Objective() float64 {
	s := cs.container().toSquare().Width
	return math.Pow(s, 2) + math.Pow(s, 3)*cs.Energy()
}

// Square ...
type Square struct {
	Center
	Width float64
}

// Retangle ...
type Retangle struct {
	left, right, bottom, top float64
}

func (s Square) toRectangle() Retangle {
	return Retangle{
		s.Center.X - s.Width/2,
		s.Center.X + s.Width/2,
		s.Center.Y - s.Width/2,
		s.Center.Y + s.Width/2,
	}
}

// UnitSquare gives a square at center
func UnitSquare() Retangle {
	return Retangle{-0.5, 0.5, -0.5, 0.5}
}

func (r Retangle) toSquare() Square {
	w := r.right - r.left
	h := r.top - r.bottom
	Width := math.Max(w, h)
	return Square{
		Center{r.left + Width/2, r.bottom + Width/2},
		Width,
	}
}

// RandCircles spreads n circles in retangle
func (r Retangle) RandCircles(n int) Circles {
	cs := Circles{}
	radius := float64((r.right - r.left) / float64(n))
	for i := 0; i < n; i++ {
		randCircle := Circle{
			Center{
				randomFloat(r.left+radius, r.right-radius),
				randomFloat(r.bottom+radius, r.top-radius),
			},
			radius,
		}
		cs = append(cs, randCircle)
	}
	return cs
}

func randomFloat(min float64, maX float64) float64 {
	return rand.Float64()*(maX-min) + min
}
