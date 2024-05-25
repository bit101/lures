// Package lures contains strange attractor functions.
package lures

// HalvorsenParams contains the default parameters for the Halvorsen attractor.
var HalvorsenParams = LureParams{
	A: 1.89,
}

// Halvorsen iterates a Halvorsen attractor one time.
// Suggest dt 0.01, scale 30
func Halvorsen(x, y, z, dt float64) (float64, float64, float64) {
	dx := -HalvorsenParams.A*x - 4*y - 4*z - y*y
	dy := -HalvorsenParams.A*y - 4*z - 4*x - z*z
	dz := -HalvorsenParams.A*z - 4*x - 4*y - x*x
	return x + dx*dt, y + dy*dt, z + dz*dt
}
