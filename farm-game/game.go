package main

func init() {
	farmland := NewScene("Farmland")

	if farmland != nil {
		player := NewPlayer(Vector2{0, 0})
		farmland.AddToLayer("Foreground", player)

	}

	// load into manager
	sm.Init(*farmland)

}
