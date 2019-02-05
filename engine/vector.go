package engine

import (
	"math"
	"math/rand"
)

// Vec3 is 3D vector
type Vec3 [3]float64

// X coordinate
func (v Vec3) X() float64 {
	return v[0]
}

// Y coordinate
func (v Vec3) Y() float64 {
	return v[1]
}

// Z coordinate
func (v Vec3) Z() float64 {
	return v[2]
}

// Add is result of vectors v and v2 sum (v + v2)
func (v Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v[0] + v2[0], v[1] + v2[1], v[2] + v2[2]}
}

// Sub is result of vectors v and v2 substraction (v - v2)
func (v Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{v[0] - v2[0], v[1] - v2[1], v[2] - v2[2]}
}

// Mul is result of vectors v and v2 multiplication (v * v2)
func (v Vec3) Mul(v2 Vec3) Vec3 {
	return Vec3{v[0] * v2[0], v[1] * v2[1], v[2] * v2[2]}
}

// Div is result of vectors v and v2 division (v / v2)
func (v Vec3) Div(v2 Vec3) Vec3 {
	return Vec3{v[0] / v2[0], v[1] / v2[1], v[2] / v2[2]}
}

// Neg is nedative vector v (-v)
func (v Vec3) Neg() Vec3 {
	return Vec3{-v[0], -v[1], -v[2]}
}

// MulF is result of vector v and const f multiplication (v / f)
func (v Vec3) MulF(f float64) Vec3 {
	return Vec3{v[0] * f, v[1] * f, v[2] * f}
}

// DivF is result of vector v and const f division (v / f)
func (v Vec3) DivF(f float64) Vec3 {
	return Vec3{v[0] / f, v[1] / f, v[2] / f}
}

// SqLen is vector v squared length
func (v Vec3) SqLen() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// Len is vector v length
func (v Vec3) Len() float64 {
	return math.Sqrt(v.SqLen())
}

// UnitV is vector whose magnitude is 1
func (v Vec3) UnitV() Vec3 {
	return v.DivF(v.Len())
}

// Reflect vector
func (v Vec3) Reflect(n Vec3) Vec3 {
	return v.Sub(n.MulF(Dot(v, n) * 2))
}

// Refract vector
func (v Vec3) Refract(n Vec3, niOverNt float64) (Vec3, bool) {
	var uv = v.UnitV()
	var dt = Dot(uv, n)
	var d = 1 - niOverNt*niOverNt*(1-dt*dt)
	if d <= 0 {
		return Vec3{}, false
	}
	return uv.Sub(n.MulF(dt)).MulF(niOverNt).Sub(n.MulF(math.Sqrt(d))), true
}

// Dot is scalar product of vectors v1 and v2
func Dot(v1, v2 Vec3) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

// Cross is vector whose magnitude is |v1||v2|sin(theta) where theta is the angle between v1 and v2
func Cross(v1, v2 Vec3) Vec3 {
	return Vec3{
		v1[1]*v2[2] - v1[2]*v2[1],
		v1[2]*v2[0] - v1[0]*v2[2],
		v1[0]*v2[1] - v1[1]*v2[0],
	}
}

// RandomInUnitSphere returns random point in unit radius sphere centered at origin
func RandomInUnitSphere() Vec3 {
	var p Vec3
	for {
		p = Vec3{rand.Float64()*2 - 1, rand.Float64()*2 - 1, rand.Float64()*2 - 1}
		if p.SqLen() < 1 {
			return p
		}
	}
}

// RandomInUnitDisk returns random point in unit radius disk centered at origin
func RandomInUnitDisk() Vec3 {
	var p Vec3
	for {
		p = Vec3{rand.Float64()*2 - 1, rand.Float64()*2 - 1, 0}
		if Dot(p, p) < 1 {
			return p
		}
	}
}
