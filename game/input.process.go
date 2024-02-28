package game

func (g *Game) isCorrectGuess(guess string) bool {
	for _, letter := range g.word {
		if string(letter) == guess {
			return true
		}
	}
	return false
}

func (g *Game) ProcessGuess(guess string) ([]string, int) {
	g.guesses = append(g.guesses, guess)
	if !g.isCorrectGuess(guess) {
		g.wrongGuesses++
	}
	return g.guesses, g.wrongGuesses
}