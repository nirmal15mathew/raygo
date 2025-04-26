package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Particle struct {
	position rl.Vector2
	vel      rl.Vector2
	color    rl.Color
	exists   bool
}

type PSystem struct {
	particles []Particle
}

func CreateParticleSystem(particleCount uint16) *PSystem {
	p := PSystem{}
	return &p
}

func updateSystem(system *PSystem) {

}

func renderSystem(system *PSystem) {

}
