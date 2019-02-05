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
