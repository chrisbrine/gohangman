package game

import (
	"math/rand"
	"time"
)

var dictionary = []string{
	"chicken",
	"dog",
	"cat",
	"mouse",
	"frog",
	"parrot",
	"hamster",
	"cow",
	"sheep",
	"elephant",
}

func GetWord() string {
	rand.Seed(time.Now().UnixNano())
	word := dictionary[rand.Intn(len(dictionary))]
	return word
}
