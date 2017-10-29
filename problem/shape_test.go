package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCenter(t *testing.T) {
	c1 := Center{0, 0}
	c2 := Center{0, 100}
	dist := c1.distance(c2)
	assert.Equal(t, 100.0, dist, "error distance")
}

func TestCircleContainer(t *testing.T) {
	container1 := UnitSquare()
	container2 := container1.RandCircles(100).container()
	assert.True(t, container2.left >= container1.left, "left error")
	assert.True(t, container2.right <= container1.right, "right error")
	assert.True(t, container2.top <= container1.top, "top error")
	assert.True(t, container2.bottom >= container1.bottom, "bottom error")
}

func TestEnergy(t *testing.T) {
	c1 := Circle{Center{0, 0}, 1}
	c2 := Circle{Center{2, 0}, 1}
	c3 := Circle{Center{0, 2}, 1}
	c4 := Circle{Center{-2, 0}, 1}
	c5 := Circle{Center{0, -2}, 1}
	cs := Circles{c1, c2, c3, c4, c5}
	assert.True(t, cs.Energy() == 0.0, "energy error")

	cs1 := Circles{}
	for i := 0; i != 2; i++ {
		cs1 = append(cs1, Circle{Center{0, 0}, 1})
	}
	assert.True(t, cs1.Energy() == 4, "energy error")
}

func TestRectSquareConvert(t *testing.T) {
	c := UnitSquare()
	s := c.toSquare()
	assert.Equal(t, Center{0, 0}, s.Center, "error")

	s1 := Square{Center{0, 0}, 1}
	c1 := s1.toRectangle()
	assert.Equal(t, s1.Center, c1.toSquare().Center, "error")
}

func TestObjective(t *testing.T) {
	c1 := Circle{Center{0, 0}, 1}
	c2 := Circle{Center{2, 0}, 1}
	c3 := Circle{Center{0, 2}, 1}
	c4 := Circle{Center{-2, 0}, 1}
	c5 := Circle{Center{0, -2}, 1}
	cs := Circles{c1, c2, c3, c4, c5}
	assert.Equal(t, 6.0*6, cs.Objective(), "error")
}
