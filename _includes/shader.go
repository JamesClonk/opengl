package _includes

import (
	"unsafe"

	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	mgl "github.com/go-gl/mathgl/mgl32"
)

type Shader struct {
	Program     gl.Program
	VertexArray gl.VertexArray
	Ortho       gl.UniformLocation
	Model       gl.UniformLocation
	View        gl.UniformLocation
	Projection  gl.UniformLocation
}

type Vertex struct {
	Position mgl.Vec2
	Color    mgl.Vec4
}

type Vertices []Vertex

func (s Shader) Use() {
	s.Program.Use()
	s.VertexArray.Bind()
}

func (s Shader) Unuse() {
	s.VertexArray.Unbind()
	s.Program.Unuse()
}

func init() {
	if int(glh.Sizeof(gl.FLOAT))*2 != int(unsafe.Sizeof(mgl.Vec2{})) {
		panic("wrong float type!")
	} else if int(glh.Sizeof(gl.FLOAT))*6 != int(unsafe.Sizeof(Vertex{})) {
		panic("wrong vertex size!")
	}
}

func NewSimpleShader(vertices *Vertices, vertexShaderSource, fragmentShaderSource string) *Shader {
	// create shader program
	vertexShader := glh.Shader{gl.VERTEX_SHADER, vertexShaderSource}
	fragmentShader := glh.Shader{gl.FRAGMENT_SHADER, fragmentShaderSource}
	shader := glh.NewProgram(vertexShader, fragmentShader)
	shader.Use()
	glh.OpenGLSentinel()

	// create vertex array object
	vertexArray := gl.GenVertexArray()
	vertexArray.Bind()
	glh.OpenGLSentinel()

	// create vertex buffer object
	triangleBuffer := gl.GenBuffer()
	triangleBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(*vertices)*int(unsafe.Sizeof(Vertex{})), *vertices, gl.STATIC_DRAW)
	glh.OpenGLSentinel()

	// enable vertex attributes
	position := shader.GetAttribLocation("position")
	position.EnableArray()
	position.AttribPointer(2, gl.FLOAT, false, int(unsafe.Sizeof(Vertex{})), nil)

	color := shader.GetAttribLocation("color")
	color.EnableArray()
	color.AttribPointer(4, gl.FLOAT, false, int(unsafe.Sizeof(Vertex{})), unsafe.Sizeof(mgl.Vec2{}))

	// uniform locations
	ortho := shader.GetUniformLocation("ortho")
	model := shader.GetUniformLocation("model")
	view := shader.GetUniformLocation("view")
	projection := shader.GetUniformLocation("projection")
	glh.OpenGLSentinel()

	// unbind
	vertexArray.Unbind()
	shader.Unuse()
	glh.OpenGLSentinel()

	return &Shader{shader, vertexArray, ortho, model, view, projection}
}
