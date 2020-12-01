package main

import (
	"log"

	"./models"
)

func main() {
	p1 := models.InitPlayer("Player 1", 0, 0, 1)
	for true {
		p1.MoveForward()
		log.Printf("Player 1 Location: [%.6f, %.6f]", p1.Loc.X, p1.Loc.Y)
	}
}
