package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/michaelmcallister/sierpinski/pkg/sierpinski"
)

const ScreenWidth, ScreenHeight = 640, 480

func main() {
	ebiten.SetWindowTitle("Sierpinski")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	s := sierpinski.NewGame(ScreenWidth, ScreenHeight)
	if err := ebiten.RunGame(s); err != nil {
		log.Fatal(err)
	}
}
