// Package lures contains strange attractor functions.
package lures

// ThreeScrollParams contains the default parameters for the ThreeScroll attractor.
var ThreeScrollParams = LureParams{
	A: 32.48,
	B: 45.84,
	C: 1.18,
	D: 0.13,
	E: 0.57,
	F: 14.7,
}

// ThreeScroll iterates a ThreeScroll attractor one time.
// Suggest dt 0.1, scale 200x
func ThreeScroll(x, y, z, dt float64) (float64, float64, float64) {
	dx := ThreeScrollParams.A*(y-x) + ThreeScrollParams.D*x*z
	dy := ThreeScrollParams.B*x - x*z + ThreeScrollParams.F*y
	dz := ThreeScrollParams.C*z + x*y - ThreeScrollParams.E*x*x
	return x + dx*dt, y + dy*dt, z + dz*dt
}
