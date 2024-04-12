package main

import "fmt"

func init() {
	sm := SceneManager{}
	sm.AddScenes(*NewScene("Farmland"))

	farmland := sm.GetScene("Farmland")

	if farmland != nil {
		player := NewPlayer(Transform{position: Vector2{200, 200}})
		farmland.AddToLayer("Foreground", *player)
		fmt.Printf("%v", farmland.layers)

	}
}
