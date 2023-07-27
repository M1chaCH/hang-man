package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var PHRASES = [6]string{"programming is fun", "hangman is a game", "happy coding", "keep learning", "go lang is awesome", "openai is amazing"}

const LIVES = 10
const BLANK_ASCII_CODE = 32
const HIDDEN_ASCII_CODE = 95

func main() {
	fmt.Println("Welcome to Hangman!")
	playNewGame()
}

func playNewGame() {
	var currentPhrase []rune
	var currentId = -1
	currentPhrase, currentId = findNewWord(currentId)
	currentGuess := generateGuessArray(currentPhrase)
	currentLives := 0
	defer fmt.Println("See you 🙋‍♂️")

	fmt.Println("You can either guess a letter of the phrase or the entire phrase")
	fmt.Println("If you miss a letter you fill lose 1 live if you miss the phrase you will lose 4 lives!")
	fmt.Println("I have selected a word, you can start guessing")

	won := false
	playing := true
	for playing {
		renderGame(currentLives, currentGuess)
		fmt.Println("Guess a letter or the phrase (press enter to send)")
		input := readInputLine()

		if len(input) == 0 || input == "error" {
			fmt.Println("invalid input")
			continue
		} else if len(input) == 1 {
			// guess one letter
			inputLetter := []rune(input)[0]
			changedOne := false
			for i, r := range currentPhrase {
				if r == inputLetter {
					currentGuess[i] = inputLetter
					changedOne = true
				}
			}

			if changedOne {
				won = checkIfWon(currentGuess, currentPhrase)
			} else {
				currentLives++
			}
		} else {
			inputPhrase := []rune(input)
			won = checkIfWon(inputPhrase, currentPhrase)
			if !won {
				currentLives += 4
			}
		}
		if won {
			fmt.Println("You won!")
			if checkIfReplay() {
				playNewGame()
			} else {
				playing = false
			}
		} else if currentLives >= LIVES {
			fmt.Println("You lost ):")
			if checkIfReplay() {
				playNewGame()
			} else {
				playing = false
			}
		}
	}
}

func findNewWord(currentIndex int) (phrase []rune, index int) {
	for {
		index = rand.Intn(len(PHRASES))
		if index != currentIndex {
			phrase = []rune(PHRASES[index])
			return
		}
	}
}

func generateGuessArray(currentPhrase []rune) (startGuess []rune) {
	startGuess = make([]rune, len(currentPhrase))
	for i, r := range currentPhrase {
		if r == BLANK_ASCII_CODE {
			startGuess[i] = r
		} else {
			startGuess[i] = HIDDEN_ASCII_CODE
		}
	}
	return
}

func readInputLine() string {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("$: ")
	if input, err := inputReader.ReadString('\n'); err == nil {
		input = strings.ReplaceAll(input, "\n", "")
		return strings.ToLower(input)
	} else {
		fmt.Println("unexpected error", err)
		return "error"
	}
}

func renderGame(currentLives int, currentGuess []rune) {
	fmt.Printf("\n     %v/%v   [ ", currentLives, LIVES)
	for i := 0; i < LIVES; i++ {
		if i < currentLives {
			fmt.Print("█ ")
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Print("] Lives\n\n     Phrase ")

	for _, char := range currentGuess {
		fmt.Printf("%c ", char)
	}
	fmt.Print("\n")
}

func checkIfReplay() bool {
	fmt.Println("Do you want to play again? (Yes, Y)")
	input := readInputLine()
	if lowerInput := strings.ToLower(input); lowerInput == "y" || lowerInput == "yes" {
		fmt.Println("Great! Lets start again.")
		return true
	} else {
		return false
	}
}

func checkIfWon(guess []rune, phrase []rune) bool {
	correct := true
	for i, r := range phrase {
		if r != guess[i] {
			correct = false
			break
		}
	}
	return correct
}
