package engine

// Ray with start and origin and direction
type Ray struct {
	Origin    Vec3
	Direction Vec3
}

// NewRay returns new ray
func NewRay(orig, dir Vec3) Ray {
	return Ray{Origin: orig, Direction: dir}
}

// PointAt is point on ray at distance d
func (r Ray) PointAt(d float64) Vec3 {
	return r.Origin.Add(r.Direction.MulF(d))
}
