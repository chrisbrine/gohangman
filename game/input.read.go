package game

import (
	"fmt"
	"os"
	"github.com/eiannone/keyboard"
)

func readChar() string {
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	return string(char)
}

func (g *Game) AnyKeyToContinue() {
	fmt.Print("Press any key to continue...")
	readChar()
}

func (g *Game) ReadLetter() string {
	fmt.Print("Enter a letter (or use '.' to quit): ")
	for {
		guess := readChar()
		if guess == "." {
			fmt.Println("Bye!")
			os.Exit(0)
		}
		// confirm guess is a letter
		if (guess >= "a" && guess <= "z") || (guess >= "A" && guess <= "Z") {
			return guess
		}
	}
}

func (g *Game) ReadDifficulty() int {
	fmt.Print("\033[H\033[2J")
	g.Title()
	fmt.Println("Difficulty options:")
	fmt.Println("  e - Easy, 12 guesses")
	fmt.Println("  m - Medium, 9 guesses")
	fmt.Println("  h - Hard, 6 guesses")
	fmt.Println("  x - eXtreme, 3 guesses")
	fmt.Println("  . - Quit")
	fmt.Print("Enter difficulty (e, m, h, x, .): ")
	for {
		difficulty := readChar()
		if difficulty == "." {
			fmt.Println("Bye!")
			os.Exit(0)
		}
		switch difficulty {
			case "e":
				return 12
			case "m":
				return 9
			case "h":
				return 6
			case "x":
				return 3
		}
	}
}
