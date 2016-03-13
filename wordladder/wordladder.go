package wordladder

import (
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/gastonsimone/go-dojo/strarray"
)

// WordSet represents a dictionary of words in memory.
type WordSet map[string]bool

// LoadWordDict loads a dictionary of words in memory and returns it as
// a 'WordSet'.  'scanner' is is used as the source of words.
func LoadWordDict(scanner *bufio.Scanner) (WordSet, error) {
	if scanner == nil {
		return nil, errors.New("nil scanner")
	}

	dict := make(WordSet)
	isWord := regexp.MustCompile(`^\w+$`)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		if isWord.MatchString(word) {
			dict[word] = false
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return dict, nil
}

// wordLadder represents a backward list of steps until the starting word
type wordLadder struct {
	word   string
	parent *wordLadder // parent word in the ladder (used to get the full conversion path)
}

// queue represents a FIFO queue of words in the ladder
type queue []*wordLadder

// Empty returns true if q has no elements
func (q *queue) empty() bool {
	return (len(*q) == 0)
}

// Enqueue adds w at the end of q
func (q *queue) enqueue(w *wordLadder) {
	*q = append(*q, w)
}

// Enqueue returns and removes the first element from q
func (q *queue) pop() *wordLadder {
	w := (*q)[0]
	*q = (*q)[1:]
	return w
}

// WordLadder calculates the 'word ladder' between two words.
// Given a 'start' word, an 'end' word, and a dictionary of words 'dict', finds
// the shortest transformation sequence from 'start' to 'end', such that only
// one letter can be changed at a time. Each intermediate word must exist in
// the dictionary.
func WordLadder(start, end string, dict WordSet) []string {
	if len(start) != len(end) {
		return nil
	}

	root := new(wordLadder)
	root.word = start

	visited := make(WordSet)
	visited[start] = true

	q := new(queue)
	q.enqueue(root)

	for !q.empty() {
		step := q.pop()
		word := step.word

		chars := strings.Split(word, "")

		for i, c := range chars {
			for newc := 'a'; newc <= 'z'; newc++ {
				newchar := fmt.Sprintf("%c", newc)

				if c == newchar {
					continue
				}

				chars[i] = newchar
				newword := strings.Join(chars, "")
				chars[i] = c

				newstep := new(wordLadder)
				newstep.word = newword
				newstep.parent = step

				if newword == end {
					return newstep.climbLadder()
				}

				if _, exists := dict[newword]; exists {
					if _, found := visited[newword]; !found {
						q.enqueue(newstep)
						visited[newword] = true
					}
				}
			}
		}
	}

	return nil
}

// ClimbLadder traverses the words in the ladder, starting from w and returns
// a slice of words, where the first word is the first word in the ladder, and
// the last word is w.word.
func (w *wordLadder) climbLadder() []string {
	var words []string
	for ; w != nil; w = w.parent {
		words = append(words, w.word)
	}

	strarray.Reverse(words)
	return words
}
