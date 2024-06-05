// Package l3d contains 3d strange attractor functions.
package l3d

import "github.com/bit101/lures"

// XingYun is a Chen-Lee attractor.
type XingYun struct {
	lures.Lure
}

// NewXingYun creates a new XingYun attractor.
func NewXingYun() XingYun {
	a := XingYun{}
	a.Params.A = 50.0
	a.Params.B = 10.0
	a.Params.C = 13.0
	a.Params.E = 6.0
	a.InitX = 3.0
	a.InitY = 3.0
	a.InitZ = 0.1
	a.Scale = 16.0
	a.Dt = 0.001
	return a
}

// Iterate iterates a XingYun attractor one time.
func (a XingYun) Iterate(x, y, z float64) (float64, float64, float64) {
	dt := a.Dt
	dx := a.Params.A*(y-x) + y*z*z
	dy := a.Params.B*(x+y) - x*z*z
	dz := -a.Params.C*z + a.Params.E*y + x*y*z
	return x + dx*dt, y + dy*dt, z + dz*dt
}
