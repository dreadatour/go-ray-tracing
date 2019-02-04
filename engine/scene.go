package engine

// Scene is list of hittable objects
type Scene []Hitable

// check if scene is hittable
var _ Hitable = &Scene{}

// Hit returns closest hit
func (s Scene) Hit(r *Ray, tMin float64, tMax float64) (bestHit *Hit, isHit bool) {
	closest := tMax

	for _, h := range s {
		if hit, ok := h.Hit(r, tMin, closest); ok {
			isHit = true
			bestHit = hit
			closest = hit.T
		}
	}

	return
}
