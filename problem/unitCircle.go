package problem

// RandomUnitCircles creates n random circles in retangle
func RandomUnitCircles(n int) (Circles, Retangle) {
	r := Retangle{0, float64(n * 2), 0, float64(n * 2)}
	radius := float64(1)
	cs := Circles{}
	for i := 0; i < n; i++ {
		randCircle := Circle{
			Center{
				RandomFloat(r.left+radius, r.right-radius),
				RandomFloat(r.bottom+radius, r.top-radius),
			},
			radius,
		}
		cs = append(cs, randCircle)
	}
	return cs, r
}

// ConvertToCanvas converts array of circle to Canvas of with l
// Note: Canvas must be square
func (cs Circles) ConvertToCanvas(l float64) Circles {
	output := Circles{}
	s := cs.Container()
	ratio := float64(l) / s.ToSquare().Width
	for _, c := range cs {
		// make sure all circles are inside square with left-bottom at center of coordinate system
		c.Center.X = (c.Center.X - s.left) * ratio
		c.Center.Y = (c.Center.Y - s.bottom) * ratio
		c.Radius *= ratio
		output = append(output, c)
	}
	return output
}
