package engine

import "github.com/veandco/go-sdl2/sdl"

type Engine struct {
	Config
	objects []GameObject

	renderer *sdl.Renderer
	Window   *sdl.Window
	running  bool
}

func NewEngine(c Config) *Engine {
	return &Engine{
		Config:  c,
		running: true,
	}
}

func (e *Engine) Register(object GameObject) {
	e.objects = append(e.objects, object)
}
func (e *Engine) Drop(object GameObject) {
	var nObjects []GameObject
	for _, gameObject := range e.objects {
		if gameObject != object {
			nObjects = append(nObjects, gameObject)
		}
	}
	e.objects = nObjects
}
