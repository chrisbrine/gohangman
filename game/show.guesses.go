package game

import (
	"fmt"
)

func (g *Game) ShowGuesses() {
	fmt.Println("Guesses: ", g.guesses)
	fmt.Println("Wrong Guesses: ", g.wrongGuesses, "/", g.maxGuesses)
}
