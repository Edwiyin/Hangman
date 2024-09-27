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
		fmt.Println(utils.Rouge("Incorrect guess!"))
	} else {
		fmt.Println(utils.Vert("Correct guess!"))
	}
}

func (g *Game) ProcessWordGuess(word string) {
	if g.word.Guess(word) {
		g.word.RevealAllLetters()
		fmt.Println(utils.Vert("Correct word guess!"))
	} else {
		g.remainingTries -= 2
		fmt.Println(utils.Rouge("Incorrect word guess! You lose 2 tries."))
	}
}

func (g *Game) IsGameOver() bool {
	return g.remainingTries <= 0 || g.word.IsFullyRevealed()
}

func (g *Game) DisplayGameState() {
	fmt.Print("\033[2J")  // Clear the screen
	fmt.Print("\033[H")   // Move cursor to top-left corner

	utils.PrintHangman(DifficultyConfig[g.difficulty].MaxTries - g.remainingTries)

	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("Word: %s\n", utils.Cyan(g.word.GetDisplayWord()))
	fmt.Printf("Remaining tries: %s\n", utils.Jaune(fmt.Sprintf("%d", g.remainingTries)))
	fmt.Printf("Guessed letters: %s\n", utils.Bleu(mapKeysToString(g.guessedLetters)))
	fmt.Println(strings.Repeat("-", 40))
}

func (g *Game) DisplayGameResult() {
	if g.word.IsFullyRevealed() {
		fmt.Println(utils.Vert("Congratulations! You've won!"))
		fmt.Println(utils.Vert(`
	        _____________ 
                '._==_==_=_.'
               .-\:        /-.
              | (|:.      |) |
               '-|:.      |-'
                 \::.     /
                  '::. .'
                    ) (
                  _.' '._
                 '-------'
		`))
	} else {
		fmt.Printf("%s %s\n", utils.Rouge("Game over! The word was:"), utils.Cyan(g.word.GetFullWord()))
		fmt.Println(utils.Rouge("Better luck next time!"))
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
	}
	time.Sleep(5 * time.Second)
}
