package main

import (
	"github.com/MarkTBSS/go-inheritance/novice"
	"github.com/MarkTBSS/go-inheritance/warrior"
)

func main() {
	player1 := novice.NewNovice("John", "Novice", 100)
	player1.DisplayInfo()
	player1.DeleteHealth(40)
	player1.DisplayInfo()

	player2 := warrior.NewWarrior("Aragon", "Warrior", 100, "Red")
	player2.DisplayInfo()
	player2.DeleteHealth(40)
	player2.DisplayInfo()
}
