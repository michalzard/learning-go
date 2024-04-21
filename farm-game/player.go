package main

func NewPlayer(spawnPosition Vector2) *GameObject {
	player := GameObject{name: "Player", transform: Transform{position: spawnPosition, scale: Vector2{2, 2}}}

	playerAnimator := &AnimatorComponent{}
	playerSprite := NewSprite("sprites/woodcutter/Woodcutter_walk.png", nil)
	playerSprite.d.size = Vector2{48, 48}

	playerAnimator.AddAnimation(NewAnimation("Idle", playerSprite))

	player.AddComponent(playerAnimator)
	player.AddComponent(&InputComponent{})
	return &player
}
