// Package lures contains strange attractor functions.
package lures

// ChenParams contains the default parameters for the Chen attractor.
var ChenParams = LureParams{
	// A: 36.0,
	// B: 3.0,
	// C: 20.0,
	A: 5.0,
	B: -10.0,
	C: -0.38,
}

// Chen iterates a Chen attractor one time.
// Suggest dt 0.01, translate z -20, scale 10
func Chen(x, y, z, dt float64) (float64, float64, float64) {
	dx := ChenParams.A*x - y*z
	dy := ChenParams.B*y + x*z
	dz := ChenParams.C*y + x*y/3.0
	return x + dx*dt, y + dy*dt, z + dz*dt
}
