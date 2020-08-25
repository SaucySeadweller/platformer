package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

//Player is a player struct
type Player struct {
	tex  *sdl.Texture
	x, y float64
}

const PlayerSpeed = 0.08

func newPlayer(renderer *sdl.Renderer) (p Player, err error) {

	playerImg, err := img.Load("/home/ferro/Downloads/GraveRobber.png")
	if err != nil {
		fmt.Println(err)
		return Player{}, fmt.Errorf("loading player object: %v", err)
	}

	defer playerImg.Free()

	p.tex, err = renderer.CreateTextureFromSurface(playerImg)
	if err != nil {
		fmt.Println(err)
		return Player{}, fmt.Errorf("creating player texture: %v", err)
	}

	return
}

func (p *Player) draw(renderer *sdl.Renderer) {
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 48, H: 48},
		&sdl.Rect{X: int32(p.x), Y: int32(p.y), W: 48, H: 48})
	renderer.Present()
}
