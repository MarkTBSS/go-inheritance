package warrior

import (
	"fmt"

	"github.com/MarkTBSS/go-inheritance/novice"
)

// Warrior struct embeds the Novice struct and adds a color field
type Warrior struct {
	*novice.Novice // Embedding the Novice struct
	color          string
}

// NewWarrior creates a new instance of Warrior
func NewWarrior(name, classname string, health int, color string) *Warrior {
	return &Warrior{
		Novice: novice.NewNovice(name, classname, health), // Creating a new Novice instance
		color:  color,
	}
}

// DisplayInfo displays the information of the Warrior including the color
func (w *Warrior) DisplayInfo() {
	w.Novice.DisplayInfo() // Calling DisplayInfo method of the embedded Novice struct
	fmt.Println("Color Changed:", w.color)
}

// DeleteHealth decreases the player's health by half the given damage
func (w *Warrior) DeleteHealth(damage int) {
	halfDamage := damage / 2
	w.Novice.DeleteHealth(halfDamage)
}
