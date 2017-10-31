package problem

import (
	"math"
	"math/rand"
	"time"
)

// In this file, all functions MUST be dimension independent.

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

// ForceInbound force circle to stay inside of retangle
func (c Circle) ForceInbound(r Retangle) Circle {
	if c.Center.X < r.left+c.Radius {
		c.Center.X = r.left + c.Radius
	}
	if c.Center.X > r.right-c.Radius {
		c.Center.X = r.right - c.Radius
	}
	if c.Center.Y < r.bottom-c.Radius {
		c.Center.Y = r.bottom - c.Radius
	}
	if c.Center.Y > r.top-c.Radius {
		c.Center.Y = r.top - c.Radius
	}
	return c
}

// Circles is an array of circle
type Circles []Circle

// Container finds the retangle that encloses the circles
func (cs Circles) Container() Retangle {
	boundary := func(c Circle) Retangle {
		return Retangle{
			c.Center.X - c.Radius,
			c.Center.X + c.Radius,
			c.Center.Y - c.Radius,
			c.Center.Y + c.Radius,
		}
	}
	rect := boundary(cs[0])
	for _, c := range cs {
		r := boundary(c)
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

// Square ...
type Square struct {
	Center
	Width float64
}

// SquareOfWidth creates a square at center (w/2,w/2)
func SquareOfWidth(w float64) Square {
	return Square{Center{w / 2, w / 2}, w}
}

// ToRectangle gets the retangle information from a square
func (s Square) ToRectangle() Retangle {
	return Retangle{
		s.Center.X - s.Width/2,
		s.Center.X + s.Width/2,
		s.Center.Y - s.Width/2,
		s.Center.Y + s.Width/2,
	}
}

// UnitSquare is a bottom-left corner at center
var UnitSquare = Retangle{0, 1, 0, 1}

// Retangle left, right, bottom top
type Retangle struct {
	left, right, bottom, top float64
}

// ToSquare creates a square from a retangle, expanding +y direction
func (r Retangle) ToSquare() Square {
	w := r.right - r.left
	h := r.top - r.bottom
	Width := math.Max(w, h)
	return Square{
		Center{r.left + Width/2, r.bottom + Width/2},
		Width,
	}
}

// RandomFloat creates a randome float64 in (min, maX)
func RandomFloat(min float64, max float64) float64 {
	// the default source won't be dynamic random, need to create new source everytime for SA
	randNewSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randNewSource.Float64()*(max-min) + min
}
