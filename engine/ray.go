package engine

// Ray with start and origin and direction
type Ray struct {
	Origin    Vector
	Direction Vector
}

// NewRay returns new ray
func NewRay(origin, direction Vector) Ray {
	return Ray{Origin: origin, Direction: direction}
}

// PointAt is point on ray at distance
func (r Ray) PointAt(distance float64) Vector {
	return r.Origin.Add(r.Direction.MulS(distance))
}
