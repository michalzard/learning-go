package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type BaseComponent interface {
	Init()
	Update()
	Render(screen *ebiten.Image)
	SetParent(parent *GameObject)
}

type Component struct {
	parent *GameObject
}

func (c *Component) SetParent(parent *GameObject) {
	c.parent = parent
}

type ImageComponent struct {
	Component
	img *ebiten.Image
	src string
	// offset Transform
}

func (ic *ImageComponent) Init() {
	img, _, err := ebitenutil.NewImageFromFile(ic.src)
	if err != nil {
		log.Fatal(err)
		return
	}
	ic.img = img
}

func (ic *ImageComponent) Update() {
	// Implement update logic here if needed
}

func (ic *ImageComponent) Render(screen *ebiten.Image) {
	if ic.parent != nil {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(ic.parent.transform.position.x, ic.parent.transform.position.y)
		options.GeoM.Scale(ic.parent.transform.scale.x, ic.parent.transform.scale.y)
		options.GeoM.Rotate(ic.parent.transform.rotation)
		if ic.img != nil {
			screen.DrawImage(ic.img, options)
		} else {
			log.Fatal("ImageComponent needs to have img pointer")
		}
	} else {
		log.Fatal("ImageComponent needs to have parent specified to render")
	}
}

type InputComponent struct {
	Component
}

func (ic *InputComponent) SetParent(parent *GameObject) {
	ic.parent = parent
}
func (ic InputComponent) Init() {}
func (ic InputComponent) Update() {
	if ic.parent != nil {
		w := ebiten.IsKeyPressed(ebiten.KeyW)
		a := ebiten.IsKeyPressed(ebiten.KeyA)
		s := ebiten.IsKeyPressed(ebiten.KeyS)
		d := ebiten.IsKeyPressed(ebiten.KeyD)

		velocity := 3.5

		// Calculate the speed for diagonal movement
		diagonalSpeed := velocity / math.Sqrt2

		// Calculate the speed for cardinal movement
		cardinalSpeed := velocity

		// Initialize movement vectors
		moveX := 0.0
		moveY := 0.0

		if a {
			moveX -= cardinalSpeed
		}
		if d {
			moveX += cardinalSpeed
		}
		if w {
			moveY -= cardinalSpeed
		}
		if s {
			moveY += cardinalSpeed
		}

		// If moving diagonally, normalize the movement vector
		if moveX != 0 && moveY != 0 {
			moveX *= diagonalSpeed / cardinalSpeed
			moveY *= diagonalSpeed / cardinalSpeed
		}

		// Apply movement
		ic.parent.transform.position.x += moveX
		ic.parent.transform.position.y += moveY

	}
}
func (ic InputComponent) Render(screen *ebiten.Image) {}

type AnimationFrame struct {
	frame      Vector2
	frameCount uint
	size       Vector2
}

type Animation struct {
	sprites []ImageComponent
	current uint
	frames  []AnimationFrame
}

func (a *Animation) Animate() {
	// Implement spritesheet looping for animation
}
