package main

import (
	"fmt"
	"math/rand"
)

func generatePassword(length int, levelPassword string) string {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	newPassword := ""

	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	symbols := []string{"!", "#", "$", "%", "&", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", "=", ">", "?", "@", "[", "]", "^", "_", "{", "|", "}", "~"}

	if levelPassword == "easy" {
		for i := 0; i < length; i++ {
			randomSymbol := alphabet[rand.Intn(52)] // я умею брать только первый символ из строки, а нужно взять один случайный символ :(
			newPassword += string(randomSymbol)
		}

	}
	if levelPassword == "medium" {
		for i := 0; i < length; i++ {
			medium := [][]string{alphabet, digits}
			mediumNew := medium[rand.Intn(len(medium))]

			mediumFull := mediumNew[rand.Intn(len(mediumNew))]

			newPassword += mediumFull
		}
	}

	if levelPassword == "hard" {
		for i := 0; i < length; i++ {
			hard := [][]string{alphabet, digits, symbols}
			hardNew := hard[rand.Intn(len(hard))]

			hardFull := hardNew[rand.Intn(len(hardNew))]
			newPassword += hardFull
		}
	}

	return newPassword
}

func main() {
	fmt.Println(generatePassword(10, "hard"))
}
