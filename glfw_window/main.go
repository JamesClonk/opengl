package main

import (
	"log"

	. "github.com/JamesClonk/opengl/_includes"
	glfw "github.com/go-gl/glfw3"
)

func main() {
	//app := NewApp(640, 480, "Go GLFW3 Window & App Defaults Example", SetViewport, draw, OnKeyDown, OnMouseDown, OnMouseMove, OnError)
	app := NewSimpleApp(640, 480, "Go GLFW3 Window & App Defaults Example", draw)
	defer app.Destroy()

	log.Println("Hello Window!")

	app.Start()

	log.Println("Goodbye Window..")
}

func draw(window *glfw.Window) {
	// nothing to draw here..
}
