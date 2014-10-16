package main

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

func TestGetBlanksForWord(t *testing.T) {
	blanks := getBlanksForWord("omg")
	if strings.Join(blanks, "") != "___" {
		t.Error("blanks should be _ _ _ for omg")
	}
}

func TestWonGame(t *testing.T) {
	if !wonGame([]string{"a", "b"}) {
		t.Error("should win game if no blanks left")
	}
	if wonGame([]string{"_", "b"}) {
		t.Error("should not win game if blanks are left")
	}
}

func TestUpdateBlanks(t *testing.T) {
	blanks := []string{"_", "_", "_"}
	word := "omg"
	letter := "m"
	updateBlanks(word, blanks, letter)
	if blanks[1] != "m" {
		t.Error("UpdateBlanks should update the correct letter")
	}
}
