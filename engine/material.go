package engine

// Material defines now object scatter light
type Material interface {
	Scatter(ray *Ray, hit *Hit) (*Color, *Ray, bool)
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
	var reflected = ray.Direction.UnitV().Reflect(hit.N)
	if Dot(reflected, hit.N) <= 0 {
		return nil, nil, false
	}

	var scattered = &Ray{
		Origin:    hit.P,
		Direction: reflected.Add(RandomInUnitSphere().MulF(m.Fuzz)),
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
		attenuation   = &ColorWhite
		scattered     *Ray
	)

	if Dot(ray.Direction, hit.N) > 0 {
		outwardNormal = hit.N.Neg()
		niOverNt = m.RefIdx
	} else {
		outwardNormal = hit.N
		niOverNt = 1 / m.RefIdx
	}

	var refracted, ok = ray.Direction.Refract(outwardNormal, niOverNt)
	if ok {
		scattered = &Ray{
			Origin:    hit.P,
			Direction: refracted,
		}
		return attenuation, scattered, true
	}

	var reflected = ray.Direction.UnitV().Reflect(hit.N)
	scattered = &Ray{
		Origin:    hit.P,
		Direction: reflected,
	}
	return attenuation, scattered, false
}
