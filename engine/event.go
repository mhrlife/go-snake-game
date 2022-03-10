package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Event struct {
	ArrowRight, ArrowUp, ArrowLeft, ArrowDown bool
}

func (e *Engine) handleEvents() Event {
	var gameEvent Event
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			e.Exit()

		case *sdl.KeyboardEvent:
			e := event.(*sdl.KeyboardEvent)
			switch e.Keysym.Scancode {
			case 82:
				gameEvent.ArrowUp = true
			case 79:
				gameEvent.ArrowRight = true
			case 81:
				gameEvent.ArrowDown = true
			case 80:
				gameEvent.ArrowLeft = true
			}
		}
	}

	return gameEvent
}
