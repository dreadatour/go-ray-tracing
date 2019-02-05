package engine

// Metal material reflects the light (steel, mirror)
type Metal struct {
	Albedo Color
	Fuzz   float64
}

// check if Metal is material
var _ Material = &Metal{}

// Scatter light by metal material
func (m Metal) Scatter(ray *Ray, hit *Hit) (*Color, *Ray, bool) {
	var reflect = Reflect(ray.Direction.Unit(), hit.Normal)
	if Dot(reflect, hit.Normal) <= 0 {
		return nil, nil, false
	}

	var scattered = &Ray{
		Origin:    hit.Point,
		Direction: reflect.Add(RandomInUnitSphere().MulS(m.Fuzz)),
	}
	return &m.Albedo, scattered, true
}
