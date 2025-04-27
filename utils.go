package main

import rl "github.com/gen2brain/raylib-go/raylib"

func showPlayerName(position rl.Vector2, name string, xOff int32, yOff int32) {
	rl.DrawText(name, int32(position.X)+xOff, int32(position.Y)+yOff, 10, rl.LightGray)
}
