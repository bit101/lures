// Package l3d contains 3d strange attractor functions.
package l3d

import "github.com/bit101/lures"

// LorenzParams contains the default parameters for the Lorenz attractor.
var LorenzParams = lures.LureParams{
	A: 28.0,
	B: 10.0,
	C: 8.0 / 3.0,
}

// Lorenz iterates a Lorenz attractor one time.
// Suggest dt 0.01, translate z -20, scale 10
func Lorenz(x, y, z, dt float64) (float64, float64, float64) {
	dx := LorenzParams.B * (y - x)
	dy := x*(LorenzParams.A-z) - y
	dz := x*y - LorenzParams.C*z
	return x + dx*dt, y + dy*dt, z + dz*dt
}
