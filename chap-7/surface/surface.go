package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width   = 600
	height  = 320
	cells   = 100
	xyrange = 30.0
	xyscale = width / 2 / xyrange
	zscale  = height * .4
)

var sin30, cos30 = 0.5, math.Sqrt(3.0 / 4.0)

// does computation given existing constants
// finds hypotenuse for z & zscale.
func corner(f func(x, y float64) float64, i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

// writes the content of the svg file. Other than boilerplate,
// applies given function to pixels.
func surface(w io.Writer, f func(x, y float64) float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i += 1 {
		for j := 0; j < cells; j += 1 {

			ax, ay := corner(f, i+1, j)
			bx, by := corner(f, i, j)
			cx, cy := corner(f, i, j+1)
			dx, dy := corner(f, i+1, j+1)

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by,
				cx, cy, dx, dy,
			)
		}
	}

	fmt.Fprintln(w, "</svg>")
}

func parseAndCheck(s string) (Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}

	expr, err := Parse(s)

	if err != nil {
		return nil, err
	}

	vars := make(map[Var]bool)

	if err := expr.Check(vars); err != nil {
		return nil, err
	}

	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}

	return expr, nil
}

func plot(w http.ResponseWriter, r *http.Request) {
	// get query parameters
	r.ParseForm()

	// Check that expression is valid
	expr, err := parseAndCheck(r.Form.Get("expr"))

	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
	}

	// set svg header
	w.Header().Set("Content-Type", "image/svg+xml")

	// write to surface, using responseWriter
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y)
		return expr.Eval(Env{"x": x, "y": y, "r": r})
	})
}

func main() {
	http.HandleFunc("/plot", plot)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}