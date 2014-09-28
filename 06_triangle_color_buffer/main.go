package main

import (
	. "github.com/JamesClonk/opengl/_includes"
	"github.com/go-gl/gl"
	mgl "github.com/go-gl/mathgl/mgl32"
)

var shader *Shader

const vertexShaderSource = `
	#version 130
		in vec4 position;
		in vec4 color;

		varying vec4 vertexColor;

		void main()	{
			vertexColor = color;
			gl_Position = position;
		}
`

const fragmentShaderSource = `
	#version 130
		varying vec4 vertexColor;

		void main() {
			gl_FragColor = vertexColor;
		}
`

func main() {
	app := NewSimpleApp(640, 480, "Go GLFW3 Triangle Color Buffer Example", draw)
	defer app.Destroy()

	triangle := Vertices{
		Vertex{
			Position: mgl.Vec4{-0.5, -0.5, 0, 1},
			Color:    mgl.Vec4{1, 0, 0, 1},
		},
		Vertex{
			Position: mgl.Vec4{-0.5, 0.5, 0, 1},
			Color:    mgl.Vec4{0, 1, 0, 1},
		},
		Vertex{
			Position: mgl.Vec4{0.5, -0.5, 0, 1},
			Color:    mgl.Vec4{0, 0, 1, 1},
		},
	}
	shader = NewSimpleShader(&triangle, vertexShaderSource, fragmentShaderSource)

	app.Start()
}

func draw(app *App) {
	shader.Use()

	gl.DrawArrays(gl.TRIANGLES, 0, 3)

	shader.Unuse()
}
