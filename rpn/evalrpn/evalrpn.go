// Evaluates Reverse Polish Notation
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gastonsimone/go-dojo/rpn"
)

// errors
const (
	_ = iota
	ERROR_INVALID_ARGS
	ERROR_EVALUATING
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tevalrpn ... \n")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	size := flag.NArg()
	if size <= 0 {
		fmt.Fprintln(os.Stderr, "No formula to evaluate.")
		os.Exit(ERROR_INVALID_ARGS)
	}

	args := flag.Args()
	result, err := rpn.Evalrpn(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error evaluating: %v.\n", err)
		os.Exit(ERROR_EVALUATING)
	} else {
		fmt.Println(result)
	}
}
