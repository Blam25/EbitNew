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

func (s *Wasd) Move() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		for _, z := range entities {
			if z.Wasd == nil {
				z.Position.Y += s.Speed
			}
		}
		//print("hej")
		//s.Entity.Position.Y += -s.Speed
		//y = C.Components.Player.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		atFloor := false
		for _, z := range Floors {
			//if z.Entity != s.Entity {
			if s.Entity.Rect.Bottom.Overlaps(z.Entity.Rect.Top) && z.Entity != s.Entity {
				atFloor = true
				//}
			}
		}
		if !atFloor {
			for _, z := range entities {
				if z.Wasd == nil {
					z.Position.Y += -s.Speed
				}
			}
		}
		//s.Entity.Position.Y += s.Speed
		//y = -C.Components.Player.MoveSpeed

	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		atFloor := false
		for _, z := range Floors {
			//if z.Entity != s.Entity {
			if s.Entity.Rect.Right.Overlaps(z.Entity.Rect.Left) && z.Entity != s.Entity {
				atFloor = true
				//}
			}
		}
		if !atFloor {
			for _, z := range entities {
				if z.Wasd == nil {
					z.Position.X += -s.Speed
				}
			}
		}
		//s.Entity.Position.X += s.Speed
		//x = -C.Components.Player.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		for _, z := range entities {
			if z.Wasd == nil {
				z.Position.X += s.Speed
			}
		}
		//s.Entity.Position.X += -s.Speed
		//x = C.Components.Player.MoveSpeed
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
		for _, z := range entities {
			if z.Wasd == nil {
				z.Position.Y += -s.Speed
			}
		}
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
		int(s.Entity.Position.Y)-s.Height)
	s.Top = image.Rect(int(s.Entity.Position.X)-s.Width/2, int(s.Entity.Position.Y)-s.Height, int(s.Entity.Position.X)+s.Width, int(s.Entity.Position.Y)-s.Height+6)
	s.Bottom = image.Rect(int(s.Entity.Position.X)-s.Width/2, int(s.Entity.Position.Y), int(s.Entity.Position.X)+s.Width, int(s.Entity.Position.Y)+6)
	s.Right = image.Rect(int(s.Entity.Position.X)+s.Width-6, int(s.Entity.Position.Y)-s.Height/2, int(s.Entity.Position.X)+s.Width, int(s.Entity.Position.Y)+s.Height)
	s.Left = image.Rect(int(s.Entity.Position.X)+6-s.Width, int(s.Entity.Position.Y)-s.Height/2, int(s.Entity.Position.X)-s.Width, int(s.Entity.Position.Y)+s.Height)
}

var Floors []*Floor

type Floor struct {
	Entity *Entity
}
