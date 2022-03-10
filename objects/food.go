package objects

import (
	"github.com/veandco/go-sdl2/sdl"
	"snakegame/engine"
	"time"
)

type Food struct {
	Cord   Cord
	engine *engine.Engine

	Width     int32
	createdAt time.Time
	lifeSpan  float64
}

func (f *Food) HasCollision(cords []engine.CollisionCord) bool {
	for _, cord := range cords {
		if cord.X == f.Cord.X && cord.Y == f.Cord.Y {
			return true
		}
	}
	return false
}

func NewFood(cord Cord, e *engine.Engine) engine.GameObject {
	return &Food{
		engine:    e,
		Cord:      cord,
		Width:     20,
		createdAt: time.Now(),
		lifeSpan:  7,
	}
}

func (f *Food) Render(renderer *sdl.Renderer, e engine.Event) error {
	lifeRemainingPercentage := 1 - time.Since(f.createdAt).Seconds()/f.lifeSpan
	// life is over!
	if lifeRemainingPercentage <= 0 {
		f.engine.Drop(f)
		return nil
	}

	width := int32(float64(f.Width) * lifeRemainingPercentage)
	x := f.Cord.X + (f.Width-width)/2
	y := f.Cord.Y + (f.Width-width)/2
	if err := renderer.SetDrawColor(199, 0, 57, 255); err != nil {
		return err
	}
	return renderer.FillRect(&sdl.Rect{
		X: x,
		Y: y,
		W: width,
		H: width,
	})
}
