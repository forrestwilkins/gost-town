package nomads

import (
	"math"
	"math/rand"
)

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

func Squared(number float64) float64 {
	return math.Pow(number, 2)
}
