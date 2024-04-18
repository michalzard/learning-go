package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

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
