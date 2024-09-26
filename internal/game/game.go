package game

import (
	"fmt"
	"hangman/internal/input"
	"hangman/internal/utils"
	"strings"
	"time"
)

func mapKeysToString(m map[rune]bool) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, string(k))
	}
	return strings.Join(keys, ", ")
}

type Game struct {
	word           *Word
	guessedLetters map[rune]bool
	remainingTries int
	difficulty     Difficulty
}

func NewGame(words []string, difficulty Difficulty) *Game {
	settings := DifficultyConfig[difficulty]
	return &Game{
		word:           NewWord(words),
		guessedLetters: make(map[rune]bool),
		remainingTries: settings.MaxTries,
		difficulty:     difficulty,
	}
}

func (g *Game) Play() {
	fmt.Printf("\nStarting a new game on %s difficulty!\n", g.difficulty)
	settings := DifficultyConfig[g.difficulty]
	g.word.RevealRandomLetters(settings.InitialReveals)

	for !g.IsGameOver() {
		g.DisplayGameState()
		guess := input.GetPlayerGuess(g.guessedLetters)

		if len(guess) == 1 {
			g.ProcessLetterGuess(rune(guess[0]))
		} else {
			g.ProcessWordGuess(guess)
		}
	}

	g.DisplayGameResult()
}

func (g *Game) ProcessLetterGuess(letter rune) {
	g.guessedLetters[letter] = true
	if !g.word.RevealLetter(letter) {
		g.remainingTries--
		fmt.Println("Incorrect guess!")
	} else {
		fmt.Println("Correct guess!")
	}
}

func (g *Game) ProcessWordGuess(word string) {
	if g.word.Guess(word) {
		g.word.RevealAllLetters()
		fmt.Println("Correct word guess!")
	} else {
		g.remainingTries -= 2
		fmt.Println("Incorrect word guess! You lose 2 tries.")
	}
	fmt.Printf("Guessed letters: %s\n", mapKeysToString(g.guessedLetters))
}

func (g *Game) IsGameOver() bool {
	return g.remainingTries <= 0 || g.word.IsFullyRevealed()
}

func (g *Game) DisplayGameState() {
	fmt.Printf("\nWord: %s\n", g.word.GetDisplayWord())
	fmt.Printf("Remaining tries: %d\n", g.remainingTries)
	fmt.Printf("Guessed letters: %s\n", mapKeysToString(g.guessedLetters))
	utils.PrintHangman(DifficultyConfig[g.difficulty].MaxTries - g.remainingTries)
}

func (g *Game) DisplayGameResult() {
	if g.word.IsFullyRevealed() {
		fmt.Println("Congratulations! You've won!")
	} else {
		fmt.Printf("Game over! The word was: %s\n", g.word.GetFullWord())
		fmt.Println("Better luck next time!")
		fmt.Println(utils.Rouge(`
		 ._________.
		 |/   |    
		 |    O    
		 |   /|\   
		 |  ° | °   
		 |   / \   
		 |       
	     ____|____

		`))
		time.Sleep(2 * time.Second)
	}
}
