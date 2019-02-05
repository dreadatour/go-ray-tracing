package engine

// Scene is list of hittable objects
type Scene []Hitable

// check if scene is hittable
var _ Hitable = &Scene{}

// Hit returns closest hit
func (s Scene) Hit(ray *Ray, dMin float64, dMax float64) (*Hit, bool) {
	var (
		hit     *Hit
		isHit   bool
		closest = dMax
	)

	for _, obj := range s {
		if h, ok := obj.Hit(ray, dMin, closest); ok {
			hit = h
			isHit = true
			closest = h.Distance
		}
	}

	return hit, isHit
}
