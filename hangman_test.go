package hangman

import "testing"
import "strings"

func TestNewGame(t *testing.T) {
	g := NewGame("omg")

	if g.Misses != 0 {
		t.Error("New game should start with 0 misses")
	}
	if len(g.Blanks) != 3 {
		t.Error("Blanks should have length 3")
	}
	for _, c := range strings.Split("omg", "") {
		if g.Blanks[c] != "_" {
			t.Error("Blanks should have underscore for value: ", c)
		}
	}
	if g.Word != "omg" {
		t.Error("Word should be 'omg'")
	}
}
