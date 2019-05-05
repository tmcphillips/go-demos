package main

import (
	"log"
	"runtime"
	"time"

	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang-ui/nuklear/nk"
	"github.com/xlab/closer"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	if err := glfw.Init(); err != nil {
		closer.Fatalln(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	win, err := glfw.CreateWindow(600, 400, "Hello World", nil, nil)
	if err != nil {
		closer.Fatalln(err)
	}

	win.MakeContextCurrent()

	width, height := win.GetSize()
	log.Printf("glfw: created window %dx%d", width, height)

	if err := gl.Init(); err != nil {
		closer.Fatalln("opengl: init failed:", err)
	}

	gl.Viewport(0, 0, int32(width), int32(height))

	ctx := nk.NkPlatformInit(win, nk.PlatformInstallCallbacks)

	atlas := nk.NewFontAtlas()
	nk.NkFontStashBegin(&atlas)
	font := nk.NkFontAtlasAddDefault(atlas, 64, nil)
	nk.NkFontStashEnd()
	if font != nil {
		nk.NkStyleSetFont(ctx, font.Handle())
	}

	exitC := make(chan struct{}, 1)
	doneC := make(chan struct{}, 1)
	closer.Bind(func() {
		close(exitC)
		<-doneC
	})

	fpsTicker := time.NewTicker(time.Second / 30)
	for {
		select {
		case <-exitC:
			nk.NkPlatformShutdown()
			glfw.Terminate()
			fpsTicker.Stop()
			close(doneC)
			return
		case <-fpsTicker.C:
			if win.ShouldClose() {
				close(exitC)
				continue
			}
			glfw.PollEvents()
			draw(win, ctx)
		}
	}
}

const pad = 8

func draw(win *glfw.Window, ctx *nk.Context) {

	nk.NkPlatformNewFrame()
	width, height := win.GetSize()
	bounds := nk.NkRect(0, 0, float32(width), float32(height))
	update := nk.NkBegin(ctx, "", bounds, nk.WindowNoScrollbar)

	if update > 0 {
		cellWidth := int32(width - pad*2)
		cellHeight := float32(height-pad*2) / 2.0
		nk.NkLayoutRowStatic(ctx, cellHeight, cellWidth, 1)
		{
			nk.NkLabel(ctx, "Hello World!", nk.TextCentered)
		}
		nk.NkLayoutRowStatic(ctx, cellHeight, cellWidth, 1)
		{
			if nk.NkButtonLabel(ctx, "Quit") > 0 {
				win.SetShouldClose(true)
			}
		}
	}
	nk.NkEnd(ctx)

	// Draw to viewport
	gl.Viewport(0, 0, int32(width), int32(height))
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.ClearColor(0x10, 0x10, 0x10, 0xff)
	nk.NkPlatformRender(nk.AntiAliasingOn, 4096, 1024)
	win.SwapBuffers()
}
