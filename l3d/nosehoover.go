// Package l3d contains 3d strange attractor functions.
package l3d

import "github.com/bit101/lures"

// NoseHoover is a NoseHoover attractor.
type NoseHoover struct {
	lures.Lure
}

// NewNoseHoover creates a new Nose-Hoover attractor.
func NewNoseHoover() NoseHoover {
	a := NoseHoover{}
	a.Params.A = 1.5
	a.InitX = 0.1
	a.InitY = 0.0
	a.InitZ = 0.0
	a.Scale = 75.0
	a.Dt = 0.01
	return a
}

// Iterate iterates a NoseHoover attractor one time.
func (n NoseHoover) Iterate(x, y, z float64) (float64, float64, float64) {
	dt := n.Dt
	dx := y
	dy := -x + y*z
	dz := n.Params.A - y*y
	return x + dx*dt, y + dy*dt, z + dz*dt
}
