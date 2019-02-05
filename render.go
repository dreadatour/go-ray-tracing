package main

import (
	"image/color"
	"math"
	"math/rand"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"

	"github.com/dreadatour/go-ray-tracing/engine"
)

const ns = 100

type render struct {
	window fyne.Window
	canvas fyne.CanvasObject
	camera engine.Camera
	scene  engine.Scene
}

func (r *render) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	r.canvas.Resize(size)
}

func (r *render) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(640, 480)
}

func (r *render) refresh() {
	// r.window.Canvas().Refresh(r.canvas)
}

func (r *render) draw(px, py, w, h int) color.Color {
	var color = engine.Color{}
	for i := 0; i < ns; i++ {
		u := (float64(px) + rand.Float64() - .5) / float64(w)
		v := (float64(h-py) + rand.Float64() - .5) / float64(h)
		ray := r.camera.Ray(u, v)
		color = color.Add(r.color(ray, 0))
	}
	color = color.DivF(float64(ns))

	return engine.Color{math.Sqrt(color.R()), math.Sqrt(color.G()), math.Sqrt(color.B())}.RGBA(1)
}

func (r *render) color(ray *engine.Ray, depth int) engine.Color {
	hit, ok := r.scene.Hit(ray, 0.001, math.MaxFloat64)
	if ok {
		if depth > 50 {
			return engine.ColorBlack
		}

		var attenuation, scatter, ok2 = hit.Material.Scatter(ray, hit)
		if !ok2 {
			return engine.ColorBlack
		}

		return attenuation.Mul(r.color(scatter, depth+1))
	}

	var dir = ray.Direction.UnitV()
	var tt = 0.5 * (dir.Y() + 1.0)
	var c = engine.Vec3{1, 1, 1}.MulF(1 - tt).Add(engine.Vec3{0.5, 0.7, 1.0}.MulF(tt))

	return engine.Color{c.X(), c.Y(), c.Z()}
}

func (r *render) onKeyDown(ev *fyne.KeyEvent) {
	switch ev.Name {
	case fyne.KeyEscape:
		r.window.Close()
	}
	// r.refresh()
}

// Render 3D scene
func Render(app fyne.App) {
	window := app.NewWindow("Render")
	window.SetTitle("Render")
	window.SetPadded(false)

	render := &render{
		window: window,
		camera: engine.NewCamera(
			engine.Vec3{13, 2, 3},
			engine.Vec3{0, 0, 0},
			engine.Vec3{0, 1, 0},
			30,
			float64(640)/float64(480),
			0.1,
			10,
		),
		scene: randomScene(),
	}
	render.canvas = canvas.NewRaster(render.draw)

	window.SetContent(fyne.NewContainerWithLayout(render, render.canvas))
	window.Canvas().SetOnKeyDown(render.onKeyDown)
	window.Show()
}

func randomScene() engine.Scene {
	var scene = engine.Scene{}

	scene = append(scene, engine.Sphere{
		Center: engine.Vec3{0, -1000, 0},
		Radius: 1000,
		Material: engine.Lambertian{
			Albedo: &engine.Color{0.5, 0.5, 0.5},
		},
	})

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			var chooseMat = rand.Float64()
			var center = engine.Vec3{float64(a) + .9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}
			if center.Sub(engine.Vec3{4, 0.2, 0}).Len() > .9 {
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
		Center: engine.Vec3{0, 1, 0},
		Radius: 1,
		Material: engine.Dielectric{
			RefIdx: 1.5,
		},
	})
	scene = append(scene, engine.Sphere{
		Center: engine.Vec3{-4, 1, 0},
		Radius: 1,
		Material: engine.Lambertian{
			Albedo: &engine.Color{0.4, 0.2, 0.1},
		},
	})
	scene = append(scene, engine.Sphere{
		Center: engine.Vec3{4, 1, 0},
		Radius: 1,
		Material: engine.Metal{
			Albedo: engine.Color{0.7, 0.6, 0.5},
			Fuzz:   0,
		},
	})

	return scene
}
