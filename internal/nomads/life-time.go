package nomads

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	ScreenWidth, ScreenHeight = 600, 600
	StartingPopulation        = 8000
	RateOfGeneration          = 5
	CellSize                  = 5
)

type Grid struct {
	cells      map[Position]Cell
	generation int
}

type Cell struct {
	color color.RGBA
}

type Position struct {
	x, y int
}

func (grid *Grid) Draw(screen *ebiten.Image) {
	for position, cell := range grid.cells {
		for y := position.y; y < position.y+CellSize; y++ {
			for x := position.x; x < position.x+CellSize; x++ {
				screen.Set(x, y, cell.color)
			}
		}
	}
}

func (grid *Grid) Update(screen *ebiten.Image) error {
	nextGeneration := make(map[Position]Cell)

	if grid.generation%RateOfGeneration == 0 {
		for gridY := 0; gridY < ScreenHeight; gridY += CellSize {
			for gridX := 0; gridX < ScreenWidth; gridX += CellSize {
				cellPosition := Position{gridX, gridY}
				cell := grid.cells[cellPosition]
				count := len(grid.getLiveNeighbors(gridX, gridY))

				// TODO: Optimize GetNeighborhoodDiversity - is currently very slow
				// diversity := grid.GetNeighborhoodDiversity(gridX, gridY)

				switch {
				// Rule 1 and 3:
				case count < 2 || count > 3:
					cell.color = color.RGBA{0, 0, 0, 255}
					nextGeneration[cellPosition] = cell

				// Rule 2:
				case (count == 2 || count == 3) && cell.isAlive():
					nextGeneration[cellPosition] = cell

				// Rule 4:
				case count == 3: // && diversity > 0:
					cell.color = grid.getLiveNeighborColorAverage(gridX, gridY)
					nextGeneration[cellPosition] = cell
				}
			}
		}
		grid.cells = nextGeneration
	}
	grid.generation++

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		grid.cells = initializeCells()
	}

	return nil
}

func (g *Grid) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func Setup() *Grid {
	ebiten.SetWindowTitle("Life")
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	grid := &Grid{cells: initializeCells()}

	return grid
}

func initializeCells() map[Position]Cell {
	cells := make(map[Position]Cell)

	for gridY := ScreenHeight / 4; gridY <= ScreenHeight*0.75; gridY += CellSize {
		for gridX := ScreenWidth / 4; gridX <= ScreenWidth*0.75; gridX += CellSize {
			if len(cells) >= StartingPopulation {
				break
			}

			cellPosition := Position{gridX, gridY}
			newCell := Cell{color: RandomRGBA()}

			cells[cellPosition] = newCell
		}
	}

	return cells
}

func (grid *Grid) getLiveNeighborColorAverage(
	cellX int,
	cellY int,
) color.RGBA {
	liveNeighbors := grid.getLiveNeighbors(cellX, cellY)
	redSum, greenSum, blueSum := 0, 0, 0

	for _, cell := range liveNeighbors {
		redSum += int(cell.color.R)
		greenSum += int(cell.color.G)
		blueSum += int(cell.color.B)
	}

	redAverage := redSum / len(liveNeighbors)
	greenAverage := greenSum / len(liveNeighbors)
	blueAverage := blueSum / len(liveNeighbors)

	return color.RGBA{
		uint8(redAverage),
		uint8(greenAverage),
		uint8(blueAverage),
		255,
	}
}

// TODO: Need to thoroughly test GetNeighborhoodDiversity
func (grid *Grid) GetNeighborhoodDiversity(cellX int, cellY int) float64 {
	liveNeighbors := grid.getLiveNeighbors(cellX, cellY)

	var differenceSum float64 = 0
	for _, cell1 := range liveNeighbors {
		for _, cell2 := range liveNeighbors {
			differenceSum += GetColorDifference(cell1.color, cell2.color)
		}
	}

	return differenceSum / Squared(float64(len(liveNeighbors)))
}

func (grid *Grid) getLiveNeighbors(cellX int, cellY int) map[Position]Cell {
	neighborhood := grid.getNeighborhood(cellX, cellY)
	liveNeighbors := make(map[Position]Cell)

	for position, cell := range neighborhood {
		if cell.isAlive() {
			liveNeighbors[position] = cell
		}
	}

	return liveNeighbors
}

func (grid *Grid) getNeighborhood(cellX int, cellY int) map[Position]Cell {
	neighborhood := make(map[Position]Cell)

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

			position := Position{neighborX, neighborY}
			neighboringCell := grid.cells[position]
			neighborhood[position] = neighboringCell
		}
	}

	return neighborhood
}

func (cell *Cell) isAlive() bool {
	return cell.color.R+cell.color.G+cell.color.B > 0
}

func (grid *Grid) DrawGlider(
	glider Spaceship,
	startingX int,
	startingY int,
	color color.RGBA,
) *Grid {
	newCell := Cell{color}

	for y, row := range glider {
		for x, cell := range row {
			if cell == 1 {
				posX := startingX + x*CellSize
				posY := startingY + y*CellSize
				grid.cells[Position{posX, posY}] = newCell
			}
		}
	}

	return grid
}
