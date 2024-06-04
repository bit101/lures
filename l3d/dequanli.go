// Package l3d contains 3d strange attractor functions.
package l3d

import "github.com/bit101/lures"

// DequanLi is a DequanLi attractor.
type DequanLi struct {
	lures.Lure
}

// NewDequanLi creates a new DequanLi attractor.
func NewDequanLi() DequanLi {
	a := DequanLi{}
	a.Params.A = 40.0
	a.Params.C = 1.833
	a.Params.D = 0.16
	a.Params.E = 0.65
	a.Params.K = 55.0
	a.Params.F = 20.0
	a.InitX = 0.349
	a.InitY = 0.0
	a.InitZ = -0.16
	a.Scale = 2.0
	a.Dt = 0.0002
	a.Cz = -100
	return a
}

// Iterate iterates a DequanLi attractor one time.
func (a DequanLi) Iterate(x, y, z float64) (float64, float64, float64) {
	dt := a.Dt
	dx := a.Params.A*(y-x) + a.Params.D*x*z
	dy := a.Params.K*x + a.Params.F*y - x*z
	dz := a.Params.C*z + x*y - a.Params.E*x*x
	return x + dx*dt, y + dy*dt, z + dz*dt
}
