package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"work/src/hangman/game"
)

func main() {
	printOut()

	for game.IsGameOver() == false {
		letter := userInput()
		game.Play(letter)
		printOut()
	}
}

func readFiles() []string {
	var contents []string

	for i := 0; i < 8; i++ {
		path := fmt.Sprintf("./data/figures/%d.txt", i)
		data, err := ioutil.ReadFile(path)

		if err != nil {
			fmt.Println(err)
		}

		contents = append(contents, string(data))
	}
	return contents
}

func userInput() string {
	var letter string

	print("Введите следующую букву: ")
	fmt.Scan(&letter)
	return strings.ToUpper(letter)
}

func figure() string {
	figures := readFiles()

	return figures[game.ErrorsMade()]
}

func showErrors() string {
	return strings.Join(game.GetErrors(), ", ")
}

func printOut() {
	fmt.Printf("Загаданное слово %v\n", game.Result())
	fmt.Printf("%v\n", figure())
	fmt.Printf("Ошибки (%v): %v\n", game.ErrorsMade(), showErrors())
	fmt.Printf("У вас осталось ошибок: %v\n\n", game.ErrorsBalance())

	if game.IsUserWon() {
		fmt.Println("Поздравляем, вы выиграли!")
	} else if game.IsUserLost() {
		fmt.Printf("Вы проиграли, загаданное слово: %v\n", game.Word())
	}
}
