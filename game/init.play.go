package game

import (
	"fmt"
)

func (g *Game) isAlreadyGuessed(guess string) bool {
	for _, letter := range g.guesses {
		if letter == guess {
			return true
		}
	}
	return false
}

func (g *Game) Play() {
	for {
		fmt.Print("\033[H\033[2J")
		g.Title()
		g.ShowHangman()
		g.ShowWord()
		g.ShowGuesses()
		if g.CheckDone() {
			g.AnyKeyToContinue()
			break
		}
		guess := g.ReadLetter()
		fmt.Println(guess);
		if g.isAlreadyGuessed(guess) {
			fmt.Printf("You already guessed %s\n", guess)
			continue
		}
		g.guesses, g.wrongGuesses = g.ProcessGuess(guess)
	}
}
