package main

import (
	"flag"

	"fyne.io/fyne/app"

	"github.com/dreadatour/go-ray-tracing/render"
)

func main() {
	var (
		width   = flag.Int("width", 320, "window width")
		height  = flag.Int("height", 240, "window height")
		samples = flag.Int("samples", 1, "samples per pixel")
	)
	flag.Parse()

	app := app.New()
	render.Render(app, *width, *height, *samples)
	app.Run()
}
