package solution

import (
	"PECS/problem"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInverseMu(t *testing.T) {
	assert.Equal(t, 0.0, inverseMu(0, 10))
	assert.Equal(t, 1.0, inverseMu(1, 1))
	assert.Equal(t, 1555.0, inverseMu(5, 5))
}

func TestSa(t *testing.T) {
	n := 5
	circles, _ := problem.RandomUnitCircles(n)
	results := SimmulatedAnnealing(circles)
	obj := objective(circles)
	objGood := objective(results)
	t.Log(obj, objGood)
	assert.True(t, obj >= objGood)
}

func TestRandomEscape(t *testing.T) {
	assert.Equal(t, 0.0, randomEscapeProbability(10, 10, 1000))
	assert.True(t, 3.7201e-44-randomEscapeProbability(10, 1, 1000) < 1e-3)
}

func TestObjective(t *testing.T) {
	type Circle = problem.Circle
	type Circles = problem.Circles
	type Center = problem.Center
	c1 := Circle{Center{0, 0}, 1}
	c2 := Circle{Center{2, 0}, 1}
	c3 := Circle{Center{0, 2}, 1}
	c4 := Circle{Center{-2, 0}, 1}
	c5 := Circle{Center{0, -2}, 1}
	cs := Circles{c1, c2, c3, c4, c5}
	assert.Equal(t, 6.0*6, objective(cs))
}
