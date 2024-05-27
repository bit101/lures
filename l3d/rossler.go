// Package l3d contains 3d strange attractor functions.
package l3d

import "github.com/bit101/lures"

// Rossler is a Rossler attractor.
type Rossler struct {
	lures.Lure
}

// NewRossler creates a new Rossler attractor.
func NewRossler() Rossler {
	a := Rossler{}
	a.Params.A = 0.2
	a.Params.B = 0.2
	a.Params.C = 5.7
	a.InitX = 0.01
	a.InitY = 0.0
	a.InitZ = 0.0
	a.Scale = 30.0
	a.Dt = 0.01
	return a
}

// Iterate iterates a Rossler attractor one time.
func (r Rossler) Iterate(x, y, z float64) (float64, float64, float64) {
	dt := r.Dt
	dx := -(y + z)
	dy := x + r.Params.A*y
	dz := r.Params.B + z*(x-r.Params.C)
	return x + dx*dt, y + dy*dt, z + dz*dt
}
