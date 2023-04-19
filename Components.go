package EbitNew

import (
	//"log"
	//"image"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewEntity[T any]() *Entity[T] {

	new := Entity[T]{}
	new.id = entityId
	//new.Position = nil
	entityId++
	//entities = append(entities, &new)
	return &new
}

var entityId int
var entities []*Entity[T]

func init() {
	entityId = 0
}

type T any
type Entity[T any] struct {
	id       int
	Position *Position[T]
	Image    *Image
	Wasd     *Wasd
	Gravity  *Gravity
	Rect     *Rect
	Ext      *T
}

func (s *Entity[T]) NewExt(new any) *T {
	s.Ext = new.(*T)
	new2 := new.(Ext[T])
	new2.SetEntity(s)
	//new.SetEntity(s)
	return s.Ext
}

type Ext[T any] interface {
	SetEntity(*Entity[T])
}

func (s *Entity[T]) NewPosition(x float64, y float64) *Entity[T] {
	new := Position[T]{}
	new.Entity = s
	new.X = x
	new.Y = y
	s.Position = &new
	return s
}

func (s *Entity[any]) GetPosition() (float64, float64) {
	if s.Position == nil {
		fmt.Println("\n")
		errorString := fmt.Sprintf("\nEntity with id %d has no component Position", s.id)
		panic(errorString)
		//log.Fatal("yo")
		//return 0, 0
		//panic("Entity has no component named Position.")
	} else {
		return s.Position.X, s.Position.Y
	}
}

func (s *Entity[any]) NewImage(image *ebiten.Image) *Entity[any] {
	new := Image{}
	//new.Entity = s
	new.Image = image
	s.Image = &new
	Images = append(Images, &new)
	return s
}

func (s *Entity[any]) NewWasd(speed float64) *Entity[any] {
	new := Wasd{}
	//new.Entity = s
	new.Speed = speed
	s.Wasd = &new
	Wasds = append(Wasds, &new)
	return s
}

func (s *Entity[any]) NewGravity(speed float64) *Entity[any] {
	new := Gravity{}
	//new.Entity = s
	new.Speed = speed
	s.Gravity = &new
	Gravitys = append(Gravitys, &new)
	return s
}

func (s *Entity[any]) NewRect(width int, height int) *Entity[any] {
	new := Rect{}
	//new.Entity = s
	new.Width = width
	new.Height = height
	//new.Rect = image.Rect(int(s.Position.X), int(s.Position.Y), int(s.Position.X)+width, int(s.Position.Y)-height)
	s.Rect = &new
	Rects = append(Rects, &new)
	return s
}

func (s *Entity[any]) NewFloor() *Entity[any] {
	new := Floor{}
	//new.Entity = s
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
