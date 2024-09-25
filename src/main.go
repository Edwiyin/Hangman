package main

import (
	"fmt"
	"github.com/Edwiyin/Hangman.git"

	"github.com/fatih/color"
)

func main() {
	for {
		displayMenu()
		choice := getUserChoice()

		switch choice {
		case 1:
			game := hangman.NewGame()
			game.Play()
		case 2:
			displayRules()
		case 3:
			color.Cyan("Merci d'avoir joué au Pendu ! Au revoir !")
			return
		default:
			color.Red("Choix invalide. Veuillez réessayer.")
		}
	}
}

func displayMenu() {
	color.Green("\n=== Menu du Jeu du Pendu ===")
	color.Yellow("1. Jouer au jeu")
	color.Yellow("2. Afficher les règles")
	color.Yellow("3. Quitter")
	fmt.Print("Entrez votre choix (1-3): ")
}

func getUserChoice() int {
	var choice int
	fmt.Scanf("%d", &choice)
	return choice
}

func displayRules() {
	color.Cyan("\n=== Règles du Jeu du Pendu ===")
	color.White("1. Un mot secret est choisi au hasard.")
	color.White("2. Vous devez deviner le mot en proposant des lettres.")
	color.White("3. Vous avez 10 essais pour deviner le mot complet.")
	color.White("4. Si vous devinez le mot avant d'épuiser vos essais, vous gagnez !")
	color.White("5. Si vous épuisez tous vos essais avant de deviner le mot, vous perdez.")
	fmt.Println()
}