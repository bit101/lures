// Package lures contains strange attractor functions.
package lures

// RabFabParams contains the default parameters for the RabFab attractor.
var RabFabParams = LureParams{
	A: 0.14,
	B: 0.1,
}

// RabFab iterates a Rabinovich-Fabrikant attractor one time.
// Suggest dt 0.01, scale 200
func RabFab(x, y, z, dt float64) (float64, float64, float64) {
	dx := y*(z-1+x*x) + RabFabParams.B*x
	dy := x*(3*z+1-x*x) + RabFabParams.B*y
	dz := -2 * z * (RabFabParams.A + x*y)
	return x + dx*dt, y + dy*dt, z + dz*dt
}
