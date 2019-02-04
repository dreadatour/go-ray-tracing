package main

import (
	"image/color"
	"math"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"

	"github.com/dreadatour/go-ray-tracing/engine"
)

type render struct {
	// window
	window fyne.Window
	canvas fyne.CanvasObject

	// camera
	bottomLeft engine.Vec3
	horisontal engine.Vec3
	vertical   engine.Vec3
	origin     engine.Vec3

	// objects
	scene engine.Scene
}

func (r *render) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	r.canvas.Resize(size)
}

func (r *render) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(640, 480)
}

func (r *render) refresh() {
	r.window.Canvas().Refresh(r.canvas)
}

func (r *render) draw(px, py, w, h int) color.Color {
	py = h - py

	var (
		u = float64(px) / float64(w)
		v = float64(py) / float64(h)
	)

	var ray = engine.Ray{
		Origin:    r.origin,
		Direction: r.bottomLeft.Add(r.horisontal.MulF(u)).Add(r.vertical.MulF(v)),
	}

	return r.color(&ray).RGBA(1)
}

func (r *render) color(ray *engine.Ray) engine.Color {
	hit, ok := r.scene.Hit(ray, 0, math.MaxFloat64)
	if ok {
		var c = hit.N.Add(engine.Vec3{1, 1, 1}).MulF(0.5)
		return engine.Color{c.X(), c.Y(), c.Z()}
	}

	var dir = ray.Direction.UnitV()
	var tt = 0.5 * (dir.Y() + 1.0)
	var c = engine.Vec3{1, 1, 1}.MulF(1 - tt).Add(engine.Vec3{0.5, 0.7, 1.0}.MulF(tt))

	return engine.Color{c.X(), c.Y(), c.Z()}
}

func (r *render) onKeyDown(ev *fyne.KeyEvent) {
	if ev.Name == fyne.KeyEscape {
		r.window.Close()
	}

	r.refresh()
}

// Render 3D scene
func Render(app fyne.App) {
	window := app.NewWindow("Render")
	window.SetTitle("Render")
	window.SetPadded(false)
	render := &render{
		window:     window,
		bottomLeft: engine.Vec3{-2, -1.5, -1},
		horisontal: engine.Vec3{4, 0, 0},
		vertical:   engine.Vec3{0, 3, 0},
		origin:     engine.Vec3{0, 0, 0},
		scene: engine.Scene{
			engine.Sphere{Center: engine.Vec3{0, 0, -1}, Radius: 0.5},
			engine.Sphere{Center: engine.Vec3{0, -100.5, -1}, Radius: 100},
		},
	}
	render.canvas = canvas.NewRaster(render.draw)

	window.SetContent(fyne.NewContainerWithLayout(render, render.canvas))
	window.Canvas().SetOnKeyDown(render.onKeyDown)
	window.Show()
}
