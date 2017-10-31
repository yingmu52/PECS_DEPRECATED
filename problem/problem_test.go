package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCenter(t *testing.T) {
	c1 := Center{0, 0}
	c2 := Center{0, 100}
	dist := c1.distance(c2)
	assert.Equal(t, 100.0, dist)
}

func TestEnergy(t *testing.T) {
	c1 := Circle{Center{0, 0}, 1}
	c2 := Circle{Center{2, 0}, 1}
	c3 := Circle{Center{0, 2}, 1}
	c4 := Circle{Center{-2, 0}, 1}
	c5 := Circle{Center{0, -2}, 1}
	cs := Circles{c1, c2, c3, c4, c5}
	assert.True(t, cs.Energy() == 0.0)

	cs1 := Circles{}
	for i := 0; i != 2; i++ {
		cs1 = append(cs1, Circle{Center{0, 0}, 1})
	}
	assert.True(t, cs1.Energy() == 4)
}

func TestRectSquareConvert(t *testing.T) {
	c := UnitSquare
	s := c.ToSquare()
	assert.Equal(t, Center{0.5, 0.5}, s.Center)

	s1 := Square{Center{0.5, 0.5}, 1}
	c1 := s1.ToRectangle()
	assert.Equal(t, s1.Center, c1.ToSquare().Center)
}
func TestRandom(t *testing.T) {
	a := []float64{}
	b := []float64{}
	for i := 0; i < 10; i++ {
		a = append(a, RandomFloat(1.0, 10.0))
		b = append(b, RandomFloat(1.0, 10.0))
	}
	for p := 0; p < 10; p++ {
		assert.NotEqual(t, a[p], b[p])
	}

	cs, r := RandomUnitCircles(100)
	for _, c := range cs {
		tmpC := c.ForceInbound(r)
		assert.Equal(t, c, tmpC)
	}
}

func TestForceInbound(t *testing.T) {
	r := SquareOfWidth(6).ToRectangle()
	c0 := Circle{Center{15, 3}, 1}
	c1 := Circle{Center{3, -15}, 1}
	c2 := Circle{Center{3, 15}, 1}
	c3 := Circle{Center{-15, 3}, 1}

	c00 := c0.ForceInbound(r)
	c11 := c1.ForceInbound(r)
	c22 := c2.ForceInbound(r)
	c33 := c3.ForceInbound(r)

	e0 := Circle{Center{5, 3}, 1}
	e1 := Circle{Center{3, -1}, 1}
	e2 := Circle{Center{3, 5}, 1}
	e3 := Circle{Center{1, 3}, 1}

	assert.Equal(t, e0.Center, c00.Center)
	assert.Equal(t, e1.Center, c11.Center)
	assert.Equal(t, e2.Center, c22.Center)
	assert.Equal(t, e3.Center, c33.Center)
}

func TestConvertToCangas(t *testing.T) {
	canvas := float64(400)
	cs, rect := RandomUnitCircles(5)
	cs1 := cs.ConvertToCanvas(canvas)
	rect1 := cs1.Container()
	s := rect.ToSquare()
	s1 := rect1.ToSquare()
	ratio := s1.Width / s.Width
	tolerance := 1e-4
	// (square) container must be scaled correctly
	assert.True(t, s1.Center.X/s.Center.X-ratio < tolerance)
	assert.True(t, s1.Center.Y/s.Center.Y-ratio < tolerance)
}
