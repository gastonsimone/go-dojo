// Rotates an array k positions (default k=1).
// Memory consumption: O(1)
// Time: O(n), n = len(array)
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gastonsimone/go-dojo/rotate"
)

// errors
const (
	_ = iota
	EmptyArray
)

var k int

func init() {
	flag.IntVar(&k, "r", 1, "how many positions the array will be rotated. Positive=right, negative=left.")
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
		os.Exit(EmptyArray)
	}

	array := flag.Args()
	rotate.Rotate(array, k)

	fmt.Println(strings.Join(array, " "))
}
