package engine

// Camera is our window to rendering scene
type Camera struct {
	BottomLeft Vec3
	Horisontal Vec3
	Vertical   Vec3
	Origin     Vec3
}

// Ray returns ray within camera u and v
func (c Camera) Ray(u, v float64) *Ray {
	return &Ray{
		Origin:    c.Origin,
		Direction: c.BottomLeft.Add(c.Horisontal.MulF(u)).Add(c.Vertical.MulF(v)),
	}
}
