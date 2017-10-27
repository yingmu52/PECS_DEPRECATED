package benchmark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCenter(t *testing.T) {
	c1 := center{0, 0}
	c2 := center{0, 100}
	dist := c1.distance(c2)
	assert.Equal(t, dist, 100.0, "error distance")
}

func TestCircleContainer(t *testing.T) {
	container1 := unitSquareContainer()
	container2 := container1.randCircles(100).container()
	assert.True(t, container2.left >= container1.left, "left error")
	assert.True(t, container2.right <= container1.right, "right error")
	assert.True(t, container2.top <= container1.top, "top error")
	assert.True(t, container2.bottom >= container1.bottom, "bottom error")
}

func TestEnergy(t *testing.T) {
	c1 := circle{center{0, 0}, 1}
	c2 := circle{center{2, 0}, 1}
	c3 := circle{center{0, 2}, 1}
	c4 := circle{center{-2, 0}, 1}
	c5 := circle{center{0, -2}, 1}
	cs := circles{c1, c2, c3, c4, c5}
	assert.True(t, cs.energy() == 0.0, "energy error")

	cs1 := circles{}
	for i := 0; i != 2; i++ {
		cs1 = append(cs1, circle{center{0, 0}, 1})
	}
	t.Logf("energy %f", cs1.energy())
	t.Log(cs1)
	assert.True(t, cs1.energy() == 4, "energy error")
}
