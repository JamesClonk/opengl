package main

import (
	. "github.com/JamesClonk/opengl/app"
	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	glh "github.com/go-gl/glh"
	mathgl "github.com/go-gl/mathgl/mgl64"
)

var shader gl.Program
var vertexShader, fragmentShader glh.Shader

const vertexShaderSource = `
	#version 120
		attribute vec4 position;
		void main(){
			gl_Position = position;
		}
`

const fragmentShaderSource = `
	#version 120
		void main(){
			gl_FragColor = vec4(1.0,1.0,1.0,1.0);
		}
`

type Triangle struct {
	Vertices []mathgl.Vec2
}

func main() {
	app := NewSimpleApp(640, 480, "Go GLFW3 Triangle Buffer Example", draw)
	defer app.Destroy()

	initOpenGL(app.Window)
	app.Start()
}

func initOpenGL(window *glfw.Window) {
	triangle := Triangle{}

	triangle.Vertices = append(triangle.Vertices, mathgl.Vec2{-1, -5})
	triangle.Vertices = append(triangle.Vertices, mathgl.Vec2{0, 1})
	triangle.Vertices = append(triangle.Vertices, mathgl.Vec2{1, -5})

	vertexShader = glh.Shader{gl.VERTEX_SHADER, vertexShaderSource}
	fragmentShader = glh.Shader{gl.FRAGMENT_SHADER, fragmentShaderSource}
	shader = glh.NewProgram(vertexShader, fragmentShader)
	shader.Use()
}

func draw(window *glfw.Window) {
	// gl.Begin(gl.TRIANGLES)
	// gl.Color3f(1, 0, 0)
	// gl.Vertex3f(-1, 0, 0)
	// gl.Color3f(0, 1, 0)
	// gl.Vertex3f(0, 1, 0)
	// gl.Color3f(0, 0, 1)
	// gl.Vertex3f(1, 0, 0)
	// gl.End()
}
