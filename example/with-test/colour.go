package color

// Reference:
// https://github.com/gerow/go-color/blob/master/color.go
// https://stackoverflow.com/a/9493060/7153181

import (
	"fmt"
	"math"
)

type RGB struct {
	R, G, B float64 // RGB are expected to be within [0, 255].
}

type HSL struct {
	H, S, L float64 // HSL are expected to be within [0, 1].
}

func (c HSL) ToRGB() RGB {
	h := c.H
	s := c.S
	l := c.L

	if s == 0 {
		// achromatic
		return RGB{math.Round(l * 255), math.Round(l * 255), math.Round(l * 255)}
	}

	var v1, v2 float64
	if l < 0.5 {
		v2 = l * (1 + s)
	} else {
		v2 = (l + s) - (s * l)
	}

	v1 = 2*l - v2

	r := hueToRGB(v1, v2, h+(1.0/3.0))
	g := hueToRGB(v1, v2, h)
	b := hueToRGB(v1, v2, h-(1.0/3.0))

	return RGB{math.Round(r * 255), math.Round(g * 255), math.Round(b * 255)}
}
func hueToRGB(v1, v2, h float64) float64 {
	if h < 0 {
		h += 1
	}
	if h > 1 {
		h -= 1
	}
	switch {
	case 6*h < 1:
		return (v1 + (v2-v1)*6*h)
	case 2*h < 1:
		return v2
	case 3*h < 2:
		return v1 + (v2-v1)*((2.0/3.0)-h)*6
	default:
		return v1
	}
}

func (c HSL) ToHex() string {
	return c.ToRGB().ToHex()
}

func (c RGB) ToHSL() HSL {
	var h, s, l float64

	r := c.R / 255
	g := c.G / 255
	b := c.B / 255

	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)

	// Luminosity is the average of the max and min rgb color intensities.
	l = (max + min) / 2

	// saturation
	delta := max - min
	if delta == 0 {
		// achromatic
		return HSL{0, 0, l}
	}

	if l < 0.5 {
		s = delta / (max + min)
	} else {
		s = delta / (2 - max - min)
	}

	// hue
	r2 := (((max - r) / 6) + (delta / 2)) / delta
	g2 := (((max - g) / 6) + (delta / 2)) / delta
	b2 := (((max - b) / 6) + (delta / 2)) / delta
	switch {
	case r == max:
		h = b2 - g2
	case g == max:
		h = (1.0 / 3.0) + r2 - b2
	case b == max:
		h = (2.0 / 3.0) + g2 - r2
	}

	// fix wraparounds
	switch {
	case h < 0:
		h += 1
	case h > 1:
		h -= 1
	}

	return HSL{h, s, l}
}

// A nudge to make truncation round to nearest number instead of flooring
const delta = 1 / 512.0

func (c RGB) ToHex() string {
	return fmt.Sprintf("%02x%02x%02x", byte((c.R + delta)), byte((c.G + delta)), byte((c.B + delta)))
}
