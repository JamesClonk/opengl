package main

import (
	"log"
	"runtime"

	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
)

func main() {
	runtime.LockOSThread()

	if !glfw.Init() {
		panic("Can't init glfw!")
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Go OpenGL Version Example", nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	window.MakeContextCurrent()

	log.Printf("OpenGL Version: %v\n", gl.GetString(gl.VERSION))
	log.Printf("GLSL Version: %v\n", gl.GetString(gl.SHADING_LANGUAGE_VERSION))
}
