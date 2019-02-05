package engine

import (
	"math"
	"math/rand"
)

// Dot is scalar product of vectors v1 and v2
func Dot(v1, v2 Vector) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

// Cross is vector whose magnitude is |v1||v2|sin(theta) where theta is the angle between v1 and v2
func Cross(v1, v2 Vector) Vector {
	return Vector{
		v1[1]*v2[2] - v1[2]*v2[1],
		v1[2]*v2[0] - v1[0]*v2[2],
		v1[0]*v2[1] - v1[1]*v2[0],
	}
}

// Reflect vector v with normal n
func Reflect(v, n Vector) Vector {
	return v.Sub(n.MulS(Dot(v, n) * 2))
}

// Refract vector v with normal n using refractive indices
func Refract(v, n Vector, niOverNt float64) (Vector, bool) {
	var uv = v.Unit()
	var dt = Dot(uv, n)
	var d = 1 - niOverNt*niOverNt*(1-dt*dt)
	if d <= 0 {
		return Vector{}, false
	}
	return uv.Sub(n.MulS(dt)).MulS(niOverNt).Sub(n.MulS(math.Sqrt(d))), true
}

func schlick(cosine, refIdx float64) float64 {
	var r0 = (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}

// RandomInUnitSphere returns random point in unit radius sphere centered at origin
func RandomInUnitSphere() Vector {
	var p Vector
	for {
		p = Vector{rand.Float64()*2 - 1, rand.Float64()*2 - 1, rand.Float64()*2 - 1}
		if p.SquareLength() < 1 {
			return p
		}
	}
}

// RandomInUnitDisk returns random point in unit radius disk centered at origin
func RandomInUnitDisk() Vector {
	var p Vector
	for {
		p = Vector{rand.Float64()*2 - 1, rand.Float64()*2 - 1, 0}
		if Dot(p, p) < 1 {
			return p
		}
	}
}
