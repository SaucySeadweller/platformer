package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

//Player is a player struct
type Player struct {
	tex *sdl.Texture
}

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
