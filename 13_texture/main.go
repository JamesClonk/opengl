package main

import (
	"math"

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
		in vec2 textureCoordinate;

		varying vec2 texCoord;

		uniform mat4 model;
		uniform mat4 view;
		uniform mat4 projection;

		void main()	{
			texCoord = textureCoordinate;
			gl_Position = projection * view * model * position;
		}
`

const fragmentShaderSource = `
	#version 130
		uniform sampler2D texture;
		varying vec2 texCoord;

		void main(void){
			gl_FragColor = texture2D(texture, texCoord);
		}
`

func main() {
	app := NewSimpleApp(640, 480, "Go GLFW3 Texture Example", draw)
	defer app.Destroy()

	slab := TextureVertices{
		TextureVertex{
			Position:          mgl.Vec4{-1, -1, 0, 1},
			TextureCoordinate: mgl.Vec2{0, 0},
		},
		TextureVertex{
			Position:          mgl.Vec4{-1, 1, 0, 1},
			TextureCoordinate: mgl.Vec2{0, 1},
		},
		TextureVertex{
			Position:          mgl.Vec4{1, 1, 0, 1},
			TextureCoordinate: mgl.Vec2{1, 1},
		},
		TextureVertex{
			Position:          mgl.Vec4{1, -1, 0, 1},
			TextureCoordinate: mgl.Vec2{1, 0},
		},
	}

	var w float32 = 24
	var h float32 = 24
	var data []mgl.Vec4
	var checkered bool

	for i := float32(0); i < w; i++ {
		r := i / w
		for j := float32(0); j < h; j++ {
			g := j / h
			data = append(data, mgl.Vec4{r, g, 0, btof(checkered)})
			checkered = !checkered
		}
		checkered = !checkered
	}

	shader = NewTexturedShader(&slab, int(w), int(h), &data, vertexShaderSource, fragmentShaderSource)

	app.Start()
}

func btof(b bool) float32 {
	if b {
		return 1
	}
	return 0
}

func draw(app *App) {
	time += 0.05

	shader.Use()

	// view, projection and model
	view := mgl.LookAtV(mgl.Vec3{0, 0, 2}, mgl.Vec3{0, 0, 0}, mgl.Vec3{0, 1, 0})
	projection := mgl.Perspective(math.Pi/3.0, app.Ratio, 0.1, -10.0)
	model := mgl.HomogRotate3D(0, mgl.Vec3{0, 0, 0})

	// send view, projection and model to shader
	shader.View.UniformMatrix4fv(false, view)
	shader.Projection.UniformMatrix4fv(false, projection)
	shader.Model.UniformMatrix4fv(false, model)

	// draw
	gl.DrawArrays(gl.QUADS, 0, 4)

	shader.Unuse()
}
