// Package l2d contains 2d strange attractor functions.
package l2d

import (
	"math"

	"github.com/bit101/lures"
)

// Pickover is a Aizawa attractor.
type Pickover struct {
	lures.Lure
}

// NewPickover creates a new Pickover attractor.
func NewPickover() Pickover {
	a := Pickover{}
	a.Params.A = -1.7
	a.Params.B = -1.8
	a.Params.C = 1.9
	a.Params.D = -0.4
	a.InitX = 0.1
	a.InitY = 0.1
	a.Scale = 200.0
	return a
}

// Iterate iterates a Pickover attractor one time.
func (p Pickover) Iterate(x, y float64) (float64, float64) {
	x1 := math.Sin(p.Params.A*y) + p.Params.C*math.Cos(p.Params.A*x)
	y1 := math.Sin(p.Params.B*x) + p.Params.D*math.Cos(p.Params.B*y)
	return x1, y1
}
