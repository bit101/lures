// Package l2d contains 2d strange attractor functions.
package l2d

import (
	"math"

	"github.com/bit101/lures"
)

// SineSine is a SineSine attractor.
type SineSine struct {
	lures.Lure
}

// NewSineSine creates a new SineSine attractor.
func NewSineSine() SineSine {
	a := SineSine{}
	a.Params.A = 2.1
	a.InitX = 0.1
	a.InitY = 0.1
	a.Scale = 200.0
	return a
}

// Iterate iterates a SineSine attractor one time.
func (s SineSine) Iterate(x, y float64) (float64, float64) {
	x1 := math.Sin(x) - math.Sin(s.Params.A*y)
	y1 := x
	return x1, y1
}
