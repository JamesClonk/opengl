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

type NormalVertex struct {
	Position mgl.Vec4
	Color    mgl.Vec4
	Normal   mgl.Vec3
}

type NormalTextureVertex struct {
	Position          mgl.Vec4
	Color             mgl.Vec4
	Normal            mgl.Vec3
	TextureCoordinate mgl.Vec2
}

type Vertices []Vertex

type ColorVertices []ColorVertex

type TextureVertices []TextureVertex

type NormalVertices []NormalVertex

type NormalTextureVertices []NormalTextureVertex

func init() {
	if int(glh.Sizeof(gl.FLOAT))*4 != int(unsafe.Sizeof(Vertex{})) {
		panic("wrong vertex size!")
	} else if int(glh.Sizeof(gl.FLOAT))*8 != int(unsafe.Sizeof(ColorVertex{})) {
		panic("wrong color vertex size!")
	} else if int(glh.Sizeof(gl.FLOAT))*6 != int(unsafe.Sizeof(TextureVertex{})) {
		panic("wrong texture vertex size!")
	} else if int(glh.Sizeof(gl.FLOAT))*11 != int(unsafe.Sizeof(NormalVertex{})) {
		panic("wrong normal vertex size!")
	} else if int(glh.Sizeof(gl.FLOAT))*13 != int(unsafe.Sizeof(NormalTextureVertex{})) {
		panic("wrong normal texture vertex size!")
	}
}
