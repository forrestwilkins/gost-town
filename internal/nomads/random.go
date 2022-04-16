package nomads

import (
	"image/color"
	"math/rand"
)

func RandomRGBA(colors ...string) color.RGBA {
	randomRed := uint8(RandomWithin(0, 255))
	randomGreen := uint8(RandomWithin(0, 255))
	randomBlue := uint8(RandomWithin(0, 255))

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

func RandomWithin(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomDivisibleBy(min int, max int, divisibleBy int) int {
	random := RandomWithin(min, max)
	remainder := random % divisibleBy

	if random > 0 {
		return random - remainder
	}
	return random + remainder
}
