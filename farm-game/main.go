package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var player GameObject

type Game struct {
}

func init() {
	tf := Transform{position: Vector2{0, 0}, scale: Vector2{3, 3}}
	player = GameObject{name: "Player", transform: tf}
	player.addComponents(&ImageComponent{imgSrc: "sprites/woodcutter/Woodcutter.png"})
	player.Init()

}

func (g *Game) Update() error {
	// Game logical update
	player.transform.position.x++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 1})
	// all of the game rendering
	player.Render(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Farm Game")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
