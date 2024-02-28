package game

import (
	"fmt"
)

func (g *Game) CheckDone() bool {
	if (len(g.guesses) - g.wrongGuesses) == g.neededGuesses {
		fmt.Println("You win!")
		return true
	}
	if g.wrongGuesses >= g.maxGuesses {
		fmt.Println("You lose!")
		return true
	}
	return false
}