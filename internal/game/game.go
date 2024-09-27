package game

import (
	"fmt"
	"hangman/internal/input"
	"hangman/internal/utils"
	"strings"
	"time"
)

func mapKeysToString(m map[rune]string) []string {
	var array []string = []string{}
	// keys := make([]string, 0, len(m))
	for _, v := range m {
		array = append(array, string(v))
		// keys = append(keys, v+string(k)+utils.ResetColor())
	}
	return array
}

type Game struct {
	word           *Word
	guessedLetters map[rune]string
	remainingTries int
	difficulty     Difficulty
}

func NewGame(words []string, difficulty Difficulty) *Game {
	settings := DifficultyConfig[difficulty]
	return &Game{
		word:           NewWord(words),
		guessedLetters: make(map[rune]string),
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
		guessedLettersBool := make(map[rune]bool)
		for k := range g.guessedLetters {
			guessedLettersBool[k] = true
		}
		guess := input.GetPlayerGuess(guessedLettersBool)

		if len(guess) == 1 {
			g.ProcessLetterGuess(rune(guess[0]))
		} else {
			g.ProcessWordGuess(guess)
		}
	}

	g.DisplayGameResult()
}

func (g *Game) ProcessLetterGuess(letter rune) {
	if g.word.RevealLetter(letter) {
		g.guessedLetters[letter] = utils.Vert(string(letter))
		fmt.Println(utils.Vert("Correct guess!"))
	} else {
		g.guessedLetters[letter] = utils.Rouge(string(letter))
		g.remainingTries--
		fmt.Println(utils.Rouge("Incorrect guess!"))
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
	fmt.Print("\033[2J")
	fmt.Print("\033[H")

	utils.PrintHangman(DifficultyConfig[g.difficulty].MaxTries - g.remainingTries)

	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("Word: %s\n", utils.Cyan(g.word.GetDisplayWord()))
	fmt.Printf("Remaining tries: %s\n", utils.Jaune(fmt.Sprintf("%d", g.remainingTries)))
	fmt.Printf("Guessed letters: ")
	array := mapKeysToString(g.guessedLetters)
	for _, v := range array {
		fmt.Printf("%s ", v)
	}
	fmt.Printf("\n")
	fmt.Println(strings.Repeat("-", 40))
}

func (g *Game) CheckGuessedLetters(input string) bool {
	for _, letter := range g.guessedLetters {
		if letter == input {
			return true
		}
	}
	return false
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
