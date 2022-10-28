package hangman

import (
	// "errors"
	"fmt"
	"strings"
)

const MaxGuesses = 10

var word *string
var wordLen int
var correctLetters []string
var occurences map[string]int

// var wordLen int
var numberOfGuesses int = 0

// all the letters in the array split into a string
var letters []string

func IncrememntGuesses() {
	numberOfGuesses++
}

// checks for underscore in text
func underScoreExists() bool {
	for _, ch := range correctLetters {
		if ch == "_" {
			return true
		}
	}
	return false
}

func CheckIfGameOver() bool {
	if numberOfGuesses < MaxGuesses && underScoreExists() {
		return false
	}
	return true

}

// do initial setup of the game
func mapSetup() {
	occurences = make(map[string]int)
	for _, letter := range letters {
		if occurences[letter] == 0 {
			occurences[letter] = 1
		} else {
			occurences[letter]++
		}
	}
	// check the map to see that it was initialized correctly
	for key, value := range occurences {
		fmt.Printf("letter: %s, occurences: %d \n", key, value)
	}
}

/*
* sets up the game
*/
func Setup(wordToStore *string) {
	word = wordToStore
	*word = strings.Replace(*word,"\n","",1)
	letters = strings.Split(*word, "")
	wordLen = len(letters)
	fmt.Println(correctLetters)
	// append blank field to the correct spot
	for i := 0; i < wordLen; i++ {
		correctLetters = append(correctLetters, "_")
	}
	fmt.Println("correct letters in array", correctLetters)
	// add letter and count into map
	mapSetup()
}

func getIndexOfChar(input string, list []string) int {
	for i, l := range list {
		if input == l {
			return i
		}
	}
	return -1
}

func CheckIfExists(input string) (bool, int) {
	//	check in the board if the input exists
	var exists = occurences[input] >= 1
	if !exists {
		return false, 0
	}
	var indexOfChar = getIndexOfChar(input, letters)
	fmt.Println(indexOfChar)
	correctLetters[indexOfChar] = input
	return true, 1
}

func InputValid(input *string) bool {
	// trim new lines
	*input = strings.Replace(*input, "\n", "", 1)
	fmt.Println(len(*input))
	// evaluate if len of new line is == to one
	if len(*input) > 1 {
		return false
	}
	return true
}

func PrintStats() {
	//Print the correct letters to the console
	fmt.Println("correct letters")
	fmt.Println(correctLetters)

	//Print the number of guesses to the console
	fmt.Println("Number of guesses remaining")
	fmt.Println(numberOfGuesses)

}
