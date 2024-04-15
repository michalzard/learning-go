package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Text struct {
	options    text.DrawOptions
	content    string
	transform  Transform
	color      color.Color
	fontSize   float64
	fontFamily *text.GoTextFaceSource
}

func (t *Text) Render(screen *ebiten.Image) {
	// Draw info
	t.options.GeoM.Translate(t.transform.position.x, t.transform.position.y)
	if t.color != nil {
		t.options.ColorScale.ScaleWithColor(t.color)
	} else {
		t.options.ColorScale.ScaleWithColor(color.White)
	}

	// Example
	if t.fontFamily == nil {
		fontFamily, err := text.NewGoTextFaceSource(bytes.NewBuffer(fonts.MPlus1pRegular_ttf))
		if err != nil {
			log.Fatal(err)
			return
		}
		t.fontFamily = fontFamily
	}

	if t.fontSize == 0.0 {
		t.fontSize = 16.0
	}

	font := text.GoTextFace{
		Source: t.fontFamily,
		Size:   t.fontSize,
	}

	text.Draw(screen, t.content, &font, &t.options)
}
