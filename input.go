package main

import "github.com/veandco/go-sdl2/sdl"

func (p *Player) update() {

	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_D] == 1 {
		//move right
		p.x += PlayerSpeed
	} else if keys[sdl.SCANCODE_A] == 1 {
		//move left
		p.x -= PlayerSpeed
	} else if keys[sdl.SCANCODE_W] == 1 {
		//move up
		p.y -= PlayerSpeed
	} else if keys[sdl.SCANCODE_S] == 1 {
		//move down
		p.y += PlayerSpeed
	}
}
