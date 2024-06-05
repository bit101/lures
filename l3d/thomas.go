// Package l3d contains 3d strange attractor functions.
package l3d

import (
	"math"

	"github.com/bit101/lures"
)

// Thomas is a Thomas attractor.
type Thomas struct {
	lures.Lure
}

// NewThomas creates a new Thomas attractor.
func NewThomas() Thomas {
	a := Thomas{}
	a.Params.A = 0.19
	a.InitX = 0.1
	a.InitY = 0.0
	a.InitZ = 0.0
	a.Scale = 75.0
	a.Dt = 0.05
	return a
}

// Iterate iterates a Thomas attractor one time.
func (t Thomas) Iterate(x, y, z float64) (float64, float64, float64) {
	dt := t.Dt
	dx := -t.Params.A*x + math.Sin(y)
	dy := -t.Params.A*y + math.Sin(z)
	dz := -t.Params.A*z + math.Sin(x)
	return x + dx*dt, y + dy*dt, z + dz*dt
}
