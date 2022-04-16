package nomads

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const (
	ScreenWidth, ScreenHeight = 600, 600
	StartingPopulation        = 1000
	RateOfGeneration          = 5
	CellSize                  = 3
)

type World struct {
	cells      map[Position]Cell
	generation int
}

type Cell struct {
	color color.Color
}

type Position struct {
	x, y int
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
				case (count == 2 || count == 3) && cell.isAlive():
					cell.color = color.RGBA{0, 50, 255, 255}
					nextGeneration[cellPosition] = cell

				// Rule 4:
				case count == 3:
					cell.color = color.RGBA{0, 255, 0, 255}
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

func Setup() *World {
	ebiten.SetWindowTitle("Life")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	world := &World{cells: make(map[Position]Cell)}

	for gridY := ScreenHeight / 4; gridY <= ScreenHeight*0.75; gridY += (CellSize * RandomWithin(0, 6)) {
		for gridX := ScreenWidth / 4; gridX <= ScreenWidth*0.75; gridX += (CellSize * RandomWithin(0, 6)) {
			if len(world.cells) >= StartingPopulation {
				break
			}

			cellPosition := Position{gridX, gridY}
			newCell := Cell{color: color.White}

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
			if neighborX < 0 || neighborY < 0 || neighborX >= ScreenWidth ||
				neighborY >= ScreenHeight {
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

func (world *World) DrawGlider(
	glider Spaceship,
	startingX int,
	startingY int,
	color color.RGBA,
) *World {
	newCell := Cell{color}

	for y, row := range glider {
		for x, cell := range row {
			if cell == 1 {
				posX := startingX + x*CellSize
				posY := startingY + y*CellSize
				world.cells[Position{posX, posY}] = newCell
			}
		}
	}

	return world
}
