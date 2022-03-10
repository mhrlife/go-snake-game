package engine

import "log"

func (e *Engine) render(event Event) {
	for _, object := range e.objects {
		if err := object.Render(e.renderer, event); err != nil {
			log.Fatal(err)
		}
	}

	e.renderer.Present()
}
