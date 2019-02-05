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
		camera: engine.Camera{
			BottomLeft: engine.Vec3{-2, -1.5, -1},
			Horisontal: engine.Vec3{4, 0, 0},
			Vertical:   engine.Vec3{0, 3, 0},
			Origin:     engine.Vec3{0, 0, 0},
		},
		scene: engine.Scene{
			engine.Sphere{
				Center: engine.Vec3{0, 0, -1},
				Radius: 0.5,
				Material: engine.Lambertian{
					Albedo: &engine.Color{0.8, 0.3, 0.3},
				},
			},
			engine.Sphere{
				Center: engine.Vec3{0, -100.5, -1},
				Radius: 100,
				Material: engine.Lambertian{
					Albedo: &engine.Color{0.8, 0.8, 0},
				},
			},
		},
	}
	render.canvas = canvas.NewRaster(render.draw)

	window.SetContent(fyne.NewContainerWithLayout(render, render.canvas))
	window.Canvas().SetOnKeyDown(render.onKeyDown)
	window.Show()
}
