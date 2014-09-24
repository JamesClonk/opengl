package app

import (
	"log"
	"runtime"

	gl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
)

type App struct {
	Window       *glfw.Window
	Width        int
	Height       int
	Title        string
	ViewportFunc func(*glfw.Window)
	DrawFunc     func(*glfw.Window)
	KeyFunc      func(*glfw.Window, glfw.Key, int, glfw.Action, glfw.ModifierKey)
	MouseFunc    func(*glfw.Window, glfw.MouseButton, glfw.Action, glfw.ModifierKey)
	CursorFunc   func(*glfw.Window, float64, float64)
}

func NewSimpleApp(width, height int, title string, drawFunc func(*glfw.Window)) *App {
	return NewApp(width, height, title, SetViewport, drawFunc, OnKeyDown, OnMouseDown, OnMouseMove)
}

func NewApp(width, height int, title string, viewportFunc func(window *glfw.Window), drawFunc func(*glfw.Window), keyFunc func(*glfw.Window, glfw.Key, int, glfw.Action, glfw.ModifierKey), mouseFunc func(*glfw.Window, glfw.MouseButton, glfw.Action, glfw.ModifierKey), cursorFunc func(*glfw.Window, float64, float64)) *App {
	runtime.LockOSThread()

	if !glfw.Init() {
		panic("can't init glfw!")
	}

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	window.SetKeyCallback(keyFunc)
	window.SetMouseButtonCallback(mouseFunc)
	window.SetCursorPositionCallback(cursorFunc)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.LineWidth(3)

	return &App{
		Window:       window,
		Width:        width,
		Height:       height,
		Title:        title,
		ViewportFunc: viewportFunc,
		DrawFunc:     drawFunc,
		KeyFunc:      keyFunc,
		MouseFunc:    mouseFunc,
		CursorFunc:   cursorFunc,
	}
}

func (a App) Start() {
	for !a.Window.ShouldClose() {
		a.ViewportFunc(a.Window)

		gl.ClearColor(0.1, 0.1, 0.1, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		a.DrawFunc(a.Window)

		a.Window.SwapBuffers()
		glfw.PollEvents()
	}
}

func (a App) Destroy() {
	a.Window.Destroy()
	glfw.Terminate()
}

func SetViewport(window *glfw.Window) {
	w, h := window.GetFramebufferSize()
	gl.Viewport(0, 0, w, h)
}

func OnKeyDown(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mod glfw.ModifierKey) {
	log.Printf("Key [%v], Scancode [%v], Action [%v], Modifier [%v]\n", key, scancode, action, mod)

	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
}

func OnMouseDown(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	log.Printf("Mouse Button [%v], Action [%v], Modifier [%v]\n", button, action, mod)
}

func OnMouseMove(window *glfw.Window, x, y float64) {
	log.Printf("Mouse Position [%v, %v]\n", x, y)
}
