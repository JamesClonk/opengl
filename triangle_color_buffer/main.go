package main

import (
	"unsafe"

	. "github.com/JamesClonk/opengl/app"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/glh"
	mgl "github.com/go-gl/mathgl/mgl32"
)

var shader gl.Program
var vertexArray gl.VertexArray

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

type Triangle struct {
	Vertices [3]Vertex
}

type Vertex struct {
	Position mgl.Vec2
	Color    mgl.Vec4
}

func main() {
	app := NewSimpleApp(640, 480, "Go GLFW3 Triangle Color Buffer Example", draw)
	defer app.Destroy()

	createTriangleShader(app.Window)
	app.Start()
}

func createTriangleShader(window *glfw.Window) {
	if int(glh.Sizeof(gl.FLOAT))*2 != int(unsafe.Sizeof(mgl.Vec2{})) {
		panic("wrong float type!")
	} else if int(glh.Sizeof(gl.FLOAT))*6 != int(unsafe.Sizeof(Vertex{})) {
		panic("wrong vertex size!")
	}

	// create triangle
	triangle := Triangle{
		[3]Vertex{
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
		},
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
	gl.BufferData(gl.ARRAY_BUFFER, len(triangle.Vertices)*int(unsafe.Sizeof(Vertex{})), &triangle.Vertices, gl.STATIC_DRAW)

	// enable vertex attributes
	position := shader.GetAttribLocation("position")
	color := shader.GetAttribLocation("color")
	position.EnableArray()
	color.EnableArray()
	position.AttribPointer(2, gl.FLOAT, false, int(unsafe.Sizeof(Vertex{})), nil)
	color.AttribPointer(4, gl.FLOAT, false, int(unsafe.Sizeof(Vertex{})), unsafe.Sizeof(mgl.Vec2{}))

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
