// Package l3d contains 3d strange attractor functions.
package l3d

import "github.com/bit101/lures"

// RosslerParams contains the default parameters for the Rossler attractor.
var RosslerParams = lures.LureParams{
	A: 0.2,
	B: 0.2,
	C: 5.7,
}

// Rossler iterates a Rossler attractor one time.
// Suggest dt 0.01, scale 30
func Rossler(x, y, z, dt float64) (float64, float64, float64) {
	dx := -(y + z)
	dy := x + RosslerParams.A*y
	dz := RosslerParams.B + z*(x-RosslerParams.C)
	return x + dx*dt, y + dy*dt, z + dz*dt
}
