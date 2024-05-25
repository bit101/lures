// Package lures contains strange attractor functions.
package lures

// SprottParams contains the default parameters for the Sprott attractor.
var SprottParams = LureParams{
	A: 2.07,
	B: 1.79,
}

// Sprott iterates a Sprott attractor one time.
// Initial point matters. Suggest 1.9, 0.1, 0.1
// Suggest dt 0.01, scale 200-300x
func Sprott(x, y, z, dt float64) (float64, float64, float64) {
	dx := y + SprottParams.A*x*y + x*z
	dy := 1 - SprottParams.B*x*x + y*z
	dz := x - x*x - y*y
	return x + dx*dt, y + dy*dt, z + dz*dt
}
