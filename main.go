package main

import (
	"go-hangman/game"
)

func main() {
	var Game game.Game
	for {
		Game = game.NewGame()
		Game.Play()
	}
}