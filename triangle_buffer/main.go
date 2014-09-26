package main

import (
	. "github.com/JamesClonk/opengl/_includes"
	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	glh "github.com/go-gl/glh"
)

var shader gl.Program
var vertexArray gl.VertexArray

// const vertexShaderSource = `
// 	#version 330
// 		in vec2 position;
// 		void main() {
// 			gl_Position = vec4(position, 0, 1);
// 		}
// `

// const fragmentShaderSource = `
// 	#version 330
// 		out vec4 colourOut;
// 		void main() {
// 			colourOut = vec4(0, 1.0, 0, 1.0);
// 		}
// `

const vertexShaderSource = `
	#version 120
		attribute vec2 position;
		void main() {
			gl_Position = vec4(position, 0, 1);
		}
`

const fragmentShaderSource = `
	#version 120
		void main() {
			gl_FragColor = vec4(0.0, 1.0, 1.0, 1.0);
		}
`

type Triangle struct {
	Vertices [6]float32
}

func main() {
	app := NewSimpleApp(640, 480, "Go GLFW3 Triangle Buffer Example", draw)
	defer app.Destroy()

	createTriangleShader(app.Window)
	app.Start()
}

func createTriangleShader(window *glfw.Window) {
	// create triangle
	triangle := Triangle{
		[6]float32{-0.5, -0.5, -0.5, 0.5, 0.5, -0.5},
	}

	// create shader
	vertexShader := glh.Shader{gl.VERTEX_SHADER, vertexShaderSource}
	fragmentShader := glh.Shader{gl.FRAGMENT_SHADER, fragmentShaderSource}
	shader = glh.NewProgram(vertexShader, fragmentShader)
	shader.Use()

	// create vertex array object
	vertexArray = gl.GenVertexArray()
	vertexArray.Bind()

	// create vertex buffer object
	triangleBuffer := gl.GenBuffer()
	triangleBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, int(glh.Sizeof(gl.FLOAT))*len(triangle.Vertices), &triangle.Vertices, gl.STATIC_DRAW)

	// enable vertex attributes
	positionLocation := shader.GetAttribLocation("position")
	positionLocation.EnableArray()
	positionLocation.AttribPointer(2, gl.FLOAT, false, int(glh.Sizeof(gl.FLOAT))*2, nil)

	vertexArray.Unbind()
	shader.Unuse()
}

func draw(window *glfw.Window) {
	shader.Use()
	vertexArray.Bind()

	gl.DrawArrays(gl.TRIANGLES, 0, 3)

	vertexArray.Unbind()
	shader.Unuse()
}
