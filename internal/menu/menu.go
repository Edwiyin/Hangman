package menu

import (
	"fmt"
	"hangman/internal/game"
	"hangman/internal/input"
	"time"

	"github.com/fatih/color"
)

func ShowMainMenu() string {
	fmt.Println("\n===== Main Menu =====")
	fmt.Println("1. Start New Game")
	fmt.Println("2. Show Rules")
	fmt.Println("3. Quit")
	return input.GetMenuChoice("Enter your choice (1-3): ", 3)
}

func SelectDifficulty() game.Difficulty {
	fmt.Println("\n===== Select Difficulty =====")
	fmt.Println("1. Easy")
	fmt.Println("2. Medium")
	fmt.Println("3. Hard")
	fmt.Println("4. Return to main menu")
	choice := input.GetMenuChoice("Enter your choice (1-4): ", 4)

	switch choice {
	case "1":
		return game.Easy
	case "2":
		return game.Medium
	case "3":
		return game.Hard
	case "4":
		MessageRapide("Returning to main menu...", 50, "vert")
	default:
		MessageRapide("Invalid choice, returning to main menu...", 50, "rouge")
	}
	return game.Easy
}

func ShowRules() {
	MessageRapide("\n===== Game Rules =====", 20, "vert")
	MessageRapide("1. The game will choose a random word based on the selected difficulty.", 20, "vert")
	MessageRapide("2. You need to guess the word by suggesting letters or the whole word.", 20, "vert")
	MessageRapide("3. You have a limited number of attempts based on the difficulty:", 20, "vert")
	MessageRapide(fmt.Sprintf("   - Easy: %d attempts, %d letters revealed\n", game.DifficultyConfig[game.Easy].MaxTries, game.DifficultyConfig[game.Easy].InitialReveals), 20, "vert")
	MessageRapide(fmt.Sprintf("   - Medium: %d attempts, %d letter revealed\n", game.DifficultyConfig[game.Medium].MaxTries, game.DifficultyConfig[game.Medium].InitialReveals), 20, "vert")
	MessageRapide(fmt.Sprintf("   - Hard: %d attempts, %d letters revealed\n", game.DifficultyConfig[game.Hard].MaxTries, game.DifficultyConfig[game.Hard].InitialReveals), 20, "vert")
	MessageRapide("4. Guessing an incorrect letter reduces your remaining attempts by 1.", 20, "vert")
	MessageRapide("5. Guessing an incorrect word reduces your remaining attempts by 2.", 20, "vert")
	MessageRapide("6. The game ends when you guess the word correctly or run out of attempts.", 20, "vert")
	MessageRapide("\nPress Enter to return to the main menu...", 20, "vert")
	input.GetPlayerGuess(nil)
}

func MessageRapide(message string, vitesse int, nomCouleur string) {
	c := color.New(color.FgWhite)
	switch nomCouleur {
	case "vert":
		c = color.New(color.FgGreen)
	case "rouge":
		c = color.New(color.FgRed)
	case "bleu":
		c = color.New(color.FgBlue)
	case "cyan":
		c = color.New(color.FgCyan)
	case "jaune":
		c = color.New(color.FgYellow)
	}

	for _, char := range message {
		fmt.Print(c.Sprint(string(char)))
		time.Sleep(time.Duration(vitesse) * time.Millisecond)
	}

	fmt.Println()
}
