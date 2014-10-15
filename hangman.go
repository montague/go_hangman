package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func loadWords(wordsFile string) ([]string, map[string]bool) {
	f, err := os.Open(wordsFile)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	wordList := make([]string, 0)
	wordsUsed := make(map[string]bool)
	for scanner.Scan() {
		word := scanner.Text()
		wordList = append(wordList, word)
		wordsUsed[word] = false
	}
	return wordList, wordsUsed
}

func getRandomWord(wordList []string, wordsUsed [string]bool) string {
	for {
		word := wordList[rand.Intn(len(wordList))]
		if !wordsUsed[word] {
			wordsUsed[word] = true
			break
		}
	}
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

func getBlanksForWord(word string) []string {
	blanks = make([]string, len(word))
	for i, _ := range blanks {
		blanks[i] = "_"
	}
	return blanks
}

func getGuess() string {
	fmt.Print("guess a letter: ")
	bio := bufio.NewReader(os.Stdin)
	line, _, _ := bio.ReadLine()
	return string(line[0])
}

func printWordStatus(blanks []string) {
	fmt.Println(strings.Join(blanks, " "))
}

func main() {
	wordList, wordsUsed := loadWords("data/valid_words_list.txt")
	rand.Seed(time.Now().Unix())
	word := getRandomWord(wordList, wordsUsed)
	blanks := getBlanksForWord(word)
	guesses := make([]string, 0)
	guess = getGuess()
	guesses = append(guesses, guess)
	//guess := getGuess()
	fmt.Println("guess: ", guess)
	//printWordStatus(blanks)
	//fmt.Println("word: ", word)
}
