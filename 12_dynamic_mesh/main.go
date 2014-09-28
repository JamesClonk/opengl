package main

import (
	"math"
	"unsafe"

	. "github.com/JamesClonk/opengl/_includes"
	"github.com/go-gl/gl"
	mgl "github.com/go-gl/mathgl/mgl32"
)

var shader *Shader
var time float64
var vertices Vertices
var indices []int32

const w = float64(10)
const h = float64(10)

const vertexShaderSource = `
	#version 130
		in vec4 position;
		in vec4 color;

		varying vec4 vertexColor;

		uniform mat4 model;
		uniform mat4 view;
		uniform mat4 projection;

		void main()	{
			vertexColor = color;
			gl_Position = projection * view * model * position;
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
	app := NewSimpleApp(640, 480, "Go GLFW3 Dynamic Mesh Example", draw)
	defer app.Destroy()

	// create vertices grid mesh
	for i := float64(0); i < w; i++ {
		tu := -1.0 + 2*i/w

		for j := float64(0); j < h; j++ {
			tv := -1.0 + 2*j/h

			vertex := Vertex{
				Position: mgl.Vec4{float32(tu), float32(tv), 0, 1},
				Color:    mgl.Vec4{float32(math.Abs(tu)), 1, float32(math.Abs(tv)), 1},
			}
			vertex.Position = vertex.Position.Mul(5)
			vertices = append(vertices, vertex)
		}
	}

	// create indices vertex array
	for i := float64(0); i < w-1; i++ {
		for j := float64(0); j < h; j++ {
			index := i*h + j

			indices = append(indices, int32(index))
			indices = append(indices, int32(index+h))

			if j == h-1 {
				indices = append(indices, int32(index+h))
			}
		}
		indices = append(indices, int32((i+1)*h))
	}

	shader = NewDynamicShader(&vertices, indices, vertexShaderSource, fragmentShaderSource)

	app.Start()
}

func draw(app *App) {
	time += 0.05

	shader.Use()

	ortho := mgl.Ortho(-app.Ratio, app.Ratio, -1.0, 1.0, -1.0, 1.0)
	shader.Ortho.UniformMatrix4fv(false, ortho)

	// view and projection
	view := mgl.LookAtV(mgl.Vec3{0, 0, 2}, mgl.Vec3{0, 0, 0}, mgl.Vec3{0, 1, 0})
	projection := mgl.Perspective(math.Pi/3.0, app.Ratio, 0.1, -10.0)

	// send view, projection and model to shader
	shader.View.UniformMatrix4fv(false, view)
	shader.Projection.UniformMatrix4fv(false, projection)

	//shader.Model.UniformMatrix4fv(false, mgl.Mat4{})
	model := mgl.HomogRotate3D(0, mgl.Vec3{0, 1, 0})
	shader.Model.UniformMatrix4fv(false, model)

	for i := 0; i < len(vertices); i++ {
		vertices[i].Color = mgl.Vec4{1, 0, 0, 1}
	}

	// bind buffer before substituting data on it
	shader.VertexBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(vertices)*int(unsafe.Sizeof(Vertex{})), vertices)
	gl.DrawElements(gl.LINE_STRIP, len(indices), gl.UNSIGNED_INT, nil)

	for i := float64(0); i < w; i++ {
		for j := float64(0); j < h; j++ {
			index := i*h + j
			pos := vertices[int(index)].Position
			vertices[int(index)].Position = mgl.Vec4{pos.X(), pos.Y(), float32(math.Sin(time + index/(w*h)*math.Pi)), pos.W()}
			vertices[int(index)].Color = mgl.Vec4{float32(i / w), float32(j / h), 0.5, 0.5}
		}
	}

	gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(vertices)*int(unsafe.Sizeof(Vertex{})), vertices)
	gl.DrawElements(gl.TRIANGLE_STRIP, len(indices), gl.UNSIGNED_INT, nil)

	shader.Unuse()
}
