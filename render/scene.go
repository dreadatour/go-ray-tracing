package render

import (
	"math/rand"

	"github.com/dreadatour/go-ray-tracing/engine"
)

func randomScene() engine.Scene {
	var scene = engine.Scene{}

	scene = append(scene, engine.Sphere{
		Center: engine.Vector{0, -1000, 0},
		Radius: 1000,
		Material: engine.Lambertian{
			Albedo: &engine.Color{0.5, 0.5, 0.5},
		},
	})

	for a := -5; a < 5; a++ {
		for b := -5; b < 5; b++ {
			var chooseMat = rand.Float64()
			var center = engine.Vector{float64(a*2) + .9*rand.Float64(), 0.2, float64(b*2) + 0.9*rand.Float64()}
			if center.Sub(engine.Vector{4, 0.2, 0}).Length() > .9 {
				if chooseMat < 0.8 {
					// diffuse
					scene = append(scene, engine.Sphere{
						Center: center,
						Radius: .2,
						Material: engine.Lambertian{
							Albedo: &engine.Color{
								rand.Float64() * rand.Float64(),
								rand.Float64() * rand.Float64(),
								rand.Float64() * rand.Float64(),
							},
						},
					})
				} else if chooseMat < 0.95 {
					// metal
					scene = append(scene, engine.Sphere{
						Center: center,
						Radius: .2,
						Material: engine.Metal{
							Albedo: engine.Color{
								0.5 * (1 + rand.Float64()),
								0.5 * (1 + rand.Float64()),
								0.5 * (1 + rand.Float64()),
							},
							Fuzz: 0.5 * rand.Float64(),
						},
					})
				} else {
					// glass
					scene = append(scene, engine.Sphere{
						Center: center,
						Radius: .2,
						Material: engine.Dielectric{
							RefIdx: 1.5,
						},
					})
				}
			}
		}
	}

	scene = append(scene, engine.Sphere{
		Center: engine.Vector{0, 1, 0},
		Radius: 1,
		Material: engine.Dielectric{
			RefIdx: 1.5,
		},
	})
	scene = append(scene, engine.Sphere{
		Center: engine.Vector{-4, 1, 0},
		Radius: 1,
		Material: engine.Lambertian{
			Albedo: &engine.Color{0.4, 0.2, 0.1},
		},
	})
	scene = append(scene, engine.Sphere{
		Center: engine.Vector{4, 1, 0},
		Radius: 1,
		Material: engine.Metal{
			Albedo: engine.Color{0.7, 0.6, 0.5},
			Fuzz:   0,
		},
	})

	return scene
}
