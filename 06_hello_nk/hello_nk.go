package main

import (
	"runtime"

	// "github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	defer glfw.Terminate()

	// if err := gl.Init(); err != nil {
	// 	panic(err)
	// }

	window, err := glfw.CreateWindow(600, 400, "Hello World", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
