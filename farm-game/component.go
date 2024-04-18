package main

import (
	"github.com/hajimehoshi/ebiten/v2"
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
