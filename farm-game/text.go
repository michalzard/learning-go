package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Text struct {
	options text.DrawOptions
	content string
}

// TODO: implement proper simple text rendering
func (t *Text) Render(screen *ebiten.Image) {
	// Draw info
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS())
	op := &text.DrawOptions{}
	op.GeoM.Translate(0, 20)
	op.ColorScale.ScaleWithColor(color.White)

	// Example
	fontFamily, err := text.NewGoTextFaceSource(bytes.NewBuffer(fonts.MPlus1pRegular_ttf))

	if err != nil {
		log.Fatal("Unable to load fontface", err)
		return
	}

	font := text.GoTextFace{
		Source: fontFamily,
		Size:   20,
	}

	text.Draw(screen, msg, &font, op)
}
