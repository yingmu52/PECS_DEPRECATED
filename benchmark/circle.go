package benchmark

type circle struct {
	center
	radius float64
}

type circles []circle

func (c circle) square() retangle {
	return retangle{
		c.center.x - c.radius,
		c.center.x + c.radius,
		c.center.y - c.radius,
		c.center.y + c.radius,
	}
}

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

func (cs circles) energy() float64 {
	e := float64(0)
	for _, c := range cs {
		for _, t := range cs {
			e += c.center.distance(t.center)
		}
	}
	return e
}
