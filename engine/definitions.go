package engine

import "github.com/veandco/go-sdl2/sdl"

type CollisionCord struct {
	X, Y int32
}

type GameObject interface {
	Render(renderer *sdl.Renderer, e Event) error
}

type HasCollision interface {
	HasCollision(cord []CollisionCord) bool
}
type OnCollision interface {
	CheckForCollisionCords() []CollisionCord
	OnCollision(obj interface{})
}
