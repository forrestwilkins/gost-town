// TODO: This should replace life/main.go once able to to animate pixels

package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const Width, Height = 500, 500

type World struct {
	generation int
	pixels     []Pixel
}

type Pixel struct {
	x, y, w, h int
	color      color.Color
}

func (world *World) render(screen *ebiten.Image) {
	for x := 0; x < Width; x++ {
		for y := 0; y < Height; y++ {
			for p := 0; p < len(world.pixels); p++ {
				pixel := world.pixels[p]

				if x >= pixel.x && x <= pixel.x+pixel.w && y >= pixel.y && y <= pixel.y+pixel.h {
					screen.Set(x, y, pixel.color)
				}
			}
		}
	}

}

func (world *World) frame(screen *ebiten.Image) error {
	var err error = nil

	if world.generation%5 == 0 {
		err = world.update()
	}

	if !ebiten.IsDrawingSkipped() {
		world.render(screen)
	}

	world.generation++

	return err
}

func (world *World) update() error {
	for p := 0; p < len(world.pixels); p++ {
		world.pixels[p].walk()
	}

	return nil
}

func (pixel *Pixel) walk() {
	pixel.x += random(-1, 1)
	pixel.y += random(-1, 1)
}

func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

func setup() *World {
	pixelSize := 1

	world := &World{pixels: []Pixel{
		{x: Width / 2, y: Height * 0.2, color: color.RGBA{255, 0, 0, 255}, w: pixelSize, h: pixelSize},
		{x: Width / 2, y: Height * 0.4, color: color.RGBA{0, 255, 0, 255}, w: pixelSize, h: pixelSize},
		{x: Width / 2, y: Height * 0.6, color: color.RGBA{0, 0, 255, 255}, w: pixelSize, h: pixelSize},
		{x: Width / 2, y: Height * 0.8, color: color.RGBA{255, 255, 255, 255}, w: pixelSize, h: pixelSize},
	}}

	return world
}

func main() {
	world := setup()

	if err := ebiten.Run(world.frame, Width, Height, 0.8, "Gost Town"); err != nil {
		log.Fatal(err)
	}
}
