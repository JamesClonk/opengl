package main

import (
	. "github.com/JamesClonk/opengl/app"
	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
)

func main() {
	app := NewSimpleApp(640, 480, "Go GLFW3 Immediate Mode Example", draw)
	defer app.Destroy()

	app.Start()
}

func draw(window *glfw.Window) {
	gl.Begin(gl.TRIANGLES)
	gl.Color3f(1, 0, 0)
	gl.Vertex3f(-1, 0, 0)
	gl.Color3f(0, 1, 0)
	gl.Vertex3f(0, 1, 0)
	gl.Color3f(0, 0, 1)
	gl.Vertex3f(1, 0, 0)
	gl.End()
}
