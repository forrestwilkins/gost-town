package main

import (
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const WorldSize = 400

type World struct {
	window fyne.Window
	canvas fyne.CanvasObject

	pixels []Pixel
}

type Pixel struct {
	x, y, w, h int
	color      color.Color
}

func (world *World) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	world.canvas.Resize(size)
}

func (world *World) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(WorldSize, WorldSize)
}

func (world *World) refresh() {
	world.window.Canvas().Refresh(world.canvas)
}

func (world *World) draw(px, py, w, h int) color.Color {
	for p := 0; p < len(world.pixels); p++ {
		pixel := world.pixels[p]

		if px >= pixel.x && px <= pixel.x+pixel.w && py >= pixel.y && py <= pixel.y+pixel.h {
			return pixel.color
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func setup(window fyne.Window) (*World, fyne.CanvasObject) {
	pixelSize := 5
	world := &World{window: window, pixels: []Pixel{
		{x: 400, y: 250, color: color.RGBA{255, 0, 0, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 350, color: color.RGBA{0, 255, 0, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 450, color: color.RGBA{0, 0, 255, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 550, color: color.RGBA{255, 255, 255, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 250, color: color.RGBA{255, 0, 0, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 350, color: color.RGBA{0, 255, 0, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 450, color: color.RGBA{0, 0, 255, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 550, color: color.RGBA{255, 255, 255, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 250, color: color.RGBA{255, 0, 0, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 350, color: color.RGBA{0, 255, 0, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 450, color: color.RGBA{0, 0, 255, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 550, color: color.RGBA{255, 255, 255, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 250, color: color.RGBA{255, 0, 0, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 350, color: color.RGBA{0, 255, 0, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 450, color: color.RGBA{0, 0, 255, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 550, color: color.RGBA{255, 255, 255, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 250, color: color.RGBA{255, 0, 0, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 350, color: color.RGBA{0, 255, 0, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 450, color: color.RGBA{0, 0, 255, 255}, w: pixelSize, h: pixelSize},
		{x: 400, y: 550, color: color.RGBA{255, 255, 255, 255}, w: pixelSize, h: pixelSize},
	}}

	raster := canvas.NewRasterWithPixels(world.draw)
	world.canvas = raster

	return world, container.New(world, world.canvas)
}

func draw(window fyne.Window, world *World) {
	for {
		time.Sleep(time.Millisecond * 10)

		for p := 0; p < len(world.pixels); p++ {
			world.pixels[p].x += walk()
			world.pixels[p].y += walk()

			raster := canvas.NewRasterWithPixels(world.draw)
			world.canvas = raster
			obj := container.New(world, world.canvas)
			window.SetContent(obj)
		}
	}
}

func walk() int {
	var (
		min = -1
		max = 1
	)
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

func main() {
	_app := app.New()
	window := _app.NewWindow("Gost Town")
	window.CenterOnScreen()

	world, worldContainer := setup(window)

	go func() {
		draw(window, world)
	}()

	window.SetContent(worldContainer)

	window.Resize(fyne.NewSize(WorldSize, WorldSize))
	window.ShowAndRun()
}
