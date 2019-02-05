package engine

// Material defines now object scatter light
type Material interface {
	Scatter(ray *Ray, hit *Hit) (*Color, *Ray, bool)
}

// Lambertian material diffuses the light
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

// Metal material reflects the light
type Metal struct {
	Albedo Color
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
		Direction: reflected,
	}
	return &m.Albedo, scattered, true
}
