package objects

import (
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"snakegame/engine"
	"time"
)

type Background struct {
	r, g, b uint8
	engine  *engine.Engine

	lastFoodSpawn time.Time
}

func NewBackground(r, g, b uint8, e *engine.Engine) *Background {
	return &Background{
		r:             r,
		g:             g,
		b:             b,
		engine:        e,
		lastFoodSpawn: time.Now(),
	}
}

func (b *Background) spawns() {
	if time.Since(b.lastFoodSpawn) > 2*time.Second {
		b.foodSpawn()
	}

}

func (b *Background) foodSpawn() {
	x := rand.Intn(int(b.engine.Width) / 20)
	y := rand.Intn(int(b.engine.Width) / 20)
	x *= 20
	y *= 20
	b.engine.Register(NewFood(Cord{
		X: int32(x),
		Y: int32(y),
	}, b.engine))
	b.lastFoodSpawn = time.Now()
}

func (b *Background) Render(renderer *sdl.Renderer, e engine.Event) error {
	b.spawns()

	if err := renderer.SetDrawColor(b.r, b.g, b.b, 255); err != nil {
		return err
	}

	return renderer.FillRect(&sdl.Rect{
		X: 0,
		Y: 0,
		W: b.engine.Width,
		H: b.engine.Height,
	})
}
