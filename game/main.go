package game

type Game struct {
	word string
	guesses []string
	wrongGuesses int
	maxGuesses int
	neededGuesses int
}

func getUniqueLetters(word string) int {
	uniqueLetters := make(map[string]bool)
	for _, letter := range word {
		uniqueLetters[string(letter)] = true
	}
	return len(uniqueLetters)
}

func NewGame() Game {
	word := GetWord()
	game := Game{
		word: word,
		guesses: []string{},
		wrongGuesses: 0,
		maxGuesses: 9,
		neededGuesses: getUniqueLetters(word),
	}
	game.maxGuesses = game.ReadDifficulty()
	return game
}
