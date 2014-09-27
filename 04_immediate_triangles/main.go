package main

import (
	"math"

	. "github.com/JamesClonk/opengl/_includes"
	"github.com/go-gl/gl"
)

var counter float64 = 0.0

func main() {
	app := NewSimpleApp(640, 480, "Go GLFW3 Immediate Mode Triangles Example", draw)
	defer app.Destroy()

	app.Start()
}

func drawTriangle() {
	gl.Begin(gl.TRIANGLES)
	gl.Color3f(1, 0, 0)
	gl.Vertex3f(-1, 0, 0)
	gl.Color3f(0, 1, 0)
	gl.Vertex3f(0, 1, 0)
	gl.Color3f(0, 0, 1)
	gl.Vertex3f(1, 0, 0)
	gl.End()
}

func draw(app *App) {
	counter += 0.01

	// transform orthogonal projection matrix
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(-float64(app.Ratio), float64(app.Ratio), -1, 2, -1.0, 1.0)

	// transform modelview matrix
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	for i := float64(0); i < 5; {
		i++
		t := i / 10
		gl.PushMatrix()

		// scale, then rotate, then translate
		gl.Translated(t*math.Sin(counter), 0, 0)
		gl.Rotated(360*t*counter, 0, 0, 1)
		gl.Scaled(1-t, 1-t, 1-t)

		drawTriangle()

		gl.PopMatrix()
	}
}
