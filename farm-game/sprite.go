package main

// &ImageComponent{src: ""}

func NewSprite(src string) *ImageComponent {
	return &ImageComponent{
		src: src,
		d: ImageDestination{
			x:    0,
			y:    0,
			size: Vector2{32, 32},
		},
	}
}
