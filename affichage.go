package hangman

import (
)

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

