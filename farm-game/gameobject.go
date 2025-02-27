package main

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vector2 struct {
	x, y float64
}

type Transform struct {
	position Vector2
	size     Vector2
	scale    Vector2
	rotation float64
}

var globalId uint = 0

type GameObject struct {
	id         uint
	active     bool
	parent     *GameObject
	name       string
	transform  Transform
	components []BaseComponent
}

func (gameObj *GameObject) setParent(parent *GameObject) {
	gameObj.parent = parent
}

func (gameObj *GameObject) Init() {
	globalId++
	gameObj.id = globalId
	gameObj.active = true
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

func (gameObj *GameObject) AddComponent(c BaseComponent) {
	c.SetParent(gameObj)
	gameObj.components = append(gameObj.components, c)
}

func (gameObj *GameObject) GetComponent(c BaseComponent) BaseComponent {
	for _, comp := range gameObj.components {

		if reflect.TypeOf(comp) == reflect.TypeOf(c) {
			return comp
		}
	}
	return nil
}
