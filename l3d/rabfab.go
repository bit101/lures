// Package l3d contains 3d strange attractor functions.
package l3d

import "github.com/bit101/lures"

// RabFab is a Rabinovich-Fabrikant attractor.
type RabFab lures.Lure

// NewRabFab creates a new RabFab attractor.
func NewRabFab() RabFab {
	a := RabFab{}
	a.Params.A = 0.14
	a.Params.B = 0.1
	a.InitX = 0.1
	a.InitY = 0.1
	a.InitZ = 0.1
	a.Scale = 200.0
	a.Dt = 0.01
	return a
}

// Iterate iterates a Rabinovich-Fabrikant attractor one time.
func (r RabFab) Iterate(x, y, z, dt float64) (float64, float64, float64) {
	dx := y*(z-1+x*x) + r.Params.B*x
	dy := x*(3*z+1-x*x) + r.Params.B*y
	dz := -2 * z * (r.Params.A + x*y)
	return x + dx*dt, y + dy*dt, z + dz*dt
}

// InitVals returns the suggested initial point values for an attractor.
func (r RabFab) InitVals() (float64, float64, float64) {
	return r.InitX, r.InitY, r.InitZ
}
