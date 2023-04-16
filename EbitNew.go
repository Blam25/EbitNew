package EbitNew

import "github.com/hajimehoshi/ebiten/v2"

type ECSystems struct {
	First       []System
	Second      []System
	Third       []System
	Fourth      []System
	Fifth       []System
	Sixth       []System
	Seventh     []System
	Eighth      []System
	Ninth       []System
	Tenth       []System
	DrawFirst   []SystemDraw
	DrawSecond  []SystemDraw
	DrawThird   []SystemDraw
	DrawFourth  []SystemDraw
	DrawFifth   []SystemDraw
	DrawSixth   []SystemDraw
	DrawSeventh []SystemDraw
	DrawEighth  []SystemDraw
	DrawNinth   []SystemDraw
	DrawTenth   []SystemDraw
}

var Systems *ECSystems = &ECSystems{}
var SystemsMap map[string]System
var SystemsRenderMap map[string]SystemDraw
var G Game = Game{}

type System interface {
	Execute()
}

type SystemDraw interface {
	Execute(*ebiten.Image)
}

type Game struct{}

func (g *Game) Update() error {

	//Wasds[0].Move()
	for _, s := range Rects {
		s.setRect()
	}

	for _, s := range Wasds {
		s.Move()
	}

	for _, s := range Gravitys {
		s.Pull()
	}

	for _, s := range Systems.First {
		s.Execute()
	}
	for _, s := range Systems.Second {
		s.Execute()
	}
	for _, s := range Systems.Third {
		s.Execute()
	}
	for _, s := range Systems.Fourth {
		s.Execute()
	}
	for _, s := range Systems.Fifth {
		s.Execute()
	}
	for _, s := range Systems.Sixth {
		s.Execute()
	}
	for _, s := range Systems.Seventh {
		s.Execute()
	}
	for _, s := range Systems.Eighth {
		s.Execute()
	}
	for _, s := range Systems.Ninth {
		s.Execute()
	}
	for _, s := range Systems.Tenth {
		s.Execute()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, s := range Images {
		s.Draw(screen)
	}

	for _, s := range Systems.DrawFirst {
		s.Execute(screen)
	}
	for _, s := range Systems.DrawSecond {
		s.Execute(screen)
	}
	for _, s := range Systems.DrawThird {
		s.Execute(screen)
	}
	for _, s := range Systems.DrawFourth {
		s.Execute(screen)
	}
	for _, s := range Systems.DrawFifth {
		s.Execute(screen)
	}
	for _, s := range Systems.DrawSixth {
		s.Execute(screen)
	}
	for _, s := range Systems.DrawSeventh {
		s.Execute(screen)
	}
	for _, s := range Systems.DrawEighth {
		s.Execute(screen)
	}
	for _, s := range Systems.DrawNinth {
		s.Execute(screen)
	}
	for _, s := range Systems.DrawTenth {
		s.Execute(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
