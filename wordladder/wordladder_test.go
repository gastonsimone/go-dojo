package wordladder

import (
	"bufio"
	"errors"
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

type badReader struct {
}

func (r *badReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("bad reader")
}

func TestLoadWordDictBadReader(t *testing.T) {
	scanner := bufio.NewScanner(&badReader{})
	if _, err := LoadWordDict(scanner); err == nil {
		t.Fatalf("expected error loading dictionary with bad reader")
	}
}

func TestLoadWordDictNilScanner(t *testing.T) {
	if _, err := LoadWordDict(nil); err == nil {
		t.Fatal("no error with nil scanner")
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

	ladder := WordLadder("hit", "cog", dict)

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

	ladder := WordLadder("small", "large", dict)

	want := []string{"small", "shall", "shale", "shave", "seave", "serve", "verve", "varve", "larve", "large"}
	if !strarray.AreEqual(ladder, want) {
		t.Fatalf("got %v, want %v", ladder, want)
	}
}

func BenchmarkLoadWordDict(b *testing.B) {
	const input = "one two three four five six seven eight nine ten eleven"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		scanner := bufio.NewScanner(strings.NewReader(input))
		scanner.Split(bufio.ScanWords)
		LoadWordDict(scanner)
	}
}

func BenchmarkWordLadder(b *testing.B) {
	file, err := os.Open("/usr/share/dict/words")
	if err != nil {
		b.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var dict WordSet
	if dict, err = LoadWordDict(scanner); err != nil {
		b.Fatalf("error loading dictionary: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		WordLadder("small", "large", dict)
	}
}
