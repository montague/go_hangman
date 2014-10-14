package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var WordList = make([]string, 0)
var WordsUsed = make(map[string]bool)

func loadWords() {
	wordsFilePath := "data/valid_words_list.txt"
	f, err := os.Open(wordsFilePath)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		word := scanner.Text()
		WordList = append(WordList, word)
		WordsUsed[word] = false
	}
}

func initializeGame() {
	loadWords()
	rand.Seed(time.Now().Unix())
}

func getWord() string {
	word := WordList[rand.Intn(len(WordList))]
	WordsUsed[word] = true
	return word
}

type GameState struct {
	Misses int
	Word   string
	Blanks map[string]string
}

func NewGame(word string) *GameState {
	g := &GameState{Word: word}
	g.Blanks = make(map[string]string)
	for _, c := range strings.Split(word, "") {
		g.Blanks[c] = "_"
	}
	return g
}

func getGuess() string {
	fmt.Print("guess a letter: ")
	bio := bufio.NewReader(os.Stdin)
	line, _, _ := bio.ReadLine()
	return string(line[0])
}

func main() {
	initializeGame()
	guess := getGuess()
	fmt.Println("guess: ", guess)
	fmt.Println("word: ", getWord())
}
