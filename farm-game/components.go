package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Component interface {
	Init()
	Update()
	Render(screen *ebiten.Image)
	SetParent(parent *GameObject)
}
type BaseComponent struct {
	parent *GameObject
}

func (c *BaseComponent) SetParent(parent *GameObject) {
	c.parent = parent
}

type ImageComponent struct {
	Component
	parent *GameObject
	img    *ebiten.Image
	imgSrc string
	offset Transform
}

func (ic *ImageComponent) SetParent(parent *GameObject) {
	ic.parent = parent
}
func (ic *ImageComponent) Init() {
	img, _, err := ebitenutil.NewImageFromFile(ic.imgSrc)
	if err != nil {
		log.Fatal(err)
	}
	ic.img = img

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
