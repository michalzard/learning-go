package main

import "github.com/hajimehoshi/ebiten/v2"

type Vector2 struct {
	x, y float64
}

type Transform struct {
	position Vector2
	scale    Vector2
	rotation float64
}

type GameObject struct {
	parent     *GameObject
	name       string
	transform  Transform
	components []Component
}

func (gameObj GameObject) Init() {
	gameObj.InitComponents()
}
func (gameObj GameObject) Update() {
	gameObj.UpdateComponents()
}
func (gameObj GameObject) Render(screen *ebiten.Image) {
	gameObj.RenderComponents(screen)
}

func (gameObj GameObject) InitComponents() {
	for _, component := range gameObj.components {
		component.Init()
	}
}
func (gameObj GameObject) UpdateComponents() {
	for _, component := range gameObj.components {
		component.Update()
	}
}
func (gameObj GameObject) RenderComponents(screen *ebiten.Image) {
	for _, component := range gameObj.components {
		component.Render(screen)
	}
}

func (gameObj *GameObject) addComponent(c Component) {
	c.SetParent(gameObj)
	gameObj.components = append(gameObj.components, c)
}
func (gameObj *GameObject) addComponents(c ...Component) {
	for _, component := range c {
		component.SetParent(gameObj)
		gameObj.components = append(gameObj.components, component)

	}
}
