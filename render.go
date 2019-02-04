package main

import (
	"image/color"

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
	sphere engine.Sphere
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

	var t = r.sphere.RayHit(ray)
	if t > 0 {
		var n = ray.PointAt(t).Sub(engine.Vec3{0, 0, -1}).UnitV()
		var c = n.Add(engine.Vec3{1, 1, 1}).MulF(0.5)

		return engine.Color{c.X(), c.Y(), c.Z()}.RGBA(1)
	}

	if r.sphere.RayIntersect(ray) {
		return engine.Color{1, 0, 0}.RGBA(1)
	}

	var dir = ray.Direction.UnitV()
	var tt = 0.5 * (dir.Y() + 1.0)
	var c = engine.Vec3{1, 1, 1}.MulF(1 - tt).Add(engine.Vec3{0.5, 0.7, 1.0}.MulF(tt))

	return engine.Color{c.X(), c.Y(), c.Z()}.RGBA(1)
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
		sphere:     engine.Sphere{Center: engine.Vec3{0, 0, -1}, Radius: 0.5},
	}
	render.canvas = canvas.NewRaster(render.draw)

	window.SetContent(fyne.NewContainerWithLayout(render, render.canvas))
	window.Canvas().SetOnKeyDown(render.onKeyDown)
	window.Show()
}
