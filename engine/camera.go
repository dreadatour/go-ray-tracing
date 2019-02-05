package engine

import "math"

// Camera is our window to rendering scene
type Camera struct {
	BottomLeft Vec3
	Horisontal Vec3
	Vertical   Vec3
	Origin     Vec3
	U, V, W    Vec3
	LensRadius float64
}

// NewCamera returns new Camera with given hfov (in degrees) and aspect ratio
func NewCamera(lookFrom, lookAt, vup Vec3, hfov, aspect, aperture, focusDist float64) Camera {
	var theta = hfov * math.Pi / 180
	var halfWidth = math.Tan(theta / 2)
	var halfHeight = halfWidth / aspect
	var w = lookFrom.Sub(lookAt).UnitV()
	var u = Cross(vup, w).UnitV()
	var v = Cross(w, u)
	return Camera{
		BottomLeft: lookFrom.Sub(u.MulF(halfWidth * focusDist)).Sub(v.MulF(halfHeight * focusDist)).Sub(w.MulF(focusDist)),
		Horisontal: u.MulF(2 * halfWidth * focusDist),
		Vertical:   v.MulF(2 * halfHeight * focusDist),
		Origin:     lookFrom,
		U:          u,
		V:          v,
		W:          w,
		LensRadius: aperture / 2,
	}
}

// Ray returns ray within camera u and v
func (c Camera) Ray(u, v float64) *Ray {
	var rd = RandomInUnitDisk().MulF(c.LensRadius)
	var offset = c.U.MulF(rd.X()).Add(c.V.MulF(rd.Y()))
	return &Ray{
		Origin:    c.Origin.Add(offset),
		Direction: c.BottomLeft.Add(c.Horisontal.MulF(u)).Add(c.Vertical.MulF(v)).Sub(c.Origin).Sub(offset),
	}
}
