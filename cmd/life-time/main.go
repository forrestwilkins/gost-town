package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/forrestwilkins/gost-town/internal/nomads"
	"github.com/hajimehoshi/ebiten"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	world := nomads.Setup()

	if err := ebiten.RunGame(world); err != nil {
		log.Fatal(err)
	}
}
