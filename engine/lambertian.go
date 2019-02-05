package engine

// Lambertian material diffuses the light (fiber)
type Lambertian struct {
	Albedo *Color
}

// check if Lambertian is material
var _ Material = &Lambertian{}

// Scatter light by lambertian material
func (m Lambertian) Scatter(ray *Ray, hit *Hit) (*Color, *Ray, bool) {
	var target = hit.Point.Add(hit.Normal).Add(RandomInUnitSphere())
	var scattered = &Ray{
		Origin:    hit.Point,
		Direction: target.Sub(hit.Point),
	}
	return m.Albedo, scattered, true
}
