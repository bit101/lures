// Package l2d contains 2d strange attractor functions.
package l2d

import (
	"math"

	"github.com/bit101/lures"
)

// Ikeda is a Ikeda attractor.
type Ikeda struct {
	lures.Lure
}

// NewIkeda creates a new Ikeda attractor.
func NewIkeda() Ikeda {
	a := Ikeda{}
	a.Params.C = 0.4
	a.Params.U = 0.9
	a.InitX = 0.1
	a.InitY = 0.1
	a.Scale = 200.0
	return a
}

// Iterate iterates a Ikeda attractor one time.
func (i Ikeda) Iterate(x, y float64) (float64, float64) {
	t := i.Params.C - 6.0/(1+x*x+y*y)
	x1 := 1 + i.Params.U*(x*math.Cos(t)-y*math.Sin(t))
	y1 := i.Params.U * (x*math.Sin(t) + y*math.Cos(t))
	return x1, y1
}
