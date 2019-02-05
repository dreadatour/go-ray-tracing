package render

import (
	"image/color"
	"math"
	"math/rand"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"

	"github.com/dreadatour/go-ray-tracing/engine"
)

type render struct {
	width   int
	height  int
	samples int
	window  fyne.Window
	canvas  fyne.CanvasObject
	camera  engine.Camera
	scene   engine.Scene
}

func (r *render) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	r.canvas.Resize(size)
}

func (r *render) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(r.width, r.height)
}

// func (r *render) refresh() {
// 	r.window.Canvas().Refresh(r.canvas)
// }

func (r *render) draw(px, py, w, h int) color.Color {
	var color = engine.Color{}
	for i := 0; i < r.samples; i++ {
		u := (float64(px) + rand.Float64() - .5) / float64(w)
		v := (float64(h-py) + rand.Float64() - .5) / float64(h)
		ray := r.camera.Ray(u, v)
		color = color.Add(r.color(ray, 0))
	}
	color = color.DivS(float64(r.samples))

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

	var dir = ray.Direction.Unit()
	var tt = 0.5 * (dir.Y() + 1.0)
	var c = engine.Vector{1, 1, 1}.MulS(1 - tt).Add(engine.Vector{0.5, 0.7, 1.0}.MulS(tt))

	return engine.Color{c.X(), c.Y(), c.Z()}
}

func (r *render) onKeyDown(ev *fyne.KeyEvent) {
	switch ev.Name {
	case fyne.KeyEscape:
		r.window.Close()
	}
}

// Render 3D scene
func Render(app fyne.App, width, height, samples int) {
	window := app.NewWindow("Render")
	window.SetTitle("Render")
	window.SetPadded(false)

	render := &render{
		width:   width,
		height:  height,
		samples: samples,
		window:  window,
		camera: engine.NewCamera(
			engine.Vector{13, 2, 3},
			engine.Vector{0, 0, 0},
			engine.Vector{0, 1, 0},
			30,
			float64(width)/float64(height),
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
