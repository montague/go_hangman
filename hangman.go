package hangman

import (
	//"bufio"
	"fmt"
	"strings"
	//"os"
)

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

func main() {
	//word := "hello world"
	//fmt.Print("Enter some input: ")
	//bio := bufio.NewReader(os.Stdin)
	//line, _, _ := bio.ReadLine()
	//fmt.Printf("input was: %s\n", line)
	m := make([]string, 12)
	m = append(m, "omg")
	g := NewGame("omg")
	fmt.Println(g)
}
