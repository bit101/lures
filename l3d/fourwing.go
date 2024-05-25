// Package l3d contains 3d strange attractor functions.
package l3d

import "github.com/bit101/lures"

// FourWingParams contains the default parameters for the FourWing attractor.
var FourWingParams = lures.LureParams{
	A: 0.2,
	B: 0.01,
	C: -0.4,
}

// FourWing iterates a FourWing attractor one time.
// Suggest dt 0.1, scale 200x
func FourWing(x, y, z, dt float64) (float64, float64, float64) {
	dx := FourWingParams.A*x + y*z
	dy := FourWingParams.B*x + FourWingParams.C*y - x*z
	dz := -z - x*y
	return x + dx*dt, y + dy*dt, z + dz*dt
}
