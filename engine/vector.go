package engine

import "math"

// Vector is a 3D vector
type Vector [3]float64

// X coordinate
func (v Vector) X() float64 {
	return v[0]
}

// Y coordinate
func (v Vector) Y() float64 {
	return v[1]
}

// Z coordinate
func (v Vector) Z() float64 {
	return v[2]
}

// Add returns vectors sum result (v + v2)
func (v Vector) Add(v2 Vector) Vector {
	return Vector{v[0] + v2[0], v[1] + v2[1], v[2] + v2[2]}
}

// Sub returns vectors substraction result (v - v2)
func (v Vector) Sub(v2 Vector) Vector {
	return Vector{v[0] - v2[0], v[1] - v2[1], v[2] - v2[2]}
}

// Mul returns vectors multiplication result (v * v2)
func (v Vector) Mul(v2 Vector) Vector {
	return Vector{v[0] * v2[0], v[1] * v2[1], v[2] * v2[2]}
}

// Div returns vectors division result (v / v2)
func (v Vector) Div(v2 Vector) Vector {
	return Vector{v[0] / v2[0], v[1] / v2[1], v[2] / v2[2]}
}

// Neg returns negative vector (-v)
func (v Vector) Neg() Vector {
	return Vector{-v[0], -v[1], -v[2]}
}

// MulS returns vector and scalar multiplication result (v * s)
func (v Vector) MulS(s float64) Vector {
	return Vector{v[0] * s, v[1] * s, v[2] * s}
}

// DivS returns vector and scalar division result (v / s)
func (v Vector) DivS(s float64) Vector {
	return Vector{v[0] / s, v[1] / s, v[2] / s}
}

// SquareLength is vector v squared length
func (v Vector) SquareLength() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// Length is vector v length
func (v Vector) Length() float64 {
	return math.Sqrt(v.SquareLength())
}

// Unit is vector whose magnitude is 1
func (v Vector) Unit() Vector {
	return v.DivS(v.Length())
}
