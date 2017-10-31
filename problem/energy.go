package problem

import "math"

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
	s := cs.Container().ToSquare().Width
	return math.Pow(s, 2) + math.Pow(s, 3)*cs.Energy()
}
