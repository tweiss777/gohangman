package main

import (
	"bufio"
	"fmt"
	hangman "hangman/game_functions"
	"os"
)

var unSolved bool = false

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("--- WELCOME TO HANGMAN ---")
	fmt.Println("--- ENTER A WORD ---")
	text, _ := reader.ReadString('\n')
	hangman.Setup(&text)
	for !unSolved {
		fmt.Println("Enter a letter")
		text, _ := reader.ReadString('\n')

		if !hangman.InputValid(&text) {
			fmt.Println("invalid input")
			continue
		}
		var found, occurences = hangman.CheckIfExists(text)

		if found {
			result := fmt.Sprintf("found %d occurences of %s", occurences, text)
			fmt.Println(result)

		} else {
			fmt.Printf("Oops, the letter %s doesn't exist here", text)
			hangman.IncrememntGuesses()

		}
		hangman.PrintStats()
		// here check if the game is still one
		var gameOver bool = hangman.CheckIfGameOver()
		if gameOver {
			unSolved = true
		}
	}
	fmt.Println("Game over")

}
