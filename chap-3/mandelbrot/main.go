package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const width = 1024
	const height = 1024

	const xmin, ymin = -2, -2
	const xmax, ymax = 2, 2

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py += 1 {
		y := float64(py)/height*(ymax-ymin) + ymin

		for px := 0; px < width; px += 1 {
			x := float64(px)/height*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	_ = png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128

	for n := uint8(0); n < iterations; n += 1 {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}

	return color.Black
}
