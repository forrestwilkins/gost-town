package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	ScreenWidth, ScreenHeight = 600, 600
	StartingPopulation        = 5000
	RateOfGeneration          = 5
	CellSize                  = 2
)

type World struct {
	generation int
	cells      map[Position]Cell
}

type Cell struct {
	color color.Color
}

type Position struct {
	x, y int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (world *World) Draw(screen *ebiten.Image) {
	for position, cell := range world.cells {
		for y := position.y; y < position.y+CellSize; y++ {
			for x := position.x; x < position.x+CellSize; x++ {
				screen.Set(x, y, cell.color)
			}
		}
	}
}

func (world *World) Update(screen *ebiten.Image) error {
	nextGeneration := make(map[Position]Cell)

	if world.generation%RateOfGeneration == 0 {
		for gridY := 0; gridY <= ScreenHeight; gridY += CellSize {
			for gridX := 0; gridX <= ScreenWidth; gridX += CellSize {
				cellPosition := Position{gridX, gridY}
				cell := world.cells[cellPosition]
				count := world.neighborCount(gridX, gridY)

				switch {
				// Rule 1 and 3:
				case count < 2 || count > 3:
					cell.color = color.Black
					nextGeneration[cellPosition] = cell

				// Rule 2:
				case (count == 2 || count == 3) && !(cell.color == color.Black || cell.color == nil):
					cell.color = color.RGBA{0, 255, 0, 255}
					nextGeneration[cellPosition] = cell

				// Rule 4:
				case count == 3:
					cell.color = color.RGBA{0, 0, 255, 255}
					nextGeneration[cellPosition] = cell
				}
			}
		}
		world.cells = nextGeneration
	}
	world.generation++

	return nil
}

func (g *World) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func setup() *World {
	ebiten.SetWindowTitle("Life")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	world := &World{cells: make(map[Position]Cell)}

	for gridY := ScreenHeight / 4; gridY <= ScreenHeight*0.75; gridY += (CellSize * randomWithin(0, 6)) {
		for gridX := ScreenWidth / 4; gridX <= ScreenWidth*0.75; gridX += (CellSize * randomWithin(0, 6)) {
			if len(world.cells) >= StartingPopulation {
				break
			}

			cellPosition := Position{gridX, gridY}

			newCell := Cell{color: randomRGBA("gray")}
			world.cells[cellPosition] = newCell
		}
	}

	return world
}

func (world *World) neighborCount(cellX int, cellY int) int {
	count := 0

	for gridY := -CellSize; gridY <= CellSize; gridY += CellSize {
		for gridX := -CellSize; gridX <= CellSize; gridX += CellSize {
			if gridX == 0 && gridY == 0 {
				continue
			}
			neighborX := cellX + gridX
			neighborY := cellY + gridY
			if neighborX < 0 || neighborY < 0 || neighborX >= ScreenWidth || neighborY >= ScreenHeight {
				continue
			}

			neighboringCell := world.cells[Position{neighborX, neighborY}]
			if neighboringCell.isAlive() {
				count++
			}
		}
	}

	return count
}

func (cell *Cell) isAlive() bool {
	return cell.color != nil && cell.color != color.Black
}

func randomRGBA(colors ...string) color.RGBA {
	randomRed := uint8(randomWithin(0, 255))
	randomGreen := uint8(randomWithin(0, 255))
	randomBlue := uint8(randomWithin(0, 255))

	if colors != nil {
		switch colors[0] {
		case "red":
			return color.RGBA{randomRed, 0, 0, 255}
		case "green":
			return color.RGBA{0, randomGreen, 0, 255}
		case "blue":
			return color.RGBA{0, 0, randomBlue, 255}
		case "gray":
			return color.RGBA{randomRed, randomRed, randomRed, 255}
		}
	}

	return color.RGBA{randomRed, randomGreen, randomBlue, 255}
}

func RandomDivisibleBy(min int, max int, divisibleBy int) int {
	random := randomWithin(min, max)
	remainder := random % divisibleBy

	if random > 0 {
		return random - remainder
	}
	return random + remainder
}

func randomWithin(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func main() {
	world := setup()

	if err := ebiten.RunGame(world); err != nil {
		log.Fatal(err)
	}
}
