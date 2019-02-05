package engine

import "math"

// Sphere is 3D sphere
type Sphere struct {
	Center   Vector
	Radius   float64
	Material Material
}

// check if sphere is hittable
var _ Hitable = &Sphere{}

// Hit returns closest hit
func (s Sphere) Hit(ray *Ray, dMin float64, dMax float64) (*Hit, bool) {
	var oc = ray.Origin.Sub(s.Center)
	var a = Dot(ray.Direction, ray.Direction)
	var b = Dot(oc, ray.Direction)
	var c = Dot(oc, oc) - s.Radius*s.Radius
	var d = b*b - a*c
	if d <= 0 {
		return nil, false
	}

	var distance = (-b - math.Sqrt(d)) / a
	if distance > dMin && distance < dMax {
		var point = ray.PointAt(distance)
		return &Hit{
			Distance: distance,
			Point:    point,
			Normal:   point.Sub(s.Center).DivS(s.Radius),
			Material: s.Material,
		}, true
	}

	distance = (-b + math.Sqrt(d)) / a
	if distance > dMin && distance < dMax {
		var point = ray.PointAt(distance)
		return &Hit{
			Distance: distance,
			Point:    point,
			Normal:   point.Sub(s.Center).DivS(s.Radius),
			Material: s.Material,
		}, true
	}

	return nil, false
}
