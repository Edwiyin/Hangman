package hangman

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Game struct {
	word            string
	guessedLetters  map[rune]bool
	remainingTries  int
	revealedLetters []rune
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	word := wordList[rand.Intn(len(wordList))]
	revealed := revealRandomLetters(word)
	return &Game{
		word:            word,
		guessedLetters:  make(map[rune]bool),
		remainingTries:  10,
		revealedLetters: revealed,
	}
}

func (g *Game) Play() {
	color.Green("Bienvenue au jeu du Pendu !")
	color.Yellow("Le mot a %d lettres. Bonne chance !\n", len(g.word))

	for g.remainingTries > 0 && !g.isWordGuessed() {
		g.displayCurrentState()
		guess := g.getPlayerGuess()
		g.processGuess(guess)
	}

	g.displayCurrentState()
	if g.isWordGuessed() {
		color.Green("Félicitations ! Vous avez gagné !")
	} else {
		color.Red("Désolé, vous avez perdu. Le mot était : %s\n", g.word)
	}
}

func (g *Game) displayCurrentState() {
	color.Cyan(hangmanArt[10-g.remainingTries])
	fmt.Printf("Mot : ")
	for _, letter := range g.word {
		if g.guessedLetters[letter] || containsRune(g.revealedLetters, letter) {
			color.Green("%c ", letter)
		} else {
			color.Yellow("_ ")
		}
	}
	fmt.Println()
	color.Blue("Essais restants : %d\n", g.remainingTries)
	color.Magenta("Lettres devinées : ")
	for letter := range g.guessedLetters {
		fmt.Printf("%c ", letter)
	}
	fmt.Println()
}

func (g *Game) getPlayerGuess() string {
	var guess string
	color.White("Entrez une lettre ou un mot : ")
	fmt.Scanln(&guess)
	return strings.ToLower(guess)
}

func (g *Game) processGuess(guess string) {
	if len(guess) == 1 {
		letter := rune(guess[0])
		if g.guessedLetters[letter] {
			color.Yellow("Vous avez déjà deviné cette lettre.")
			return
		}
		g.guessedLetters[letter] = true
		if !strings.ContainsRune(g.word, letter) {
			g.remainingTries--
			color.Red("Lettre incorrecte.")
		} else {
			color.Green("Bonne devinette !")
		}
	} else if len(guess) == len(g.word) {
		if guess == g.word {
			for _, letter := range g.word {
				g.guessedLetters[letter] = true
			}
		} else {
			g.remainingTries -= 2
			color.Red("Mot incorrect. Vous perdez 2 essais.")
		}
	} else {
		color.Yellow("Veuillez entrer une seule lettre ou le mot entier.")
	}
}

func (g *Game) isWordGuessed() bool {
	for _, letter := range g.word {
		if !g.guessedLetters[letter] && !containsRune(g.revealedLetters, letter) {
			return false
		}
	}
	return true
}

func revealRandomLetters(word string) []rune {
	numToReveal := len(word)/4 + 1
	revealed := make([]rune, 0, numToReveal)
	for len(revealed) < numToReveal {
		idx := rand.Intn(len(word))
		letter := rune(word[idx])
		if !containsRune(revealed, letter) {
			revealed = append(revealed, letter)
		}
	}
	return revealed
}

func containsRune(slice []rune, r rune) bool {
	for _, v := range slice {
		if v == r {
			return true
		}
	}
	return false
}