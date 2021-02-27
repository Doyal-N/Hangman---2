package game

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

const errorsAllowed int = 7

var userGuesses []string
var letters []string = getRandomWord("./data/words.txt")
var normalLetters []string = normalizedLetters(letters, normalizeLetter)

func difference(a, b []string) (diff []string) {
	m := make(map[string]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}

//GetErrors return errors
func GetErrors() []string {
	return difference(userGuesses, normalLetters)
}

//ErrorsMade return number done errors
func ErrorsMade() int {
	return len(GetErrors())
}

//ErrorsBalance return rest of attempts
func ErrorsBalance() int {
	return errorsAllowed - ErrorsMade()
}

func getRandomWord(path string) []string {
	rand.Seed(time.Now().UnixNano())

	data, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	}

	allWords := strings.Fields(string(data))

	word := allWords[rand.Intn(len(allWords))]

	return strings.Split(string(word), "")
}

func normalizeLetter(letter string) string {
	if letter == "Ё" {
		return "Е"
	} else if letter == "Й" {
		return "И"
	}
	return letter
}

func normalizedLetters(arrLetters []string, f func(string) string) []string {
	normLetters := make([]string, len(arrLetters))

	for i, val := range arrLetters {
		normLetters[i] = f(val)
	}
	return normLetters
}

//Result return word
func Result() string {
	return strings.Join(lettersToGuess(), "")
}

//LettersToGuess return guess letter or __
func lettersToGuess() []string {
	result := make([]string, utf8.RuneCountInString(Word()))

	for _, letter := range letters {
		if isLetterContains(userGuesses, letter) {
			result = append(result, letter)
		} else {
			result = append(result, ".__.")
		}
	}
	return result
}

//IsUserLost return true if attempt null
func IsUserLost() bool {
	if num := ErrorsBalance(); num == 0 {
		return true
	}
	return false
}

//IsUserWon return true if user won
func IsUserWon() bool {
	if len(difference(normalLetters, userGuesses)) == 0 {
		return true
	}
	return false
}

//IsGameOver return false while not game over
func IsGameOver() bool {
	if IsUserLost() || IsUserWon() {
		return true
	}
	return false
}

//Word return entirely
func Word() string {
	return strings.Join(letters, "")
}

func isLetterContains(arr []string, elem string) bool {
	for _, val := range arr {
		if normalizeLetter(elem) == val {
			return true
		}
	}
	return false
}

//Play - start game
func Play(letter string) {
	if IsGameOver() == false && isLetterContains(userGuesses, letter) == false {
		normLetter := normalizeLetter(letter)
		userGuesses = append(userGuesses, normLetter)
	}
}
