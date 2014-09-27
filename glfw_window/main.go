package main

import (
	"log"

	. "github.com/JamesClonk/opengl/_includes"
)

func main() {
	//app := NewApp(640, 480, "Go GLFW3 Window & App Defaults Example", SetViewport, draw, OnKeyDown, OnMouseDown, OnMouseMove, OnError)
	app := NewSimpleApp(640, 480, "Go GLFW3 Window & App Defaults Example", draw)
	defer app.Destroy()

	log.Println("Hello Window!")

	app.Start()

	log.Println("Goodbye Window..")
}

func draw(app *App) {
	// nothing to draw here..
}
