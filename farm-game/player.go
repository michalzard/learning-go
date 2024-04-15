package main

func NewPlayer(spawnPosition Vector2) *GameObject {
	player := GameObject{name: "Player", transform: Transform{position: spawnPosition, scale: Vector2{2, 2}}}
	playerSprite := NewSprite("sprites/woodcutter/Woodcutter.png")
	playerSprite.d.size.y = 48
	player.AddComponent(playerSprite)
	player.AddComponent(&InputComponent{})
	return &player
}
