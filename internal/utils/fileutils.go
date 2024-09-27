package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

func ReadWordsFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func Jaune(texte string) string {
	return "\033[33m" + texte + "\033[0m"
}

func Vert(texte string) string {
	return "\033[32m" + texte + "\033[0m"
}

func Rouge(texte string) string {
	return "\033[31m" + texte + "\033[0m"
}

func Cyan(texte string) string {
	return "\033[36m" + texte + "\033[0m"
}

func Bleu(texte string) string {
	return "\033[34m" + texte + "\033[0m"
}

func AfficherLigneMenu(texte string, largeur int) {
	espaces := largeur - len([]rune(texte)) - 6
	fmt.Printf("%s║ %s%s ║%s\n", Vert("║"), Cyan(texte), strings.Repeat(" ", espaces), Vert("║"))
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
