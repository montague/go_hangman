package hangman

import "testing"

func TestNewGame(t *testing.T) {
	g := NewGame("omg")

	if g.Misses != 0 {
		t.Error("New game should start with 0 misses")
	}
	if len(g.Blanks) != 3 {
		t.Error("Blanks should have length 3")
	}
	if g.Blanks["o"] != "_" {
		t.Error("Blanks should have underscore for value: o")
	}
	if g.Blanks["m"] != "_" {
		t.Error("Blanks should have underscore for value: m")
	}
	if g.Blanks["g"] != "_" {
		t.Error("Blanks should have underscore for value: g")
	}
	if g.Word != "omg" {
		t.Error("Word should be 'omg'")
	}
}
