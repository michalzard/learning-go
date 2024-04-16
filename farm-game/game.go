package main

func init() {
	farmland := NewScene("Farmland")

	if farmland != nil {
		player := NewPlayer(Vector2{0, 0})
		player.AddComponent(&AnimatorComponent{})
		farmland.AddToLayer("Foreground", player)

	}

	// load into manager
	sm.Init(*farmland)

}
