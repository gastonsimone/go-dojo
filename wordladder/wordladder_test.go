package wordladder

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/gastonsimone/go-dojo/strarray"
)

func TestLoadWordDict(t *testing.T) {
	const input = "one two three four five"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	var dict WordSet
	var err error
	if dict, err = LoadWordDict(scanner); err != nil {
		t.Fatalf("error loading dictionary: %v", err)
	}

	if len(dict) != 5 {
		t.Fatal("unexpected number of words in dictionary")
	}
}

func TestWordLadderNoDict(t *testing.T) {
	ladder := WordLadder("hit", "cog", nil)

	var want []string = nil
	if !strarray.AreEqual(ladder, want) {
		t.Fatalf("got %v, want %v", ladder, want)
	}
}

func TestWordLadderDifferentLength(t *testing.T) {
	ladder := WordLadder("small", "larger", nil)

	var want []string = nil
	if !strarray.AreEqual(ladder, want) {
		t.Fatalf("got %v, want %v", ladder, want)
	}
}

func TestWordLadderExactDict(t *testing.T) {
	const input = "hot dot dog"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	var dict WordSet
	var err error
	if dict, err = LoadWordDict(scanner); err != nil {
		t.Fatalf("error loading dictionary: %v", err)
	}

	var ladder []string
	ladder = WordLadder("hit", "cog", dict)

	want := []string{"hit", "hot", "dot", "dog", "cog"}
	if !strarray.AreEqual(ladder, want) {
		t.Fatalf("got %v, want %v", ladder, want)
	}
}

func TestWordLadderLargeConversion(t *testing.T) {
	file, err := os.Open("/usr/share/dict/words")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var dict WordSet
	if dict, err = LoadWordDict(scanner); err != nil {
		t.Fatalf("error loading dictionary: %v", err)
	}

	var ladder []string
	ladder = WordLadder("small", "large", dict)

	want := []string{"small", "shall", "shale", "shave", "seave", "serve", "verve", "varve", "larve", "large"}
	if !strarray.AreEqual(ladder, want) {
		t.Fatalf("got %v, want %v", ladder, want)
	}
}
