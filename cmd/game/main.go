// TODO: This should replace life/main.go once able to to animate pixels

package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const Width, Height = 400, 500

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

	if !ebiten.IsDrawingSkipped() {
		world.update()
		world.render(screen)
	}

	return err
}

func (world *World) update() {
	for p := 0; p < len(world.pixels); p++ {
		world.pixels[p].walk()
	}
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
		{x: 100, y: 100, color: color.RGBA{255, 0, 0, 255}, w: pixelSize, h: pixelSize},
		{x: 100, y: 200, color: color.RGBA{0, 255, 0, 255}, w: pixelSize, h: pixelSize},
		{x: 100, y: 300, color: color.RGBA{0, 0, 255, 255}, w: pixelSize, h: pixelSize},
		{x: 100, y: 400, color: color.RGBA{255, 255, 255, 255}, w: pixelSize, h: pixelSize},
	}}

	return world
}

func main() {
	world := setup()

	if err := ebiten.Run(world.frame, Width, Height, 2, "Gost Town"); err != nil {
		log.Fatal(err)
	}
}
