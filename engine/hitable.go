package engine

// Hit is result of ray hitting any object
type Hit struct {
	T float64
	P Vec3
	N Vec3
}

// Hitable is interface for hitable by ray objects
type Hitable interface {
	Hit(r *Ray, tMin float64, tMax float64) (*Hit, bool)
}
