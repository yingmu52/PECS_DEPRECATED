package benchmark

import (
	"math"
)

type circle struct {
	center
	radius float64
}

func (c circle) square() retangle {
	return retangle{
		c.center.x - c.radius,
		c.center.x + c.radius,
		c.center.y - c.radius,
		c.center.y + c.radius,
	}
}

type circles []circle

func (cs circles) container() retangle {
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

func (cs circles) energy() float64 {
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
			dist := c.center.distance(t.center) - c.radius - t.radius
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
