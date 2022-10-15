package hangman

import (
	"errors"
	"fmt"
	"strings"
)

const maxGuesses = 10

var word *string
var wordLen int
var correctLetters []string
var occurences map[string]int

// var wordLen int
var numberOfGuesses int = 0

// all the letters in the array split into a string
var letters []string

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

func Setup(wordToStore *string) {
	word = wordToStore
	letters = strings.Split(*word, "")
	wordLen = len(letters)
	fmt.Println(correctLetters)
	// append blank field to the correct spot
	for i := 0; i < wordLen; i++ {
		correctLetters = append(correctLetters, "")
	}
	fmt.Println("correct letters in array", correctLetters)
	// add letter and count into map
	mapSetup()
}

func getIndexOfChar(input string, list []string) int{
	for i,l := range list{
		if input == l {
			return i
		}
	}
	return -1
}

func CheckIfExists(input string) (bool, int) {
	var exists = occurences[input] >= 1
	if !exists {
		return false, 0
	}
	var indexOfChar = getIndexOfChar(input, letters)
	fmt.Println(indexOfChar)
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

func PrintStats(){

}
