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

func getRandomWord(wordList []string, wordsUsed map[string]bool) string {
	var word string
	for {
		word = wordList[rand.Intn(len(wordList))]
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
	blanks := make([]string, len(word))
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

func updateBlanks(word string, blanks []string, letter string) {
	letters := strings.Split(word, "")
	for i, _ := range letters {
		if letters[i] == letter {
			blanks[i] = letter
		}
	}
}

func printGameStatus(blanks []string, guesses []string) {
	if len(guesses) > 0 {
		fmt.Println("guesses: ", strings.Join(guesses, ""))
	}
	fmt.Println(strings.Join(blanks, " "))
}

func wonGame(blanks []string) bool {
	for _, letter := range blanks {
		if letter == "_" {
			return false
		}
	}
	return true
}

func main() {
	wordList, wordsUsed := loadWords("data/valid_words_list.txt")
	rand.Seed(time.Now().Unix())
	word := getRandomWord(wordList, wordsUsed)
	blanks := getBlanksForWord(word)
	guesses := make([]string, 0)
	for {
		printGameStatus(blanks, guesses)
		guess := getGuess()
		guesses = append(guesses, guess)
		updateBlanks(word, blanks, guess)
		if wonGame(blanks) {
			fmt.Println(strings.Join(strings.Split(word, ""), " "))
			fmt.Println("YOU WON!!")
			return
		}
	}
}
