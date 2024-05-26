// Package l3d contains 3d strange attractor functions.
package l3d

import "github.com/bit101/lures"

// Aizawa is a Aizawa attractor.
type Aizawa lures.Lure

// NewAizawa creates a new Aizawa attractor.
func NewAizawa() Aizawa {
	a := Aizawa{}
	a.Params.A = 0.97
	a.Params.B = 0.7
	a.Params.C = 0.6
	a.Params.D = 3.5
	a.Params.E = 0.25
	a.Params.F = 0.1
	a.InitX = 0.1
	a.InitY = 0.0
	a.InitZ = 0.0
	a.Scale = 200.0
	a.Dt = 0.01
	return a
}

// Iterate iterates a Aizawa attractor one time.
func (a Aizawa) Iterate(x, y, z float64) (float64, float64, float64) {
	dt := a.Dt
	dx := (z-a.Params.B)*x - a.Params.D*y
	dy := a.Params.D*x + (z-a.Params.B)*y
	dz := a.Params.C + a.Params.A*z - z*z*z/3 - (x*x+y*y)*(1+a.Params.E*z) + a.Params.F*z*x*x*x
	return x + dx*dt, y + dy*dt, z + dz*dt
}

// InitVals returns the suggested initial point values for an attractor.
func (a Aizawa) InitVals() (float64, float64, float64) {
	return a.InitX, a.InitY, a.InitZ
}
