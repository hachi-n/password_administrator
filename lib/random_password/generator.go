package random_password

import (
	"github.com/hachi-n/passwd_gen/lib/util"
	"math/rand"
	"time"
)

const (
	lowerCase  = "abcdefghijklmnopqrstuvwxyz"
	upperCase  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberCase = "0123456789"
	symbolCase = "`~!@#$%^&*()_+-={}[]\\|:;\"'<>,.?/"
)

const (
	NormalType = iota
	NumberType
	SymbolType
)

func GeneratePassword(length, numberLength, symbolLength int) string {
	runeTypeElements := createRuneTypeElements(length, numberLength, symbolLength)
	passwordStr := createRandomRunes(runeTypeElements)
	return passwordStr
}

func createRuneTypeElements(length, numberLength, symbolLength int) []int {
	charLength := length - numberLength - symbolLength

	runeTypeNums := []int{charLength, numberLength, symbolLength}
	var runeTypeElements []int
	for index, num := range runeTypeNums {
		runeTypeElements = collectRuneTypeElements(runeTypeElements, index, num)
	}

	util.Shuffle(runeTypeElements[1:])

	return runeTypeElements
}

func createRandomRunes(runeTypeElements []int) string {
	var passwordRunes []rune
	for _, e := range runeTypeElements {
		passwordRunes = append(passwordRunes, GenerateRandomRune(e))
	}

	return string(passwordRunes)
}

func collectRuneTypeElements(slice []int, index int, num int) []int {
	for i := 0; i < num; i++ {
		slice = append(slice, index)
	}
	return slice
}

func GenerateRandomRune(runeType int) rune {
	var selection string
	switch runeType {
	case NormalType:
		selection = lowerCase+upperCase
	case NumberType:
		selection = numberCase
	case SymbolType:
		selection = symbolCase
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(selection))

	r := []rune(selection)[index]
	return r
}

