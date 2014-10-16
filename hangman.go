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
	MissesLeft     int
	Word           string
	LettersCorrect int
	Blanks         []string
	LettersGuessed []string
}

func NewGame(word string) *GameState {
	g := &GameState{Word: word}
	g.MissesLeft = 8
	g.LettersCorrect = 0
	g.Blanks = make([]string, 0)
	for i := 0; i < len(word); i++ {
		g.Blanks = append(g.Blanks, "_")
	}
	return g
}

func (this *GameState) PrintStatus() {
	if len(this.LettersGuessed) > 0 {
		fmt.Println("already guessed: ", strings.Join(this.LettersGuessed, " "))
		fmt.Println("guesses left: ", this.MissesLeft)
	}

	fmt.Println(strings.Join(this.Blanks, " "))
}

func (this *GameState) Guess() {
	this.PrintStatus()
	fmt.Print("guess a letter: ")
	bio := bufio.NewReader(os.Stdin)
	line, _, _ := bio.ReadLine()
	guess := string(line[0])
	alreadyGuessed := false
	for _, c := range this.LettersGuessed {
		if c == guess {
			alreadyGuessed = true
			break
		}
	}
	if !alreadyGuessed {
		this.LettersGuessed = append(this.LettersGuessed, guess)
	}
	this.UpdateGuesses(guess)
}

func (this *GameState) UpdateGuesses(guess string) {
	letters := strings.Split(this.Word, "")
	missedGuess := true
	for i, _ := range letters {
		if letters[i] == guess {
			this.Blanks[i] = guess
			missedGuess = false
			this.LettersCorrect += 1
		}
	}
	if missedGuess {
		this.MissesLeft -= 1
	}
}

func (this *GameState) WonGame() bool {
	return len(this.Word) == this.LettersCorrect
}

func (this *GameState) LostGame() bool {
	return this.MissesLeft == 0
}

func main() {
	wordList, wordsUsed := loadWords("data/valid_words_list.txt")
	rand.Seed(time.Now().Unix())
	game := NewGame(getRandomWord(wordList, wordsUsed))
	for {
		game.Guess()
		if game.WonGame() {
			fmt.Println("YOU WON!!")
			return
		}
		if game.LostGame() {
			fmt.Println(strings.Join(strings.Split(game.Word, ""), " "))
			fmt.Println("YOU LOST!!")
			return
		}

	}
}
