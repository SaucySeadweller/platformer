package main

import (
	"fmt"

	//#include <SDL/SDL_image.h>

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

//go build -tags static -ldflags "-s -w"
func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL ", err)
		return
	}
	window, err := sdl.CreateWindow(
		"Platformer",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window", err)
		return
	}

	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer", err)
		return
	}
	defer renderer.Destroy()
	//	player, err := game.newPlayer(renderer)
	//img.Init(img.INIT_JPG | img.INIT_PNG)

	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}

			renderer.Clear()
			renderer.SetDrawColor(0, 100, 240, 222)
			//	renderer.Copy(playerTex, &sdl.Rect{X: 0, Y: 0, W: 48, H: 48}, &sdl.Rect{X: 100, Y: 200, W: 48, H: 48})
			renderer.Present()
		}
	}
}
