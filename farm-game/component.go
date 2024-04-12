package main

import (
	"log"

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
	img    *ebiten.Image
	src    string
	offset Transform
}

func (ic *ImageComponent) Init() {
	img, _, err := ebitenutil.NewImageFromFile(ic.src)
	if err != nil {
		log.Fatal(err)
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
		//TODO: implement SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image)
		screen.DrawImage(ic.img, options)
	} else {
		log.Fatal("ImageComponent needs to have parent specified to render")
	}
}
