package wordladder

import (
	"bufio"
	"errors"
	"strings"
	"testing"
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
