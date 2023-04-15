package EbitNew

import (
	//"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewEntity() *Entity {

	new := Entity{}
	new.id = entityId
	//new.Position = nil
	entityId++
	entities = append(entities, &new)
	return &new
}

var entityId int
var entities []*Entity

func init() {
	entityId = 0
}

type Entity struct {
	id       int
	Position *Position
	Image    *Image
	Wasd     *Wasd
	Gravity  *Gravity
	Rect     *Rect
}

func (s *Entity) NewPosition(x float64, y float64) *Entity {
	new := Position{}
	new.Entity = s
	new.X = x
	new.Y = y
	s.Position = &new
	return s
}

func (s *Entity) NewImage(image *ebiten.Image) *Entity {
	new := Image{}
	new.Entity = s
	new.Image = image
	s.Image = &new
	Images = append(Images, &new)
	return s
}

func (s *Entity) NewWasd(speed float64) *Entity {
	new := Wasd{}
	new.Entity = s
	new.Speed = speed
	s.Wasd = &new
	Wasds = append(Wasds, &new)
	return s
}

func (s *Entity) NewGravity(speed float64) *Entity {
	new := Gravity{}
	new.Entity = s
	new.Speed = speed
	s.Gravity = &new
	Gravitys = append(Gravitys, &new)
	return s
}

func (s *Entity) NewRect(width int, height int) *Entity {
	new := Rect{}
	new.Entity = s
	new.Width = width
	new.Height = height
	//new.Rect = image.Rect(int(s.Position.X), int(s.Position.Y), int(s.Position.X)+width, int(s.Position.Y)-height)
	s.Rect = &new
	Rects = append(Rects, &new)
	return s
}

func (s *Entity) NewFloor() *Entity {
	new := Floor{}
	new.Entity = s
	Floors = append(Floors, &new)
	return s
}

type test struct {
	*yo
}

type yo struct {
	hej int
}

func init() {
	hmm := test{}
	hmm.yo = &yo{}
	hmm.hej = 5
	println(hmm.hej)
}
