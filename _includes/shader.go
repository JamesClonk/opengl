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
	Position mgl.Vec4
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
	var b byte = 255
	if int(glh.Sizeof(gl.FLOAT))*4 != int(unsafe.Sizeof(mgl.Vec4{})) {
		panic("wrong float type!")
	} else if int(glh.Sizeof(gl.FLOAT))*8 != int(unsafe.Sizeof(Vertex{})) {
		panic("wrong vertex size!")
	} else if int(glh.Sizeof(gl.BYTE)) != int(unsafe.Sizeof(b)) { // is this silly? probably..
		panic("wrong byte size!")
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
	vertexBuffer := gl.GenBuffer()
	vertexBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(*vertices)*int(unsafe.Sizeof(Vertex{})), *vertices, gl.STATIC_DRAW)
	glh.OpenGLSentinel()

	// enable vertex attributes
	position := shader.GetAttribLocation("position")
	position.EnableArray()
	position.AttribPointer(4, gl.FLOAT, false, int(unsafe.Sizeof(Vertex{})), nil)
	glh.OpenGLSentinel()

	color := shader.GetAttribLocation("color")
	color.EnableArray()
	color.AttribPointer(4, gl.FLOAT, false, int(unsafe.Sizeof(Vertex{})), unsafe.Sizeof(mgl.Vec4{}))
	glh.OpenGLSentinel()

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

func NewElementShader(vertices *Vertices, indices []byte, vertexShaderSource, fragmentShaderSource string) *Shader {
	shader := createShader(vertexShaderSource, fragmentShaderSource)

	createVertexArray(shader)

	createVertexBuffer(vertices)

	// create element array buffer object
	elementBuffer := gl.GenBuffer()
	elementBuffer.Bind(gl.ELEMENT_ARRAY_BUFFER)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*int(glh.Sizeof(gl.BYTE)), indices, gl.STATIC_DRAW)
	glh.OpenGLSentinel()

	// enable vertex attributes
	position := shader.Program.GetAttribLocation("position")
	position.EnableArray()
	position.AttribPointer(4, gl.FLOAT, false, int(unsafe.Sizeof(Vertex{})), nil)
	glh.OpenGLSentinel()

	color := shader.Program.GetAttribLocation("color")
	color.EnableArray()
	color.AttribPointer(4, gl.FLOAT, false, int(unsafe.Sizeof(Vertex{})), unsafe.Sizeof(mgl.Vec4{}))
	glh.OpenGLSentinel()

	// uniform locations
	ortho := shader.Program.GetUniformLocation("ortho")
	model := shader.Program.GetUniformLocation("model")
	view := shader.Program.GetUniformLocation("view")
	projection := shader.Program.GetUniformLocation("projection")
	glh.OpenGLSentinel()

	// unbind
	shader.Unuse()
	glh.OpenGLSentinel()

	return shader
}

func createShader(vertexShaderSource, fragmentShaderSource string) *Shader {
	// create shader program
	vertexShader := glh.Shader{gl.VERTEX_SHADER, vertexShaderSource}
	fragmentShader := glh.Shader{gl.FRAGMENT_SHADER, fragmentShaderSource}
	shader := glh.NewProgram(vertexShader, fragmentShader)
	shader.Use()
	glh.OpenGLSentinel()

	return &Shader{
		Program: shader,
	}
}

func createVertexArray(shader *Shader) {
	// create vertex array object
	vertexArray := gl.GenVertexArray()
	vertexArray.Bind()
	glh.OpenGLSentinel()

	shader.VertexArray = vertexArray
}

func createVertexBuffer(vertices *Vertices) {
	// create vertex buffer object
	vertexBuffer := gl.GenBuffer()
	vertexBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(*vertices)*int(unsafe.Sizeof(Vertex{})), *vertices, gl.STATIC_DRAW)
	glh.OpenGLSentinel()
}
