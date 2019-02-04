package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
)

type render struct {
	window fyne.Window
	canvas fyne.CanvasObject
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
	return color.RGBA{
		uint8((float32(px)*255 + .5) / float32(w)),
		uint8((float32(py)*255 + .5) / float32(h)),
		0,
		0xff,
	}
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
		window: window,
	}
	render.canvas = canvas.NewRaster(render.draw)

	window.SetContent(fyne.NewContainerWithLayout(render, render.canvas))
	window.Canvas().SetOnKeyDown(render.onKeyDown)
	window.Show()
}
