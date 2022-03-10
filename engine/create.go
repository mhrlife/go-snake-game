package engine

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

func (e *Engine) Create() {

	e.Window = e.createWindow()
	defer e.Window.Destroy()

	e.renderer = e.createRenderer(e.Window)
	defer e.renderer.Destroy()

	e.loop()

}

func (e *Engine) createWindow() *sdl.Window {
	window, err := sdl.CreateWindow(e.WindowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		e.Width, e.Height, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalln(err)
	}
	return window
}

func (e *Engine) createRenderer(window *sdl.Window) *sdl.Renderer {
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatalln(err)
	}
	return renderer
}

func (e *Engine) loop() {
	delay := uint32(1000 / e.FPS)

	for e.running {
		e.checkCollision()
		event := e.handleEvents()
		e.render(event)
		sdl.Delay(delay)
	}
}

func (e *Engine) Exit() {
	e.running = false
}
