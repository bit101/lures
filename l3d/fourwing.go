// Package l3d contains 3d strange attractor functions.
package l3d

import "github.com/bit101/lures"

// FourWings is a FourWings attractor.
type FourWings struct {
	lures.Lure
}

// NewFourWings creates a new FourWings attractor.
func NewFourWings() FourWings {
	a := FourWings{}
	a.Params.A = 0.2
	a.Params.B = 0.01
	a.Params.C = -0.4
	a.InitX = 0.1
	a.InitY = 0.0
	a.InitZ = 0.0
	a.Scale = 125.0
	a.Dt = 0.1
	return a
}

// Iterate iterates a FourWings attractor one time.
func (f FourWings) Iterate(x, y, z float64) (float64, float64, float64) {
	dt := f.Dt
	dx := f.Params.A*x + y*z
	dy := f.Params.B*x + f.Params.C*y - x*z
	dz := -z - x*y
	return x + dx*dt, y + dy*dt, z + dz*dt
}
