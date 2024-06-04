// Package l3d contains 3d strange attractor functions.
package l3d

import "github.com/bit101/lures"

// ChenLee is a Chen-Lee attractor.
type ChenLee struct {
	lures.Lure
}

// NewChenLee creates a new ChenLee attractor.
func NewChenLee() ChenLee {
	a := ChenLee{}
	a.Params.A = 5.0
	a.Params.B = -10.0
	a.Params.C = -0.38
	a.InitX = 1.0
	a.InitY = 0.0
	a.InitZ = 1.5
	a.Scale = 16.0
	a.Dt = 0.004
	a.Cz = -15
	return a
}

// Iterate iterates a ChenLee attractor one time.
func (c ChenLee) Iterate(x, y, z float64) (float64, float64, float64) {
	dt := c.Dt
	dx := c.Params.A*x - y*z
	dy := c.Params.B*y + x*z
	dz := c.Params.C*z + x*y/3.0
	return x + dx*dt, y + dy*dt, z + dz*dt
}
