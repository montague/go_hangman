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
	Letters        []string
	LettersCorrect int
	Blanks         []string
	LettersGuessed []string
}

func NewGame(word string) *GameState {
	g := &GameState{Word: word}
	g.MissesLeft = 8
	g.LettersCorrect = 0
	g.Blanks = make([]string, 0)
	for _, c := range word {
		g.Blanks = append(g.Blanks, "_")
		g.Letters = append(g.Letters, string(c))
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

func (this *GameState) Guess(guess string) {
	if guess == "\n" {
		return
	}
	alreadyGuessed := false
	for _, c := range this.LettersGuessed {
		if c == guess {
			alreadyGuessed = true
			break
		}
	}
	if !alreadyGuessed {
		this.LettersGuessed = append(this.LettersGuessed, guess)
		this.UpdateGuesses(guess)
	}
}

func (this *GameState) UpdateGuesses(guess string) {
	missedGuess := true
	for i, _ := range this.Letters {
		if this.Letters[i] == guess {
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

func getInputFromUser() string {
	bio := bufio.NewReader(os.Stdin)
	line, _ := bio.ReadString('\n')

	return string(line[0])
}

func getGuess() string {
	fmt.Print("guess a letter: ")
	return getInputFromUser()
}

func wantsToPlayAgain() bool {
	var response string
	for {
		fmt.Println("Play again? (y/n)")
		response = getInputFromUser()
		if response == "y" {
			return true
		}
		if response == "n" {
			return false
		}
	}
}

func gameLoop(wordList []string, wordsUsed map[string]bool) {
	game := NewGame(getRandomWord(wordList, wordsUsed))
	for {
		game.PrintStatus()
		game.Guess(getGuess())
		if game.WonGame() {
			fmt.Println(strings.Join(strings.Split(game.Word, ""), " "))
			fmt.Println("YOU WON!!")
			break
		}
		if game.LostGame() {
			fmt.Println(strings.Join(strings.Split(game.Word, ""), " "))
			fmt.Println("YOU LOST!!")
			break
		}
	}
}

func main() {
	wordList, wordsUsed := loadWords("data/valid_words_list.txt")
	rand.Seed(time.Now().Unix())
	for {
		gameLoop(wordList, wordsUsed)
		if !wantsToPlayAgain() {
			fmt.Println("Bye!")
			break
		}
	}
}
