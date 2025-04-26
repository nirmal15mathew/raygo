package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	position rl.Vector2
	velocity rl.Vector2
	acc      rl.Vector2
	maxVel   float32
	color    rl.Color
}

func CreatePlayer() *Player {
	s := Player{
		position: rl.Vector2{X: 0, Y: 0},
		velocity: rl.Vector2{X: 0, Y: 0},
		acc:      rl.Vector2{X: 0, Y: 0},
		maxVel:   15.0,
	}

	return &s
}

func CreatePlayerWithPosition(x float32, y float32) *Player {
	s := Player{
		position: rl.Vector2{X: x, Y: y},
		velocity: rl.Vector2{X: 0, Y: 0},
		acc:      rl.Vector2{X: 0, Y: 0},
		maxVel:   10.0,
		color:    rl.Blue,
	}

	return &s
}

func updatePlayer(p *Player) {
	p.velocity = rl.Vector2Add(p.acc, p.velocity)
	p.velocity = rl.Vector2ClampValue(p.velocity, 0, p.maxVel)
	p.position = rl.Vector2Add(p.velocity, p.position)

	p.acc = rl.Vector2Scale(p.acc, 0)
}

func showPlayer(p Player) {
	rl.DrawRectangle(int32(p.position.X), int32(p.position.Y), 20, 20, p.color)
}

func applyForce(p *Player, force rl.Vector2) rl.Vector2 {
	p.acc = rl.Vector2Add(p.acc, force)
	return p.acc
}
