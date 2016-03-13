// Given a start word, an end word and a dictionary (list of valid words)
// calculates the shortest "word ladder" from the start to the end words.
// A "word ladder" is a sequence of words (in the dictionary) that beginning in
// the start word, ending in the end word, and changing just one character from
// one word to the next in the sequence.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gastonsimone/go-dojo/wordladder"
)

// errors
const (
	_ = iota
	InvalidArgs
	OpenDict
	LoadDict
)

var dictfile string

func init() {
	flag.StringVar(&dictfile, "d", "/usr/share/dict/words", "dictionary file. Use - for standard input")
}

func usage() {
	fmt.Println(`Given a start word, an end word and a dictionary (list of valid words)
calculates the shortest "word ladder" from the start to the end words.
A "word ladder" is a sequence of words (in the dictionary) that beginning in
the start word, ending in the end word, and changing just one character from
one word to the next in the sequence.`)
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tladder [flags] start end\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
	fmt.Println(`Examples:
  Executing
  	ladder small large
  May produce the following output (depending on the dictionary):
  	[small shall shale shave seave serve verve varve larve large]
  Using a smaller dictionary:
  	ladder -d=words3 cog car
  Produces:
  	[cog cag car]
  From 'cog' to 'car' using only words that start with 'd':
  	grep '^d' words3 | ./ladder -d=- cog car
  Produces:
  	[cog dog dag dar car]`)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	size := flag.NArg()
	if size != 2 {
		fmt.Fprintln(os.Stderr, "Invalid number of arguments.")
		os.Exit(InvalidArgs)
	}

	array := flag.Args()
	start, end := array[0], array[1]

	var file *os.File
	var err error

	if dictfile == "-" {
		file = os.Stdin
	} else if file, err = os.Open(dictfile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(OpenDict)
	}

	scanner := bufio.NewScanner(file)
	dict, err := wordladder.LoadWordDict(scanner)
	file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading dictionary: %v", err)
		os.Exit(LoadDict)
	}

	ladder := wordladder.WordLadder(start, end, dict)

	fmt.Println(strings.Join(ladder, "\n"))
}
