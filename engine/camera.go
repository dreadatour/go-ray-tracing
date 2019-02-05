package engine

import "math"

// Camera is our window to rendering scene
type Camera struct {
	bottomLeft Vector
	horisontal Vector
	vertical   Vector
	origin     Vector
	u, v, w    Vector
	lensRadius float64
}

// NewCamera returns new Camera with given hfov (in degrees) and aspect ratio
func NewCamera(lookFrom, lookAt, vup Vector, hfov, aspect, aperture, focusDist float64) Camera {
	var theta = hfov * math.Pi / 180
	var halfWidth = math.Tan(theta / 2)
	var halfHeight = halfWidth / aspect
	var w = lookFrom.Sub(lookAt).Unit()
	var u = Cross(vup, w).Unit()
	var v = Cross(w, u)
	return Camera{
		bottomLeft: lookFrom.Sub(u.MulS(halfWidth * focusDist)).Sub(v.MulS(halfHeight * focusDist)).Sub(w.MulS(focusDist)),
		horisontal: u.MulS(2 * halfWidth * focusDist),
		vertical:   v.MulS(2 * halfHeight * focusDist),
		origin:     lookFrom,
		u:          u,
		v:          v,
		w:          w,
		lensRadius: aperture / 2,
	}
}

// Ray returns ray within camera u and v
func (c Camera) Ray(u, v float64) *Ray {
	var rd = RandomInUnitDisk().MulS(c.lensRadius)
	var offset = c.u.MulS(rd.X()).Add(c.v.MulS(rd.Y()))
	return &Ray{
		Origin:    c.origin.Add(offset),
		Direction: c.bottomLeft.Add(c.horisontal.MulS(u)).Add(c.vertical.MulS(v)).Sub(c.origin).Sub(offset),
	}
}
