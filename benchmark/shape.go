package benchmark

// Point represents a coordinates in Cartesian Coordinate System
type Point struct {
	X, Y float64
}

// Square represents a container or a damaged area
type Square struct {
	Point
	Width float64
}

// Circle represents a circle to be packed
type Circle struct {
	Point
	Radius float64
}

// Container is a square at 0,0 with side equals to 1
func Container() Square {
	return Square{Point{0, 0}, 1}
}
