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
		in vec3 norm;

		varying vec4 vertexColor;

		uniform mat4 model;
		uniform mat4 view;
		uniform mat4 projection;
		uniform mat3 normal;

		vec4 doColor() {
			vec3 normalized  = normalize(normal * normalize(norm));
			vec3 light = normalize(vec3(1.0, 1.0, 1.0));
			float df = max(dot(normalized, light), 0.0);
			return vec4((color * df).xyz, 1.0);
		}

		void main()	{
			vertexColor = doColor();
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
	app := NewSimpleApp(640, 480, "Go GLFW3 Normal Lighting Example", draw)
	defer app.Destroy()

	cube := NormalVertices{
		NormalVertex{
			Position: mgl.Vec4{1, -1, 1, 1},
			Color:    mgl.Vec4{1, 1, 0, 1},
			Normal:   mgl.Vec3{1, -1, 1},
		},
		NormalVertex{
			Position: mgl.Vec4{1, 1, 1, 1},
			Color:    mgl.Vec4{0, 1, 0, 1},
			Normal:   mgl.Vec3{1, 1, 1},
		},
		NormalVertex{
			Position: mgl.Vec4{-1, 1, 1, 1},
			Color:    mgl.Vec4{1, 1, 0, 1},
			Normal:   mgl.Vec3{-1, 1, 1},
		},
		NormalVertex{
			Position: mgl.Vec4{-1, -1, 1, 1},
			Color:    mgl.Vec4{1, 0, 0, 1},
			Normal:   mgl.Vec3{-1, -1, 1},
		},
		NormalVertex{
			Position: mgl.Vec4{1, -1, -1, 1},
			Color:    mgl.Vec4{0, 1, 0, 1},
			Normal:   mgl.Vec3{1, -1, -1},
		},
		NormalVertex{
			Position: mgl.Vec4{1, 1, -1, 1},
			Color:    mgl.Vec4{0, 0, 1, 1},
			Normal:   mgl.Vec3{1, 1, -1},
		},
		NormalVertex{
			Position: mgl.Vec4{-1, 1, -1, 1},
			Color:    mgl.Vec4{1, 0, 0, 1},
			Normal:   mgl.Vec3{-1, 1, -1},
		},
		NormalVertex{
			Position: mgl.Vec4{-1, -1, -1, 1},
			Color:    mgl.Vec4{0, 0, 1, 1},
			Normal:   mgl.Vec3{-1, -1, -1},
		},
	}
	/*
	       //6-------------/5
	     //  .           // |
	   //2--------------1   |
	   //    .          |   |
	   //    .          |   |
	   //    .          |   |
	   //    .          |   |
	   //    7.......   |   /4
	   //               | //
	   //3--------------/0
	*/

	indices := []int32{
		0, 1, 2, 3, // front
		7, 6, 5, 4, // back
		3, 2, 6, 7, // left
		4, 5, 1, 0, // right
		1, 5, 6, 2, // top
		4, 0, 3, 7, // bottom
	}

	shader = NewNormalShader(&cube, indices, vertexShaderSource, fragmentShaderSource)

	app.Start()
}

func draw(app *App) {
	time += 0.01

	shader.Use()

	ortho := mgl.Ortho(-app.Ratio, app.Ratio, -1.0, 1.0, -1.0, 1.0)
	shader.Ortho.UniformMatrix4fv(false, ortho)

	// view and projection
	view := mgl.LookAtV(mgl.Vec3{0, 0, 5}, mgl.Vec3{0, 0, 0}, mgl.Vec3{0, 1, 0})
	projection := mgl.Perspective(math.Pi/3.0, app.Ratio, 0.1, -10.0)

	// send view and projection to shader
	shader.View.UniformMatrix4fv(false, view)
	shader.Projection.UniformMatrix4fv(false, projection)

	// transformation matrix for rotation
	model := mgl.HomogRotate3D(float32(time), mgl.Vec3{0, 1, 0})
	shader.Model.UniformMatrix4fv(false, model)

	// calculate normal matrix and send to shader
	normal := view.Mul4(model).Mat3().Inv().Transpose()
	shader.Normal.UniformMatrix3fv(false, normal)

	gl.DrawElements(gl.QUADS, 24, gl.UNSIGNED_INT, nil)

	shader.Unuse()
}
