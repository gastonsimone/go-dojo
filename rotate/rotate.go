// Rotates an array k positions (default k=1).
// Memory consumption: O(1)
// Time: O(n), n = len(array)
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
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
	err := rotate(array, k)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error rotating: %v.\n", err)
		os.Exit(ERROR_ROTATING)
	}
	fmt.Println(array)
}

// rotates slice in place k positions
// Memory: O(1)
// Time: O(n), n = len(array)
func rotate(array []string, k int) error {
	if k < 0 {
		return errors.New("invalid number of positions to rotate")
	}

	if len(array) <= 0 {
		return errors.New("empty array")
	}

	// Do not rotate more than necessary
	k = k % len(array)

	split := len(array) - k
	reverse(array[:split])
	reverse(array[split:])
	reverse(array)
	return nil
}

// reverses a slice in place.
func reverse(a []string) {
	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}
