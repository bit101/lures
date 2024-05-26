// Package l3d contains 3d strange attractor functions.
package l3d

import "github.com/bit101/lures"

// Sprott is a Sprott attractor.
type Sprott lures.Lure

// NewSprott creates a new Sprott attractor.
func NewSprott() Sprott {
	a := Sprott{}
	a.Params.A = 2.07
	a.Params.B = 1.79
	a.InitX = 1.9
	a.InitY = 0.1
	a.InitZ = 0.1
	a.Scale = 200.0
	a.Dt = 0.01
	return a
}

// Iterate iterates a Sprott attractor one time.
func (s Sprott) Iterate(x, y, z, dt float64) (float64, float64, float64) {
	dx := y + s.Params.A*x*y + x*z
	dy := 1 - s.Params.B*x*x + y*z
	dz := x - x*x - y*y
	return x + dx*dt, y + dy*dt, z + dz*dt
}

// InitVals returns the suggested initial point values for an attractor.
func (s Sprott) InitVals() (float64, float64, float64) {
	return s.InitX, s.InitY, s.InitZ
}
