package engine

import "math/rand"

// Dielectric material refracts the light (glass, water)
type Dielectric struct {
	RefIdx float64
}

// check if Dielectric is material
var _ Material = &Dielectric{}

// Scatter light by dielectric material
func (m Dielectric) Scatter(ray *Ray, hit *Hit) (*Color, *Ray, bool) {
	var (
		outwardNormal Vector
		niOverNt      float64
		cosine        float64
		reflectProb   float64
		attenuation   = &ColorWhite
		scattered     *Ray
	)

	if Dot(ray.Direction, hit.Normal) > 0 {
		outwardNormal = hit.Normal.Neg()
		niOverNt = m.RefIdx
		cosine = m.RefIdx * Dot(ray.Direction, hit.Normal) / ray.Direction.Length()
	} else {
		outwardNormal = hit.Normal
		niOverNt = 1 / m.RefIdx
		cosine = -Dot(ray.Direction, hit.Normal) / ray.Direction.Length()
	}

	var reflect = Reflect(ray.Direction, hit.Normal)
	var refract, ok = Refract(ray.Direction, outwardNormal, niOverNt)
	if ok {
		reflectProb = schlick(cosine, m.RefIdx)
	} else {
		scattered = &Ray{
			Origin:    hit.Point,
			Direction: reflect,
		}
		reflectProb = 1
	}

	if rand.Float64() < reflectProb {
		scattered = &Ray{
			Origin:    hit.Point,
			Direction: reflect,
		}
	} else {
		scattered = &Ray{
			Origin:    hit.Point,
			Direction: refract,
		}
	}

	return attenuation, scattered, true
}
