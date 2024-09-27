package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	textLength := len([]rune(texte))
	espaces := largeur - textLength - 6
	if espaces < 0 {
		espaces = 0
		texte = string([]rune(texte)[:largeur-6])
	}
	fmt.Printf("%s║ %s%s ║%s\n", Vert("║"), Cyan(texte), strings.Repeat(" ", espaces), Vert("║"))
}

func ResetColor() string {

    return "\033[0m"

}