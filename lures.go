// Package lures contains strange attractor functions.
package lures

// Lure is a struct defining a single attractor type.
type Lure struct {
	Params                         Params
	InitX, InitY, InitZ, Dt, Scale float64
}

// Iterator is an interface for an attractor iterator function.
type Iterator func(x, y, z float64) (float64, float64, float64)

// Params defines the possible parameters for an attractor function.
type Params struct {
	A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z float64
}
