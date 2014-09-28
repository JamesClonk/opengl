package main

import (
	"math"

	. "github.com/JamesClonk/opengl/_includes"
	"github.com/go-gl/gl"
	mgl "github.com/go-gl/mathgl/mgl32"
)

var shader *Shader
var time float64

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
	app := NewSimpleApp(640, 480, "Go GLFW3 Model Projection Example", draw)
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
	time += 0.01

	shader.Use()

	ortho := mgl.Ortho(-app.Ratio, app.Ratio, -1.0, 1.0, -1.0, 1.0)
	shader.Ortho.UniformMatrix4fv(false, ortho)

	// view and projection
	view := mgl.LookAtV(mgl.Vec3{0, 0, 2}, mgl.Vec3{0, 0, 0}, mgl.Vec3{0, 1, 0})
	projection := mgl.Perspective(math.Pi/3.0+float32(1+math.Sin(time)), app.Ratio, 0.1, -10.0)

	// send view and projection to shader
	shader.View.UniformMatrix4fv(false, view)
	shader.Projection.UniformMatrix4fv(false, projection)

	// draw lots of triangles (10 * 10 * 10)
	length := float32(10)
	total := length * length * length
	for i := float32(1); i <= length; i++ {
		for j := float32(1); j <= length; j++ {
			for k := float32(1); k <= length; k++ {
				index := i*length*length + j*length + k

				// create transformation matrices
				translate := mgl.Translate3D(-length/2+i, -length/2+j, -length/2+k)
				rotate := mgl.HomogRotate3D(float32(time)*math.Pi, mgl.Vec3{0, 0, 1})
				scale := mgl.Scale3D(1-(index/total), 1-(index/total), 1-(index/total))

				// scale first, then rotate, then translate..
				// which means -> translate * rotate * scale
				model := translate.Mul4(rotate).Mul4(scale)

				// send model to shader
				shader.Model.UniformMatrix4fv(false, model)

				gl.DrawArrays(gl.TRIANGLES, 0, 3)
			}
		}
	}

	shader.Unuse()
}
