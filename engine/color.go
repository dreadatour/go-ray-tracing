package engine

import "image/color"

// Color is 3D vector
type Color [3]float64

var (
	// ColorBlack is black color (#000000)
	ColorBlack = Color{0, 0, 0}
	// ColorWhite is white color (#ffffff)
	ColorWhite = Color{1, 1, 1}
)

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

// Add returns colors sum (c + c2)
func (c Color) Add(c2 Color) Color {
	return Color{c[0] + c2[0], c[1] + c2[1], c[2] + c2[2]}
}

// Mul returns colors multiplication (c * c2)
func (c Color) Mul(c2 Color) Color {
	return Color{c[0] * c2[0], c[1] * c2[1], c[2] * c2[2]}
}

// MulS returns color and scalar multiplication (c * s)
func (c Color) MulS(s float64) Color {
	return Color{c[0] * s, c[1] * s, c[2] * s}
}

// DivS returns color and scalar division (c / s)
func (c Color) DivS(s float64) Color {
	return Color{c[0] / s, c[1] / s, c[2] / s}
}

// RGBA color from color and alpha-channel value
func (c Color) RGBA(a float64) color.RGBA {
	return color.RGBA{
		uint8(255.999 * c[0]),
		uint8(255.999 * c[1]),
		uint8(255.999 * c[2]),
		uint8(255.999 * a),
	}
}
