package _includes

import (
	"unsafe"

	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	mgl "github.com/go-gl/mathgl/mgl32"
)

type Vertex struct {
	Position mgl.Vec4
}

type ColorVertex struct {
	Position mgl.Vec4
	Color    mgl.Vec4
}

type TextureVertex struct {
	Position          mgl.Vec4
	TextureCoordinate mgl.Vec2
}

type Vertices []Vertex

type ColorVertices []ColorVertex

type TextureVertices []TextureVertex

func init() {
	if int(glh.Sizeof(gl.FLOAT))*4 != int(unsafe.Sizeof(Vertex{})) {
		panic("wrong vertex size!")
	} else if int(glh.Sizeof(gl.FLOAT))*8 != int(unsafe.Sizeof(ColorVertex{})) {
		panic("wrong color vertex size!")
	} else if int(glh.Sizeof(gl.FLOAT))*6 != int(unsafe.Sizeof(TextureVertex{})) {
		panic("wrong texture vertex size!")
	}
}
