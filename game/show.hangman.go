package game

import (
	"fmt"
	"io/ioutil"
)

func (g *Game) calculateHangmanLevel() int {
	return int((float64(g.wrongGuesses) / float64(g.maxGuesses)) * 9.0)
}

func (g *Game) ShowHangman() {
	currentState := g.calculateHangmanLevel()
	currentFile := fmt.Sprintf("states/hangman%d", currentState)
	state, err := ioutil.ReadFile(currentFile)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	fmt.Println(string(state))
}
