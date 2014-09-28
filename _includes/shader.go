package _includes

import (
	"unsafe"

	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	mgl "github.com/go-gl/mathgl/mgl32"
)

type Shader struct {
	Program       gl.Program
	VertexArray   gl.VertexArray
	VertexBuffer  gl.Buffer
	ElementBuffer gl.Buffer
	Ortho         gl.UniformLocation
	Model         gl.UniformLocation
	View          gl.UniformLocation
	Projection    gl.UniformLocation
}

type Vertex struct {
	Position mgl.Vec4
	Color    mgl.Vec4
}

type Vertices []Vertex

func init() {
	var b byte = 255
	var i int32 = 1234567890
	if int(glh.Sizeof(gl.FLOAT))*4 != int(unsafe.Sizeof(mgl.Vec4{})) {
		panic("wrong float type!")
	} else if int(glh.Sizeof(gl.FLOAT))*8 != int(unsafe.Sizeof(Vertex{})) {
		panic("wrong vertex size!")
	} else if int(glh.Sizeof(gl.UNSIGNED_BYTE)) != int(unsafe.Sizeof(b)) {
		panic("wrong byte size!")
	} else if int(glh.Sizeof(gl.UNSIGNED_INT)) != int(unsafe.Sizeof(i)) {
		panic("wrong int size!")
	}
}

func NewSimpleShader(vertices *Vertices, vertexShaderSource, fragmentShaderSource string) *Shader {
	shader := CreateShader(vertexShaderSource, fragmentShaderSource)

	shader.SetVertexArray()
	shader.SetStaticVertexArrayBuffer(vertices)

	shader.EnableVertexAttributes()
	shader.SetUniformLocations()

	shader.Unuse()
	glh.OpenGLSentinel()

	return shader
}

func NewElementShader(vertices *Vertices, indices []int32, vertexShaderSource, fragmentShaderSource string) *Shader {
	shader := CreateShader(vertexShaderSource, fragmentShaderSource)

	shader.SetVertexArray()
	shader.SetStaticVertexArrayBuffer(vertices)
	shader.SetStaticElementArrayBuffer(indices)

	shader.EnableVertexAttributes()
	shader.SetUniformLocations()

	shader.Unuse()
	glh.OpenGLSentinel()

	return shader
}

func NewDynamicShader(vertices *Vertices, indices []int32, vertexShaderSource, fragmentShaderSource string) *Shader {
	shader := CreateShader(vertexShaderSource, fragmentShaderSource)

	shader.SetVertexArray()
	shader.SetDynamicVertexArrayBuffer(vertices)
	shader.SetStaticElementArrayBuffer(indices)

	shader.EnableVertexAttributes()
	shader.SetUniformLocations()

	shader.Unuse()
	glh.OpenGLSentinel()

	return shader
}

func CreateShader(vertexShaderSource, fragmentShaderSource string) *Shader {
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

func (shader *Shader) Use() {
	shader.Program.Use()
	shader.VertexArray.Bind()
	shader.VertexBuffer.Bind(gl.ARRAY_BUFFER)
	shader.ElementBuffer.Bind(gl.ELEMENT_ARRAY_BUFFER)
}

func (shader *Shader) Unuse() {
	shader.VertexArray.Unbind()
	shader.VertexBuffer.Unbind(gl.ARRAY_BUFFER)
	shader.ElementBuffer.Unbind(gl.ELEMENT_ARRAY_BUFFER)
	shader.Program.Unuse()
}

func (shader *Shader) SetVertexArray() {
	// create vertex array object
	vertexArray := gl.GenVertexArray()
	vertexArray.Bind()
	glh.OpenGLSentinel()

	shader.VertexArray = vertexArray
}

func (shader *Shader) SetStaticVertexArrayBuffer(vertices *Vertices) {
	// create vertex buffer object
	vertexBuffer := gl.GenBuffer()
	vertexBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(*vertices)*int(unsafe.Sizeof(Vertex{})), *vertices, gl.STATIC_DRAW)
	glh.OpenGLSentinel()

	shader.VertexBuffer = vertexBuffer
}

func (shader *Shader) SetDynamicVertexArrayBuffer(vertices *Vertices) {
	// create vertex buffer object
	vertexBuffer := gl.GenBuffer()
	vertexBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(*vertices)*int(unsafe.Sizeof(Vertex{})), *vertices, gl.DYNAMIC_DRAW)
	glh.OpenGLSentinel()

	shader.VertexBuffer = vertexBuffer
}

func (shader *Shader) SetStaticElementArrayBuffer(indices []int32) {
	// create element array buffer object
	elementBuffer := gl.GenBuffer()
	elementBuffer.Bind(gl.ELEMENT_ARRAY_BUFFER)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*int(glh.Sizeof(gl.UNSIGNED_INT)), indices, gl.STATIC_DRAW)
	glh.OpenGLSentinel()

	shader.ElementBuffer = elementBuffer
}

func (shader *Shader) EnableVertexAttributes() {
	position := shader.Program.GetAttribLocation("position")
	position.EnableArray()
	position.AttribPointer(4, gl.FLOAT, false, int(unsafe.Sizeof(Vertex{})), nil)
	glh.OpenGLSentinel()

	color := shader.Program.GetAttribLocation("color")
	color.EnableArray()
	color.AttribPointer(4, gl.FLOAT, false, int(unsafe.Sizeof(Vertex{})), unsafe.Sizeof(mgl.Vec4{}))
	glh.OpenGLSentinel()
}

func (shader *Shader) SetUniformLocations() {
	shader.Ortho = shader.Program.GetUniformLocation("ortho")
	shader.Model = shader.Program.GetUniformLocation("model")
	shader.View = shader.Program.GetUniformLocation("view")
	shader.Projection = shader.Program.GetUniformLocation("projection")
	glh.OpenGLSentinel()
}
