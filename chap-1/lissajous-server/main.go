package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
)

var palette = []color.Color{
	color.White,
	color.Black,
}

const whiteIndex = 0 // first color in palette
const blackIndex = 1 // next color in palette

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	lissajous(w)
}

func lissajous(out io.Writer) {
	const nframes = 64 // animation frames
	const size = 100   // image canvas covers [-size..size]
	const cycles = 5   // number of complete x oscillator revolutions
	const delay = 6    // delay between frames in 10ms units
	const res = 0.001  // angular resolution

	freq := rand.Float64() * 3.0
	phase := 0.1 // phase difference
	anim := gif.GIF{LoopCount: nframes}

	for i := 0; i < nframes; i += 1 {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(
				size+int(x*size+0.5),
				size+int(y*size+0.5),
				blackIndex,
			)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
