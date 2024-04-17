package main

import (
	"image"
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

type ImageDestination struct {
	x    int
	y    int
	size Vector2
}
type ImageComponent struct {
	Component
	img *ebiten.Image
	src string
	d   ImageDestination
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
			destRect := image.Rect(ic.d.x, ic.d.y, ic.d.x+int(ic.d.size.x), ic.d.y+int(ic.d.size.y))
			screen.DrawImage(ic.img.SubImage(destRect).(*ebiten.Image), options)
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
	image ImageComponent
}

type Animation struct {
	name         string
	frames       []AnimationFrame
	currentFrame AnimationFrame
}

func (a *Animation) Render(screen *ebiten.Image) {
	a.currentFrame.image.Render(screen)
}

type AnimatorComponent struct {
	Component
	current    *Animation
	animations []Animation
}

func (ac *AnimatorComponent) SetParent(parent *GameObject) {
	ac.parent = parent
}

func (ac *AnimatorComponent) Init() {}
func (ac *AnimatorComponent) Update() {

}
func (ac *AnimatorComponent) Render(screen *ebiten.Image) {
	if ac.current != nil {
		ac.current.Render(screen)
	}
}
func (ac *AnimatorComponent) GetAnimation(animationName string) *Animation {
	for _, anim := range ac.animations {
		if anim.name == animationName {
			return &anim
		}
	}
	return nil
}

func (ac *AnimatorComponent) SetAnimation(animationName string) {
	animation := ac.GetAnimation(animationName)

	if animation != nil {
		ac.current = animation
	}
}

/*
	Animator Component

	- render method should render currentAnimation
	- update method should loop trough all the available frames
	- sequence of frames slice should be based on state -> "running" , "idle" , "attacking", etc...
	-
*/

/*
	Animation

	- primitive that holds name so that animator can get its reference based on name
	- should hold ref to image
	- method to loop over image with subImage to have precise cutout frame
*/
