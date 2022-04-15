package main

import (
	"image/color"
	"math"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const WindowSize = 800

type fractal struct {
	currIterations          uint
	currScale, currX, currY float64

	window fyne.Window
	canvas fyne.CanvasObject
}

func (f *fractal) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	f.canvas.Resize(size)
}

func (f *fractal) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(WindowSize, WindowSize)
}

//lint:ignore U1000  See TODO inside the .Show() method.
func (f *fractal) refresh() {
	if f.currScale >= 1.0 {
		f.currIterations = 100
	} else {
		f.currIterations = uint(100 * (1 + math.Pow((math.Log10(1/f.currScale)), 1.25)))
	}

	f.window.Canvas().Refresh(f.canvas)
}

func (f *fractal) scaleChannel(c float64, start, end uint32) uint8 {
	if end >= start {
		return (uint8)(c*float64(uint8(end-start))) + uint8(start)
	}

	return (uint8)((1-c)*float64(uint8(start-end))) + uint8(end)
}

func (f *fractal) scaleColor(c float64, start, end color.Color) color.Color {
	r1, g1, b1, _ := start.RGBA()
	r2, g2, b2, _ := end.RGBA()
	return color.RGBA{
		f.scaleChannel(c, r1, r2),
		f.scaleChannel(c, g1, g2),
		f.scaleChannel(c, b1, b2),
		0xff,
	}
}

func (f *fractal) mandelbrot(px, py, w, h int) color.Color {
	backgroundColor := color.RGBA{0, 0, 0, 255}
	edgeColor := color.RGBA{
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		uint8(rand.Intn(255)),
		255,
	}
	innerColor := color.RGBA{0, 0, 0, 255}

	drawScale := 3.5 * f.currScale
	aspect := (float64(h) / float64(w))
	cRe := ((float64(px)/float64(w))-0.5)*drawScale + f.currX
	cIm := ((float64(py)/float64(w))-(0.5*aspect))*drawScale - f.currY

	var i uint
	var x, y, xsq, ysq float64

	for i = 0; i < f.currIterations && (xsq+ysq <= 4); i++ {
		xNew := float64(xsq-ysq) + cRe
		y = 2*x*y + cIm
		x = xNew

		xsq = x * x
		ysq = y * y
	}

	if i == f.currIterations {
		return innerColor
	}

	mu := (float64(i) / float64(f.currIterations))
	c := math.Sin((mu / 2) * math.Pi)

	return f.scaleColor(c, backgroundColor, edgeColor)
}

//lint:ignore U1000 See TODO inside the .Show() method.
func (f *fractal) fractalRune(r rune) {
	if r == '+' || r == '=' {
		f.currScale /= 1.1
	} else if r == '-' || r == '_' {
		f.currScale *= 1.1
	}

	f.refresh()
}

//lint:ignore U1000 See TODO inside the .Show() method.
func (f *fractal) fractalKey(ev *fyne.KeyEvent) {
	delta := f.currScale * 0.2
	if ev.Name == fyne.KeyDown {
		f.currY -= delta
	} else if ev.Name == fyne.KeyUp {
		f.currY += delta
	} else if ev.Name == fyne.KeyRight {
		f.currX += delta
	} else if ev.Name == fyne.KeyLeft {
		f.currX -= delta
	}

	f.refresh()
}

// Show loads a Mandelbrot fractal example window for the specified app context
func Show(win fyne.Window, currIterations uint) fyne.CanvasObject {
	fractal := &fractal{window: win}
	fractal.canvas = canvas.NewRasterWithPixels(fractal.mandelbrot)

	fractal.currIterations = currIterations
	fractal.currScale = 1.0
	fractal.currX = -0.75
	fractal.currY = 0.0

	// TODO: Register, and unregister, these keys:
	win.Canvas().SetOnTypedKey(fractal.fractalKey)
	win.Canvas().SetOnTypedRune(fractal.fractalRune)

	return container.New(fractal, fractal.canvas)
}

func main() {
	a := app.New()
	w := a.NewWindow("Mandelbrot")
	w.CenterOnScreen()

	mandelbrot := Show(w, 500)
	w.SetContent(mandelbrot)

	// go func() {
	// 	for i := 1; i < 10000; i++ {
	// 		time.Sleep(time.Millisecond * 100)

	// 		obj1 := Show(w, uint(i))
	// 		w.SetContent(container.NewWithoutLayout(obj1))
	// 	}
	// }()

	w.Resize(fyne.NewSize(WindowSize, WindowSize))
	w.ShowAndRun()
}
