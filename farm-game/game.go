package main

func init() {
	farmland := NewScene("Farmland")

	if farmland != nil {
		player := NewPlayer(Vector2{300, 200})
		farmland.AddToLayer("Foreground", player)

	}

	// load into manager
	sm.Init(*farmland)

}
