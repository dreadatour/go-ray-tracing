package engine

import "math"

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

// MulF is result of vector v and const c multiplication (v / c)
func (v Vec3) MulF(c float64) Vec3 {
	return Vec3{v[0] * c, v[1] * c, v[2] * c}
}

// DivF is result of vector v and const c division (v / c)
func (v Vec3) DivF(c float64) Vec3 {
	return Vec3{v[0] / c, v[1] / c, v[2] / c}
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
