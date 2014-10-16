package main

import "testing"
import "strings"

func TestNewGame(t *testing.T) {
	g := NewGame("omg")

	if g.MissesLeft != 8 {
		t.Error("New game should start with 8 misses left")
	}
	if len(g.Blanks) != 3 {
		t.Error("Blanks should have length 3")
	}
	for i, _ := range strings.Split("omg", "") {
		if g.Blanks[i] != "_" {
			t.Error("Blanks should have underscore for index: ", i)
		}
	}
	if g.Word != "omg" {
		t.Error("Word should be 'omg'")
	}
	if g.LettersCorrect != 0 {
		t.Error("Letters correct should start at 0")
	}
}

func TestLoadWords(t *testing.T) {
	wordList, wordsUsed := loadWords("data/test_list.txt")
	for i, word := range []string{"firstword", "secondword", "thirdword"} {
		if wordList[i] != word {
			t.Error("word list should contain ", word, " at index ", i)
		}
		if wordsUsed[word] {
			t.Error("word used: \"", word, "\" should be false")
		}
	}
}

func TestWonGame(t *testing.T) {
	g = NewGame("omg")

}
