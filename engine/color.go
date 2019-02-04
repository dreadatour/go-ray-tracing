package engine

import "image/color"

// Color is 3D vector
type Color [3]float64

// R returns color's red component
func (c Color) R() float64 {
	return c[0]
}

// G returns color's green component
func (c Color) G() float64 {
	return c[1]
}

// B returns color's blue component
func (c Color) B() float64 {
	return c[2]
}

// RGBA color
func (c Color) RGBA(a float64) color.RGBA {
	return color.RGBA{
		uint8(255.999 * c[0]),
		uint8(255.999 * c[1]),
		uint8(255.999 * c[2]),
		uint8(255.999 * a),
	}
}
