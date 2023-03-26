package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyz"
	LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers = "0123456789"
	symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

var Reader io.Reader

type Password struct {
	length   int
	letters  string
	LETTERS  string
	numbers  string
	symbols  string
	password string
}

func main() {
	// router group with save passwords, create db, view specific password (protected), generate password
	var pw *Password = &Password{}
	router := gin.Default()
	router.GET("/password/:length", pw.PasswordInit)

	router.Run("localhost:8080")
	password, err := pw.PasswordInit(12)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(password)
}

func (pw *Password) PasswordInit(input int) (string, error) {
	if input < 8 || input > 64 {
		return "", errors.New("invalid chosen password length, please use a password between 8 and 64 characters.")
	}
	pw.length = input
	return pw.GeneratePassword(), nil
}

func (pw *Password) GeneratePassword() string {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	altChars := pw.length / 5
	length := pw.GenerateCapitals(altChars + random.Intn(pw.length/12+1))
	length += pw.GenerateNumbers(altChars + random.Intn(pw.length/12+1))
	length += pw.GenerateSymbols(altChars + random.Intn(pw.length/12+1))
	pw.GenerateLetters(pw.length - length)
	passArray := []rune(pw.LETTERS + pw.numbers + pw.symbols + pw.letters)
	fmt.Println(pw.LETTERS + pw.numbers + pw.symbols + pw.letters)
	rand.Shuffle(pw.length, func(i, j int) {
		passArray[i], passArray[j] = passArray[j], passArray[i]
	})
	return string(passArray)
}

func (pw *Password) GenerateLetters(length int) {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	for i := 0; i < length; i++ {
		rnd := random.Intn(len(letters))
		pw.letters += letters[rnd : rnd+1]
	}
}

func (pw *Password) GenerateCapitals(length int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	for i := 0; i < length; i++ {
		rnd := random.Intn(len(LETTERS))
		pw.LETTERS += LETTERS[rnd : rnd+1]
	}
	return len(pw.LETTERS)
}

func (pw *Password) GenerateNumbers(length int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	for i := 0; i < length; i++ {
		rnd := random.Intn(len(numbers))
		pw.numbers += numbers[rnd : rnd+1]
	}
	return len(pw.numbers)
}

func (pw *Password) GenerateSymbols(length int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	for i := 0; i < length; i++ {
		rnd := random.Intn(len(symbols))
		pw.symbols += symbols[rnd : rnd+1]
	}
	return len(pw.symbols)
}

func (pw *Password) Iterator(values string) (string, int) {
	if len(values) < 1 {
		return "", -1
	}

	pw.password += values[0:1]
	return values[1:], 0
}
