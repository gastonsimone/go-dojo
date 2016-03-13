// if two strings s and t are isomorphic.
// Two strings are isomorphic if you can get one of them by replacing all the
// characters in the other.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gastonsimone/go-dojo/isomorphic"
)

// errors
const (
	_ = iota
	InvalidArgs
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tcheckisom s t  # Checks if strings s and t are isomorphic.\n")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Expecting two strings as parameters.\n")
		os.Exit(InvalidArgs)
	}

	s := args[0]
	t := args[1]

	if isomorphic.AreIsomorphic(s, t) {
		fmt.Printf("Yes, %v is isomorphic with %v.\n", s, t)
	} else {
		fmt.Printf("No, %v is not isomorphic with %v.\n", s, t)
	}
}
