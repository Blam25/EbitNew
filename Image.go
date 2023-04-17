package EbitNew

import (
	"image"
	I "image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Position struct {
	Entity *Entity
	X      float64
	Y      float64
}

type Image struct {
	Entity *Entity
	Image  *ebiten.Image
	Op     ebiten.DrawImageOptions
}

var Images []*Image

func (s *Image) Draw(screen *ebiten.Image) {
	s.Op.GeoM.Reset()
	s.Op.GeoM.Translate(s.Entity.Position.X, s.Entity.Position.Y)
	screen.DrawImage(s.Image, &s.Op)
}

type Wasd struct {
	Entity *Entity
	Speed  float64
}

var Wasds []*Wasd

func (s *Wasd) movePlayer(speedx, speedy float64) {
	for _, z := range entities {
		if z.Wasd == nil {
			z.Position.X += speedx
			z.Position.Y += speedy
		}
	}

}

type Direction int

const (
	Left Direction = iota
	Up
	Right
	Down
)

func (s *Rect) colliding(direction Direction) bool {
	colliding := false
	for _, z := range Floors {
		var rect1 image.Rectangle
		var rect2 image.Rectangle
		switch direction {
		case Left:
			rect1 = s.Left
			rect2 = z.Entity.Rect.Right
		case Up:
			rect1 = s.Top
			rect2 = z.Entity.Rect.Bottom
		case Right:
			rect1 = s.Right
			rect2 = z.Entity.Rect.Left
		case Down:
			rect1 = s.Bottom
			rect2 = z.Entity.Rect.Top
		}
		if rect1.Overlaps(rect2) && z.Entity != s.Entity {
			colliding = true
			if direction == Down && s.Entity.Wasd != nil {
				moveY := s.Entity.Position.Y - z.Entity.Position.Y + float64(s.Entity.Rect.Height) - 2
				s.Entity.Wasd.movePlayer(0, moveY)
			}
		}
	}
	return colliding
}

func (s *Wasd) Move() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		s.movePlayer(0, s.Speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if !s.Entity.Rect.colliding(Down) {
			s.movePlayer(0, -s.Speed)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		if !s.Entity.Rect.colliding(Right) {
			s.movePlayer(-s.Speed, 0)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		if !s.Entity.Rect.colliding(Left) {
			s.movePlayer(s.Speed, 0)
		}
	}
}

type Gravity struct {
	Entity *Entity
	Speed  float64
}

var Gravitys []*Gravity

func (s *Gravity) Pull() {

	atFloor := false
	for _, z := range Floors {
		//if z.Entity != s.Entity {
		if s.Entity.Rect.Bottom.Overlaps(z.Entity.Rect.Top) && z.Entity != s.Entity {
			atFloor = true
			//}
		}
	}
	if !atFloor && s.Entity.Wasd == nil {
		s.Entity.Position.Y += s.Speed
	} else if !atFloor && s.Entity.Wasd != nil {
		s.Entity.Wasd.movePlayer(0, -s.Speed)
	}

}

var Rects []*Rect

type Rect struct {
	Entity *Entity
	Rect   I.Rectangle
	Width  int
	Height int
	Top    I.Rectangle
	Right  I.Rectangle
	Bottom I.Rectangle
	Left   I.Rectangle
}

func (s *Rect) setRect() {
	s.Rect = image.Rect(int(s.Entity.Position.X),
		int(s.Entity.Position.Y),
		int(s.Entity.Position.X)+s.Width,
		int(s.Entity.Position.Y)+s.Height)
	//.Add(image.Point{int(s.Entity.Position.X), int(s.Entity.Position.X)}
	s.Top = image.Rect(
		int(s.Entity.Position.X)+2,
		int(s.Entity.Position.Y)-s.Height,
		int(s.Entity.Position.X)+s.Width-2,
		int(s.Entity.Position.Y)-s.Height+10)
	s.Bottom = image.Rect(
		int(s.Entity.Position.X)+2,
		int(s.Entity.Position.Y),
		int(s.Entity.Position.X)+s.Width-2,
		int(s.Entity.Position.Y)-10)
	s.Right = image.Rect(
		int(s.Entity.Position.X)+s.Width-10,
		int(s.Entity.Position.Y)+2,
		int(s.Entity.Position.X)+s.Width,
		int(s.Entity.Position.Y)+s.Height-2)
	s.Left = image.Rect(
		int(s.Entity.Position.X),
		int(s.Entity.Position.Y)+2,
		int(s.Entity.Position.X)+10,
		int(s.Entity.Position.Y)+s.Height-2)
}

var Floors []*Floor

type Floor struct {
	Entity *Entity
}
