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
	Texture       gl.Texture
	Ortho         gl.UniformLocation
	Model         gl.UniformLocation
	View          gl.UniformLocation
	Projection    gl.UniformLocation
}

func init() {
	var b byte = 255
	var i int32 = 1234567890
	if int(glh.Sizeof(gl.FLOAT))*4 != int(unsafe.Sizeof(mgl.Vec4{})) {
		panic("wrong float type!")
	} else if int(glh.Sizeof(gl.UNSIGNED_BYTE)) != int(unsafe.Sizeof(b)) {
		panic("wrong byte size!")
	} else if int(glh.Sizeof(gl.UNSIGNED_INT)) != int(unsafe.Sizeof(i)) {
		panic("wrong int size!")
	}
}

func NewSimpleShader(vertices *Vertices, vertexShaderSource, fragmentShaderSource string) *Shader {
	shader := NewShader(vertexShaderSource, fragmentShaderSource)

	shader.SetVertexArray()
	shader.SetVertexArrayBuffer(*vertices, gl.STATIC_DRAW)

	shader.EnableVertexAttributes()
	shader.SetUniformLocations()

	shader.Unuse()
	glh.OpenGLSentinel()

	return shader
}

func NewColoredShader(vertices *ColorVertices, vertexShaderSource, fragmentShaderSource string) *Shader {
	shader := NewShader(vertexShaderSource, fragmentShaderSource)

	shader.SetVertexArray()
	shader.SetVertexArrayBuffer(*vertices, gl.STATIC_DRAW)

	shader.EnableColorVertexAttributes()
	shader.SetUniformLocations()

	shader.Unuse()
	glh.OpenGLSentinel()

	return shader
}

func NewElementShader(vertices *ColorVertices, indices []int32, vertexShaderSource, fragmentShaderSource string) *Shader {
	shader := NewShader(vertexShaderSource, fragmentShaderSource)

	shader.SetVertexArray()
	shader.SetVertexArrayBuffer(*vertices, gl.STATIC_DRAW)
	shader.SetElementArrayBuffer(indices, gl.STATIC_DRAW)

	shader.EnableColorVertexAttributes()
	shader.SetUniformLocations()

	shader.Unuse()
	glh.OpenGLSentinel()

	return shader
}

func NewDynamicShader(vertices *ColorVertices, indices []int32, vertexShaderSource, fragmentShaderSource string) *Shader {
	shader := NewShader(vertexShaderSource, fragmentShaderSource)

	shader.SetVertexArray()
	shader.SetVertexArrayBuffer(*vertices, gl.DYNAMIC_DRAW)
	shader.SetElementArrayBuffer(indices, gl.STATIC_DRAW)

	shader.EnableColorVertexAttributes()
	shader.SetUniformLocations()

	shader.Unuse()
	glh.OpenGLSentinel()

	return shader
}

func NewTexturedShader(vertices *TextureVertices, textureWidth, textureHeight int, data *[]mgl.Vec4, vertexShaderSource, fragmentShaderSource string) *Shader {
	shader := NewShader(vertexShaderSource, fragmentShaderSource)

	shader.SetVertexArray()
	shader.SetVertexArrayBuffer(*vertices, gl.STATIC_DRAW)

	shader.EnableTextureVertexAttributes()
	shader.SetUniformLocations()

	// shader.VertexArray.Unbind()
	// shader.VertexBuffer.Unbind(gl.ARRAY_BUFFER)

	shader.SetTexture(textureWidth, textureHeight, data)

	shader.Unuse()
	glh.OpenGLSentinel()

	return shader
}

func NewShader(vertexShaderSource, fragmentShaderSource string) *Shader {
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
	shader.Texture.Bind(gl.TEXTURE_2D)
	shader.VertexArray.Bind()
	shader.VertexBuffer.Bind(gl.ARRAY_BUFFER)
	shader.ElementBuffer.Bind(gl.ELEMENT_ARRAY_BUFFER)
}

