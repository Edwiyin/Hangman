package main

import (
	"fmt"
	"os"

	"hangman/internal/game"
	"hangman/internal/menu"
	"hangman/internal/utils"
)

func main() {
	utils.PrintGameTitle()

	wordFiles := map[game.Difficulty]string{
		game.Easy:   "assets/words_easy.txt",
		game.Medium: "assets/words_medium.txt",
		game.Hard:   "assets/words_hard.txt",
	}

	words := make(map[game.Difficulty][]string)
	for diff, file := range wordFiles {
		wordList, err := utils.ReadWordsFile(file)
		if err != nil {
			fmt.Printf("Error reading words file for %s difficulty: %v\n", diff, err)
			os.Exit(1)
		}
		words[diff] = wordList
	}

	for {
		choice := menu.ShowMainMenu()
		switch choice {
		case "1":
			difficulty, startGame := menu.SelectDifficulty()
			if startGame {
				g := game.NewGame(words[difficulty], difficulty)
				g.Play()
			}
		case "2":
			menu.ShowRules()
		case "3":
			fmt.Println("Thanks for playing Hangman! Goodbye!")
			return
		}
	}
}
