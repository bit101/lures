// Package l3d contains 3d strange attractor functions.
package l3d

import (
	"math"

	"github.com/bit101/lures"
)

// Pickover is a Chen-Lee attractor.
type Pickover struct {
	lures.Lure
}

// NewPickover creates a new Pickover attractor.
func NewPickover() Pickover {
	a := Pickover{}
	a.Params.A = -0.759
	a.Params.B = 2.449
	a.Params.C = 1.253
	a.Params.D = 1.5
	a.InitX = 0.1
	a.InitY = 0.1
	a.InitZ = 0.1
	a.Scale = 600.0
	a.Cx = -0.3
	a.Cy = 0.5
	a.Cz = -0.25
	return a
}

// Iterate iterates a Pickover attractor one time.
func (p Pickover) Iterate(x, y, z float64) (float64, float64, float64) {
	x1 := math.Sin(p.Params.A*y) - z*math.Cos(p.Params.B*x)
	y1 := z*math.Sin(p.Params.C*x) - math.Cos(p.Params.D*y)
	z1 := math.Sin(x)
	return x1, y1, z1
}
