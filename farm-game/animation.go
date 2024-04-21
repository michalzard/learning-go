package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	name   string
	sprite *ImageComponent
}

func NewAnimation(name string, sprite *ImageComponent) *Animation {
	return &Animation{
		name:   name,
		sprite: sprite,
	}
}

func (a *Animation) Render(screen *ebiten.Image) {
	a.sprite.Render(screen)
}

// loop trough image
//

func (a *Animation) Update() {

	// Horizontal

	// grab horizontal width
	imgWidth := a.sprite.img.Bounds().Max.X
	// grab sprite desination data and loop over

	frameWidth := a.sprite.d.size.x

	if a.sprite.d.x < imgWidth {
		a.sprite.d.x += int(frameWidth)
	} else {
		a.sprite.d.x = 0
	}

}

type AnimatorComponent struct {
	Component
	currentAnim *Animation
}

func (ac *AnimatorComponent) SetParent(parent *GameObject) {
	ac.parent = parent
}

func (ac *AnimatorComponent) Init() {

	if ac.currentAnim != nil && ac.parent != nil {
		ac.currentAnim.sprite.SetParent(ac.parent)
		ac.currentAnim.sprite.Init()
	}
}
func (ac *AnimatorComponent) Update() {
	ac.currentAnim.Update()
}
func (ac *AnimatorComponent) Render(screen *ebiten.Image) {
	if ac.currentAnim != nil {
		ac.currentAnim.Render(screen)
	}
}

func (ac *AnimatorComponent) AddAnimation(animation *Animation) {
	ac.currentAnim = animation
}

// func (ac *AnimatorComponent) GetAnimation(animationName string) *Animation {
// 	for _, anim := range ac.animations {
// 		if anim.name == animationName {
// 			return &anim
// 		}
// 	}
// 	return nil
// }

// func (ac *AnimatorComponent) SetAnimation(animationName string) {
// 	animation := ac.GetAnimation(animationName)

// 	if animation != nil {
// 		ac.current = animation
// 	}
// }
