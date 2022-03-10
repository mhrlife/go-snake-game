package objects

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"snakegame/engine"
	"time"
)

type Cord struct {
	X int32
	Y int32
}

type Direction int8

const (
	Up Direction = iota
	Right
	Left
	Down
)

type Snake struct {
	body      []Cord
	direction Direction
	createdAt time.Time
	width     int32
	engine    *engine.Engine

	lastMove         time.Time
	waitBetweenMoves time.Duration
}

func NewSnake(e *engine.Engine) engine.GameObject {
	return &Snake{
		createdAt: time.Now(),
		body: []Cord{
			{
				X: 20,
				Y: 20,
			}, {
				X: 20,
				Y: 20,
			},
			{
				X: 20,
				Y: 40,
			},
		},
		direction:        Down,
		width:            20,
		engine:           e,
		lastMove:         time.Now(),
		waitBetweenMoves: 200 * time.Millisecond,
	}
}

func (s *Snake) HasCollision(cord []engine.CollisionCord) bool {
	for _, collisionCord := range cord {
		for i, c := range s.body {
			if i > len(s.body)-3 {
				continue
			}
			if c.Y == collisionCord.Y && c.X == collisionCord.X {
				return true
			}
		}
	}
	return false
}

func (s *Snake) CheckForCollisionCords() []engine.CollisionCord {
	h := s.head()
	return []engine.CollisionCord{
		{
			X: h.X,
			Y: h.Y,
		},
	}
}

func (s *Snake) gameOver() {
	sdl.ShowMessageBox(&sdl.MessageBoxData{
		Window:  s.engine.Window,
		Title:   "Game Over",
		Message: fmt.Sprintf("You earned %d scores.", int(time.Since(s.createdAt).Seconds())*5),
	})
	s.engine.Exit()
}

func (s *Snake) OnCollision(obj interface{}) {
	switch t := obj.(type) {
	case *Food:
		s.engine.Drop(obj.(engine.GameObject))
		s.grow()
	case *Snake:
		s.gameOver()
	default:
		fmt.Println("Unhandled collision for ", t)
	}
}

func (s *Snake) moveByDirection() (int32, int32) {
	switch s.direction {
	case Up:
		return 0, -1
	case Right:
		return 1, 0
	case Left:
		return -1, 0
	case Down:
		return 0, 1
	}
	log.Fatalln("Unknown direction ", s.direction)
	return 0, 0
}

func (s *Snake) popBody() {
	s.body = s.body[1:]
}

func (s *Snake) head() Cord {
	return s.body[len(s.body)-1]
}

func (s *Snake) grow() {
	head := s.head()
	mvX, mvY := s.moveByDirection()
	head.X += mvX * s.width
	head.Y += mvY * s.width

	s.body = append(s.body, head)
}

func (s *Snake) move() {
	if time.Since(s.lastMove) < s.waitBetweenMoves {
		return
	}
	s.lastMove = time.Now()
	s.grow()
	s.popBody()
}

func (s *Snake) handleEvent(e engine.Event) {
	if e.ArrowLeft && s.direction != Right {
		s.direction = Left
	}
	if e.ArrowRight && s.direction != Left {
		s.direction = Right
	}
	if e.ArrowDown && s.direction != Up {
		s.direction = Down
	}
	if e.ArrowUp && s.direction != Down {
		s.direction = Up
	}
}

func (s *Snake) checkIfLeftBoard() {
	h := s.head()
	if h.X < 0 || h.X > s.engine.Width || h.Y < 0 || h.Y > s.engine.Width {
		s.gameOver()
	}
}

func (s *Snake) Render(renderer *sdl.Renderer, e engine.Event) error {
	s.handleEvent(e)
	s.move()
	s.checkIfLeftBoard()

	if err := renderer.SetDrawColor(25, 25, 25, 255); err != nil {
		return err
	}

	var body []sdl.Rect
	for _, cord := range s.body {
		body = append(body, sdl.Rect{
			X: cord.X,
			Y: cord.Y,
			W: s.width,
			H: s.width,
		})
	}
	return renderer.FillRects(body)
}
