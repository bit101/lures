// Package l3d contains 3d strange attractor functions.
package l3d

import "github.com/bit101/lures"

// Halvorsen is a Halvorsen attractor.
type Halvorsen struct {
	lures.Lure
}

// NewHalvorsen creates a new Halvorsen attractor.
func NewHalvorsen() Halvorsen {
	a := Halvorsen{}
	a.Params.A = 1.89
	a.InitX = 0.1
	a.InitY = 0.0
	a.InitZ = 0.0
	a.Scale = 30.0
	a.Dt = 0.01
	a.Cx = 3
	a.Cy = 5
	a.Cz = 5
	return a
}

// Iterate iterates a Halvorsen attractor one time.
func (h Halvorsen) Iterate(x, y, z float64) (float64, float64, float64) {
	dt := h.Dt
	dx := -h.Params.A*x - 4*y - 4*z - y*y
	dy := -h.Params.A*y - 4*z - 4*x - z*z
	dz := -h.Params.A*z - 4*x - 4*y - x*x
	return x + dx*dt, y + dy*dt, z + dz*dt
}
