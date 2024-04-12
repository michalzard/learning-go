package main

import (
	"fmt"
	"log"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
)

type BaseScene interface {
	Init()
	Update()
	Render(screen *ebiten.Image)
}

// TODO: TURN gos into map that maps gameobject to its own id so we can do quick lookups if needed
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
	BaseScene
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
			fmt.Println(layer.name, gameObj)
			fmt.Println(layer)

			layer.gos = append(layer.gos, gameObj)

			fmt.Println(layer)

			break
		}
	}
}

func NewScene(name string) *Scene {
	return &Scene{
		name:   name,
		layers: []SceneLayer{{name: "Foreground"}, {name: "Background"}, {name: "Overlay"}},
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

func (sm *SceneManager) GetScene(sceneName string) *Scene {
	var selectedScene *Scene

	for _, scene := range sm.scenes {
		if scene.name == sceneName {
			selectedScene = &scene
			break
		}
	}
	if selectedScene == nil {
		return nil
	}

	return selectedScene
}
func (sm *SceneManager) AddScenes(scene ...Scene) {
	if len(scene) > 0 {
		sm.scenes = append(sm.scenes, scene...)
	}
}
