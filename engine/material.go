package engine

// Material defines now object scatter light
type Material interface {
	Scatter(ray *Ray, hit *Hit) (*Color, *Ray, bool)
}