func (shader *Shader) Unuse() {
	shader.VertexArray.Unbind()
	shader.VertexBuffer.Unbind(gl.ARRAY_BUFFER)
	shader.ElementBuffer.Unbind(gl.ELEMENT_ARRAY_BUFFER)
	shader.Texture.Unbind(gl.TEXTURE_2D)
	shader.Program.Unuse()
}

func (shader *Shader) SetVertexArray() {
	// create vertex array object
	vertexArray := gl.GenVertexArray()
	vertexArray.Bind()
	glh.OpenGLSentinel()

	shader.VertexArray = vertexArray
}

func (shader *Shader) SetVertexArrayBuffer(data interface{}, mode gl.GLenum) {
	var size int
	switch d := data.(type) {
	case Vertices:
		size = len(d) * int(unsafe.Sizeof(Vertex{}))
	case ColorVertices:
		size = len(d) * int(unsafe.Sizeof(ColorVertex{}))
	case TextureVertices:
		size = len(d) * int(unsafe.Sizeof(TextureVertex{}))
	default:
		panic("unknown vertex type provided!")
	}

	// create vertex buffer object
	vertexBuffer := gl.GenBuffer()
	vertexBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, size, data, mode)
	glh.OpenGLSentinel()

	shader.VertexBuffer = vertexBuffer
}

func (shader *Shader) SetElementArrayBuffer(indices []int32, mode gl.GLenum) {
	// create element array buffer object
	elementBuffer := gl.GenBuffer()
	elementBuffer.Bind(gl.ELEMENT_ARRAY_BUFFER)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*int(glh.Sizeof(gl.UNSIGNED_INT)), indices, mode)
	glh.OpenGLSentinel()

	shader.ElementBuffer = elementBuffer
}

func (shader *Shader) SetTexture(width, height int, data *[]mgl.Vec4) {
	// create texture
	texture := gl.GenTexture()
	texture.Bind(gl.TEXTURE_2D)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.FLOAT, nil)
	gl.TexSubImage2D(gl.TEXTURE_2D, 0, 0, 0, width, height, gl.RGBA, gl.FLOAT, &((*data)[0]))

	gl.GenerateMipmap(gl.TEXTURE_2D)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	glh.OpenGLSentinel()

	shader.Texture = texture
}

func (shader *Shader) EnableVertexAttribute(name string, length uint, size int, offset interface{}) {
	position := shader.Program.GetAttribLocation(name)
	position.EnableArray()
	position.AttribPointer(length, gl.FLOAT, false, size, offset)
	glh.OpenGLSentinel()
}

func (shader *Shader) EnableVertexAttributes() {
	shader.EnableVertexAttribute("position", 4, int(unsafe.Sizeof(Vertex{})), nil)
}

func (shader *Shader) EnableColorVertexAttributes() {
	shader.EnableVertexAttribute("position", 4, int(unsafe.Sizeof(ColorVertex{})), nil)
	shader.EnableVertexAttribute("color", 4, int(unsafe.Sizeof(ColorVertex{})), unsafe.Sizeof(mgl.Vec4{}))
}

func (shader *Shader) EnableTextureVertexAttributes() {
	shader.EnableVertexAttribute("position", 4, int(unsafe.Sizeof(TextureVertex{})), nil)
	shader.EnableVertexAttribute("textureCoordinate", 2, int(unsafe.Sizeof(TextureVertex{})), unsafe.Sizeof(mgl.Vec4{}))
}

func (shader *Shader) SetUniformLocations() {
	shader.Ortho = shader.Program.GetUniformLocation("ortho")
	shader.Model = shader.Program.GetUniformLocation("model")
	shader.View = shader.Program.GetUniformLocation("view")
	shader.Projection = shader.Program.GetUniformLocation("projection")
	glh.OpenGLSentinel()
}
