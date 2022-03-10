package main

import (
	"math/rand"
	"snakegame/engine"
	"snakegame/objects"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	e := engine.NewEngine(engine.Config{
		Width:       300,
		Height:      300,
		WindowTitle: "Snake! - Pure golang game",
		FPS:         60,
	})

	e.Register(objects.NewBackground(218, 247, 166, e))
	e.Register(objects.NewSnake(e))

	e.Create()
}
