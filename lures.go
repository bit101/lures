// Package lures contains strange attractor functions.
package lures

// Lure is a struct defining a single attractor type.
type Lure struct {
	Params              Params
	InitX, InitY, InitZ float64
	Dt, Scale           float64
	Cx, Cy, Cz          float64
}

// Params defines the possible parameters for an attractor function.
type Params struct {
	A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z float64
}

// InitVals2d returns the suggested initial point values for a 2d attractor.
func (l Lure) InitVals2d() (float64, float64) {
	return l.InitX, l.InitY
}

// InitVals3d returns the suggested initial point values for a 3d attractor.
func (l Lure) InitVals3d() (float64, float64, float64) {
	return l.InitX, l.InitY, l.InitZ
}

// Center2d returns suggested center point values for a 2d attractor.
func (l Lure) Center2d() (float64, float64) {
	return l.Cx, l.Cy
}

// Center3d returns suggested center point values for a 3d attractor.
func (l Lure) Center3d() (float64, float64, float64) {
	return l.Cx, l.Cy, l.Cz
}
