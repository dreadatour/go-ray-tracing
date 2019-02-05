package engine

// Hit is result of ray hitting any object
type Hit struct {
	Distance float64
	Point    Vector
	Normal   Vector
	Material Material
}

// Hitable is interface for hitable by ray objects
type Hitable interface {
	Hit(ray *Ray, dMin float64, dMax float64) (*Hit, bool)
}
