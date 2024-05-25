// Package main renders an image or video
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	cairo "github.com/bit101/blcairo"
	"github.com/bit101/blcairo/render"
	"github.com/bit101/blcairo/target"
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
	fileName := "rossler"

	if renderTarget == target.Image {
		render.CreateAndViewImage(800, 800, "out/"+fileName+".png", scene1, 0.0)
	} else if renderTarget == target.Video {
		program := render.NewProgram(400, 400, 30)
		program.AddSceneWithFrames(scene1, 360)
		program.RenderAndPlayVideo("out/frames", "out/"+fileName+".mp4")
	}
}

func scene1(context *cairo.Context, width, height, percent float64) {
	context.WhiteOnBlack()
	wire.InitWorld(context, 200, 200, 800)
	x := 1.1
	y := 1.1
	z := -0.1
	// x := random.FloatRange(-1, 1)
	// y := random.FloatRange(-1, 1)
	// z := random.FloatRange(-1, 1)

	shape := wire.NewShape()
	for range 10000 {
		x, y, z = l3d.FourWing(x, y, z, 0.1)
		// fmt.Println(x, y, z)
		shape.AddXYZ(x, y, z)
	}
	// shape.TranslateZ(-25)
	shape.UniScale(200)
	shape.Rotate(percent*tau, percent*2*tau, 0)

	shape.RenderPoints(2)

	// shape.TranslateZ(-20)
	// context.GaussianBlur(20)
	// shape.RenderPoints(1)

}
