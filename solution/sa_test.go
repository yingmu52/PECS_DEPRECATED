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
	n := 3
	circles, _ := problem.RandomUnitCircles(n)
	results := SimmulatedAnnealing(n)
	obj := circles.Objective()
	objGood := results.Objective()
	t.Log(obj, objGood)
	assert.True(t, obj >= objGood)
}
