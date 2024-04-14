package main

func NewPlayer(spawnPosition Vector2) *GameObject {
	player := GameObject{name: "Player", transform: Transform{position: spawnPosition, scale: Vector2{2, 2}}}
	player.addComponent(&ImageComponent{src: "sprites/woodcutter/Woodcutter.png"})
	player.addComponent(&InputComponent{})
	return &player
}
