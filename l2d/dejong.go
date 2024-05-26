// Package l2d contains 2d strange attractor functions.
package l2d

import (
	"math"

	"github.com/bit101/lures"
)

// Dejong is a Aizawa attractor.
type Dejong lures.Lure

// NewDejong creates a new Dejong attractor.
func NewDejong() Dejong {
	a := Dejong{}
	a.Params.A = 4.0
	a.Params.B = 60.0
	a.Params.C = 10
	a.Params.D = 1.6
	a.InitX = 0.1
	a.InitY = 0.1
	a.Scale = 200.0
	return a
}

// Iterate iterates a Dejong attractor one time.
func (d Dejong) Iterate(x, y float64) (float64, float64) {
	x1 := math.Sin(d.Params.A*y) - math.Cos(d.Params.B*x)
	y1 := math.Sin(d.Params.C*x) - math.Cos(d.Params.D*y)
	return x1, y1
}

// InitVals returns the suggested initial point values for an attractor.
func (d Dejong) InitVals() (float64, float64) {
	return d.InitX, d.InitY
}
