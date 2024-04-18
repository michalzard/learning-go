package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ImageDestination struct {
	x    int
	y    int
	size Vector2
}
type ImageComponent struct {
	Component
	img *ebiten.Image
	src string
	d   *ImageDestination
}

func (ic *ImageComponent) SetParent(parent *GameObject) {
	ic.parent = parent
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
			log.Fatal("ImageComponent img is nil, might have invalid src")
		}
	} else {
		log.Fatal("ImageComponent needs to have parent specified to render")
	}
}

func NewSprite(src string, opts *ImageDestination) *ImageComponent {

	if opts != nil {
		return &ImageComponent{
			src: src,
			d:   opts,
		}
	} else {
		return &ImageComponent{
			src: src,
			d: &ImageDestination{
				x:    0,
				y:    0,
				size: Vector2{32, 32},
			},
		}
	}
}
