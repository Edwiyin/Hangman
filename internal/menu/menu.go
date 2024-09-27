package menu

import (
	"fmt"
	"hangman/internal/game"
	"hangman/internal/input"
	"hangman/internal/utils"
	"strings"
	"time"

	"github.com/fatih/color"
)

const menuWidth = 50

func drawHorizontalLine(left, right string) {
	fmt.Print(utils.Vert(left))
	fmt.Print(utils.Vert(strings.Repeat("═", menuWidth-2)))
	fmt.Println(utils.Vert(right))
}

func showTableau(title string, options []string) string {
	drawHorizontalLine("╔", "╗")
	utils.AfficherLigneMenu(title, menuWidth)
	drawHorizontalLine("╠", "╣")
	for i, option := range options {
		utils.AfficherLigneMenu(fmt.Sprintf("%d. %s", i+1, option), menuWidth)
	}
	drawHorizontalLine("╚", "╝")
	return input.GetMenuChoice(fmt.Sprintf("Enter your choice (1-%d): ", len(options)), len(options))
}

func ShowMainMenu() string {
	return showTableau("Main Menu", []string{
		"Start New Game",
		"Show Rules",
		"Quit",
	})
}

func SelectDifficulty() (game.Difficulty, bool) {
	choice := showTableau("Select Difficulty", []string{
		"Easy",
		"Medium",
		"Hard",
		"Return to main menu",
	})

	switch choice {
	case "1":
		return game.Easy, true
	case "2":
		return game.Medium, true
	case "3":
		return game.Hard, true
	case "4":
		MessageRapide("Returning to main menu...", 50, "vert")
		return "", false
	default:
		MessageRapide("Invalid choice, returning to main menu...", 50, "rouge")
		return "", false
	}
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
	fmt.Scanln()
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