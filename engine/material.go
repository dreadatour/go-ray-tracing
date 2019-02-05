package engine

import (
	"math"
	"math/rand"
)

// Material defines now object scatter light
type Material interface {
	Scatter(ray *Ray, hit *Hit) (*Color, *Ray, bool)
}

func schlick(cosine, refIdx float64) float64 {
	var r0 = (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}

// Lambertian material diffuses the light (fiber)
type Lambertian struct {
	Albedo *Color
}

// check if Lambertian is material
var _ Material = &Lambertian{}

// Scatter light by lambertian material
func (m Lambertian) Scatter(ray *Ray, hit *Hit) (*Color, *Ray, bool) {
	var target = hit.P.Add(hit.N).Add(RandomInUnitSphere())
	var scattered = &Ray{
		Origin:    hit.P,
		Direction: target.Sub(hit.P),
	}
	return m.Albedo, scattered, true
}

// Metal material reflects the light (steel, mirror)
type Metal struct {
	Albedo Color
	Fuzz   float64
}

// check if Metal is material
var _ Material = &Metal{}

// Scatter light by metal material
func (m Metal) Scatter(ray *Ray, hit *Hit) (*Color, *Ray, bool) {
	var reflect = ray.Direction.UnitV().Reflect(hit.N)
	if Dot(reflect, hit.N) <= 0 {
		return nil, nil, false
	}

	var scattered = &Ray{
		Origin:    hit.P,
		Direction: reflect.Add(RandomInUnitSphere().MulF(m.Fuzz)),
	}
	return &m.Albedo, scattered, true
}

// Dielectric material refracts the light (glass, water)
type Dielectric struct {
	RefIdx float64
}

// check if Dielectric is material
var _ Material = &Dielectric{}

// Scatter light by dielectric material
func (m Dielectric) Scatter(ray *Ray, hit *Hit) (*Color, *Ray, bool) {
	var (
		outwardNormal Vec3
		niOverNt      float64
		cosine        float64
		reflectProb   float64
		attenuation   = &ColorWhite
		scattered     *Ray
	)

	if Dot(ray.Direction, hit.N) > 0 {
		outwardNormal = hit.N.Neg()
		niOverNt = m.RefIdx
		cosine = m.RefIdx * Dot(ray.Direction, hit.N) / ray.Direction.Len()
	} else {
		outwardNormal = hit.N
		niOverNt = 1 / m.RefIdx
		cosine = -Dot(ray.Direction, hit.N) / ray.Direction.Len()
	}

	var reflect = ray.Direction.Reflect(hit.N)
	var refract, ok = ray.Direction.Refract(outwardNormal, niOverNt)
	if ok {
		reflectProb = schlick(cosine, m.RefIdx)
	} else {
		scattered = &Ray{
			Origin:    hit.P,
			Direction: reflect,
		}
		reflectProb = 1
	}

	if rand.Float64() < reflectProb {
		scattered = &Ray{
			Origin:    hit.P,
			Direction: reflect,
		}
	} else {
		scattered = &Ray{
			Origin:    hit.P,
			Direction: refract,
		}
	}

	return attenuation, scattered, true
}
