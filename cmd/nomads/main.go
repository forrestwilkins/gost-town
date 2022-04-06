package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const ScreenWidth, ScreenHeight = 500, 500

type World struct {
	generation int
	nomads     []Nomad
}

type Nomad struct {
	w, h     int
	position Position
	color    color.Color
	colorMap map[Position]color.Color
}

type Position struct {
	x, y int
}

func (world *World) Update(screen *ebiten.Image) error {
	if world.generation%4 == 0 {
		for p := 0; p < len(world.nomads); p++ {
			world.nomads[p].walk()
		}
	}

	world.generation++

	return nil
}

func (world *World) Draw(screen *ebiten.Image) {
	for x := 0; x < ScreenWidth; x++ {
		for y := 0; y < ScreenHeight; y++ {
			for p := 0; p < len(world.nomads); p++ {
				nomad := world.nomads[p]
				position := nomad.position

				if x >= position.x && x <= position.x+nomad.w && y >= position.y && y <= position.y+nomad.h {
					screen.Set(x, y, nomad.color)
					world.nomads[p].colorMap[position] = nomad.color
				} else if val, ok := world.nomads[p].colorMap[Position{x: x, y: y}]; ok {
					screen.Set(x, y, val)
				}
			}
		}
	}
}

func (g *World) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (nomad *Nomad) walk() {
	speed := 1
	nomad.position.x += random(-speed, speed)
	nomad.position.y += random(-speed, speed)
}

func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

func setup() *World {
	ebiten.SetWindowTitle("Nomads")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	nomadSize := 2

	world := &World{nomads: []Nomad{
		{position: Position{x: ScreenWidth * 0.5, y: ScreenHeight * 0.2}, color: color.RGBA{255, 0, 0, 255}, w: nomadSize, h: nomadSize, colorMap: make(map[Position]color.Color)},
		{position: Position{x: ScreenWidth * 0.5, y: ScreenHeight * 0.4}, color: color.RGBA{0, 255, 0, 255}, w: nomadSize, h: nomadSize, colorMap: make(map[Position]color.Color)},
		{position: Position{x: ScreenWidth * 0.5, y: ScreenHeight * 0.6}, color: color.RGBA{0, 0, 255, 255}, w: nomadSize, h: nomadSize, colorMap: make(map[Position]color.Color)},
		{position: Position{x: ScreenWidth * 0.5, y: ScreenHeight * 0.8}, color: color.RGBA{255, 255, 255, 255}, w: nomadSize, h: nomadSize, colorMap: make(map[Position]color.Color)},
	}}

	return world
}

func main() {
	world := setup()

	if err := ebiten.RunGame(world); err != nil {
		log.Fatal(err)
	}
}
