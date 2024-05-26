// Package l3d contains 3d strange attractor functions.
package l3d

import "github.com/bit101/lures"

// Lorenz is a Lorenz-Fabrikant attractor.
type Lorenz lures.Lure

// NewLorenz creates a new Lorenz attractor.
func NewLorenz() Lorenz {
	a := Lorenz{}
	a.Params.A = 28.0
	a.Params.B = 10.0
	a.Params.C = 8.0 / 3.0
	a.InitX = 0.1
	a.InitY = 0.0
	a.InitZ = 0.0
	a.Scale = 10.0
	a.Dt = 0.01
	return a
}

// Iterate iterates a Lorenz attractor one time.
// Suggest translate z -20
func (l Lorenz) Iterate(x, y, z float64) (float64, float64, float64) {
	dt := l.Dt
	dx := l.Params.B * (y - x)
	dy := x*(l.Params.A-z) - y
	dz := x*y - l.Params.C*z
	return x + dx*dt, y + dy*dt, z + dz*dt
}

// InitVals returns the suggested initial point values for an attractor.
func (l Lorenz) InitVals() (float64, float64, float64) {
	return l.InitX, l.InitY, l.InitZ
}
