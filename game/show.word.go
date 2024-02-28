package game

import (
	"fmt"
)

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func (g *Game) ShowWord() {
	var display string
	for _, letter := range g.word {
		if contains(g.guesses, string(letter)) {
			display += string(letter)
		} else {
			display += "_"
		}
	}
	fmt.Println(display)
}
