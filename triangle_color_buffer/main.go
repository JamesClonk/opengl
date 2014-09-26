package main

import (
	. "github.com/JamesClonk/opengl/_includes"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	mgl "github.com/go-gl/mathgl/mgl32"
)

var shader *Shader

const vertexShaderSource = `
	#version 130
		in vec2 position;
		in vec4 color;

		varying vec4 vertexColor;

		void main()	{
			vertexColor = color;
			gl_Position = vec4(position, 0, 1);
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
			mgl.Vec2{-0.5, -0.5},
			mgl.Vec4{1.0, 0.0, 0.0, 1.0},
		},
		Vertex{
			mgl.Vec2{-0.5, 0.5},
			mgl.Vec4{0.0, 1.0, 0.0, 1.0},
		},
		Vertex{
			mgl.Vec2{0.5, -0.5},
			mgl.Vec4{0.0, 0.0, 1.0, 1.0},
		},
	}
	shader = NewSimpleShader(app.Window, &triangle, vertexShaderSource, fragmentShaderSource)

	app.Start()
}

func draw(window *glfw.Window) {
	shader.Use()

	gl.DrawArrays(gl.TRIANGLES, 0, 3)

	shader.Unuse()
}
