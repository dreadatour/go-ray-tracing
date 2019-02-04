package engine

import "math"

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

// RayHit checks if ray intersects with sphere and returns normal length
func (s *Sphere) RayHit(ray Ray) float64 {
	var oc = ray.Origin.Sub(s.Center)
	var a = Dot(ray.Direction, ray.Direction)
	var b = float64(2) * Dot(oc, ray.Direction)
	var c = Dot(oc, oc) - s.Radius*s.Radius
	var d = b*b - float64(4)*a*c
	if d < 0 {
		return -1
	}
	return (-b - math.Sqrt(d)) / (2 * a)
}
