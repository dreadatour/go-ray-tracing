package engine

import "math"

// Sphere is 3D sphere
type Sphere struct {
	Center   Vec3
	Radius   float64
	Material Material
}

// check if sphere is hittable
var _ Hitable = &Sphere{}

// Hit returns closest hit
func (s Sphere) Hit(ray *Ray, tMin float64, tMax float64) (*Hit, bool) {
	var oc = ray.Origin.Sub(s.Center)
	var a = Dot(ray.Direction, ray.Direction)
	var b = Dot(oc, ray.Direction)
	var c = Dot(oc, oc) - s.Radius*s.Radius
	var d = b*b - a*c
	if d <= 0 {
		return nil, false
	}

	var t = (-b - math.Sqrt(d)) / a
	if t > tMin && t < tMax {
		var p = ray.PointAt(t)
		return &Hit{
			T:        t,
			P:        p,
			N:        p.Sub(s.Center).DivF(s.Radius),
			Material: s.Material,
		}, true
	}

	t = (-b + math.Sqrt(d)) / a
	if t > tMin && t < tMax {
		var p = ray.PointAt(t)
		return &Hit{
			T:        t,
			P:        p,
			N:        p.Sub(s.Center).DivF(s.Radius),
			Material: s.Material,
		}, true
	}

	return nil, false
}
