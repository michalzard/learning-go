package main

import (
	"log"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
)

type BaseScene interface {
	Init()
	Update()
	Render(screen *ebiten.Image)
}
type SceneLayer struct {
	name string
	gos  []GameObject
}

func (sl *SceneLayer) AddToLayer(gameObj GameObject) {
	sl.gos = append(sl.gos, gameObj)
}
func (sl *SceneLayer) RemoveFromLayer(gameObj GameObject) {
	for i, layerObj := range sl.gos {
		if layerObj.id == gameObj.id {
			sl.gos = slices.Delete(sl.gos, i, 1)
			return
		}
	}
}

type Scene struct {
	name   string
	layers []SceneLayer
}

func (s *Scene) Init() {
	for _, layer := range s.layers {
		for _, gameObj := range layer.gos {
			gameObj.Init()
		}
	}
}
func (s *Scene) Update() {
	for _, layer := range s.layers {
		for _, gameObj := range layer.gos {
			gameObj.Update()
		}
	}
}
func (s *Scene) Render(screen *ebiten.Image) {
	for _, layer := range s.layers {
		for _, gameObj := range layer.gos {
			gameObj.Render(screen)
		}
	}
}

func (s *Scene) AddToLayer(layerName string, gameObj GameObject) {
	for _, layer := range s.layers {
		if layer.name == layerName {
			layer.gos = append(layer.gos, gameObj)
		}
	}
}

type SceneManager struct {
	scenes  []Scene
	current Scene
}

func (sm *SceneManager) Init(scenes ...Scene) {
	if len(scenes) > 0 {
		sm.scenes = append(sm.scenes, scenes...)
		sm.current = sm.scenes[0]
		// initialize all the available scenes
		for _, scene := range scenes {
			scene.Init()
		}
	} else {
		log.Fatal("Cannot initialize manager without any scene")
	}
}
func (sm *SceneManager) Update() {
	sm.current.Update()
}
func (sm *SceneManager) Render(screen *ebiten.Image) {
	sm.current.Render(screen)
}
func (sm *SceneManager) Switch(scene Scene) {
	sm.current = scene
	sm.current.Init()
}
