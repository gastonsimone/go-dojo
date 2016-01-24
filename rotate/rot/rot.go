// Rotates an array k positions (default k=1).
// Memory consumption: O(1)
// Time: O(n), n = len(array)
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gastonsimone/go-dojo/rotate"
)

// errors
const (
	_ = iota
	EMPTY_ARRAY
	ERROR_ROTATING
)

var k int

func init() {
	flag.IntVar(&k, "r", 1, "how many positions the array will be rotated")
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\trotate_array [flags] e1 e2 ... # Rotates the array defined as [e1 e2 ...]\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	size := flag.NArg()
	if size <= 0 {
		fmt.Fprintln(os.Stderr, "Empty array.")
		os.Exit(EMPTY_ARRAY)
	}

	array := flag.Args()
	err := rotate.Rotate(array, k)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error rotating: %v.\n", err)
		os.Exit(ERROR_ROTATING)
	}
	fmt.Println(array)
}
