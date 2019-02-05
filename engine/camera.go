package engine

import "math"

// Camera is our window to rendering scene
type Camera struct {
	BottomLeft Vec3
	Horisontal Vec3
	Vertical   Vec3
	Origin     Vec3
}

// NewCamera returns new Camera with given hfov (in degrees) and aspect ratio
func NewCamera(lookFrom, lookAt, vup Vec3, hfov, aspect float64) Camera {
	var theta = hfov * math.Pi / 180
	var halfWidth = math.Tan(theta / 2)
	var halfHeight = halfWidth / aspect
	var w = lookFrom.Sub(lookAt).UnitV()
	var u = Cross(vup, w).UnitV()
	var v = Cross(w, u)
	return Camera{
		BottomLeft: lookFrom.Sub(u.MulF(halfWidth)).Sub(v.MulF(halfHeight)).Sub(w),
		Horisontal: u.MulF(2 * halfWidth),
		Vertical:   v.MulF(2 * halfHeight),
		Origin:     lookFrom,
	}
}

// Ray returns ray within camera u and v
func (c Camera) Ray(u, v float64) *Ray {
	return &Ray{
		Origin:    c.Origin,
		Direction: c.BottomLeft.Add(c.Horisontal.MulF(u)).Add(c.Vertical.MulF(v)).Sub(c.Origin),
	}
}
