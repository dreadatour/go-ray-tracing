package engine

// Scene is list of hittable objects
type Scene []Hitable

// check if scene is hittable
var _ Hitable = &Scene{}

// Hit returns closest hit
func (s Scene) Hit(ray *Ray, tMin float64, tMax float64) (*Hit, bool) {
	var (
		bestHit *Hit
		isHit   bool
		closest = tMax
	)

	for _, h := range s {
		if hit, ok := h.Hit(ray, tMin, closest); ok {
			bestHit = hit
			isHit = true
			closest = hit.T
		}
	}

	return bestHit, isHit
}
