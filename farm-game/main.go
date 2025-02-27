package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var sm = SceneManager{}

type Game struct {
}

func (g *Game) Update() error {
	sm.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	sm.Render(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Harvest Homestead")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
