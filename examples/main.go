// Package main renders an image or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
	"github.com/bit101/lures/l2d"
	"github.com/bit101/lures/l3d"
	"github.com/bit101/wire"
)

//revive:disable:unused-parameter
const (
	tau = blmath.Tau
	pi  = math.Pi
)

func main() {
	renderTarget := target.Video
	fileName := "aizawa"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)

		// 3d
		// program.AddSceneWithFrames(scene1, 180)

		// 2d
		program.AddSceneWithFrames(scene2, 180)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	attr := l3d.NewAizawa()
	x, y, z := attr.InitVals3d()
	scale := attr.Scale

	shape := wire.NewShape()
	for range 20000 {
		x, y, z = attr.Iterate(x, y, z)
		shape.AddXYZ(x, y, z)
	}
	shape.UniScale(scale)
	shape.Rotate(percent*tau, percent*2*tau, 0)

	shape.RenderPoints(2)

	shape.TranslateZ(-20)
	context.GaussianBlur(20)
	shape.RenderPoints(1)

}

func scene2(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)

	attr := l2d.NewIkeda()
	attr.Params.C = blmath.LoopSin(percent, 0.0, 0.4)
	attr.Params.U = blmath.LoopSin(percent, 0.97, 0.9)
	x, y := attr.InitVals2d()
	scale := attr.Scale

	shape := wire.NewShape()
	for range 10000 {
		x, y = attr.Iterate(x, y)
		shape.AddXYZ(x, y, 0)
	}
	shape.UniScale(scale)
	// shape.Rotate(percent*tau, percent*2*tau, 0)

	shape.RenderPoints(2)

	context.GaussianBlur(20)
	shape.RenderPoints(1)

}
