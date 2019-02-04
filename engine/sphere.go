package engine

// Sphere is 3D sphere
type Sphere struct {
	Center Vec3
	Radius float64
}

// RayIntersect checks if ray intersects with sphere
func (s *Sphere) RayIntersect(ray Ray) bool {
	var oc = ray.Origin.Sub(s.Center)
	var a = Dot(ray.Direction, ray.Direction)
	var b = float64(2) * Dot(oc, ray.Direction)
	var c = Dot(oc, oc) - s.Radius*s.Radius
	var d = b*b - float64(4)*a*c
	return d > 0
}
