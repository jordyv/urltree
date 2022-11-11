package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jordyv/urltree/internal/builder"
	"github.com/jordyv/urltree/internal/output"
)

var (
	noQueries bool
)

func main() {
	flag.BoolVar(&noQueries, "noQueries", false, "don't print query strings")
	flag.Parse()

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "no input, use 'cat urls | urltree'")
		os.Exit(1)
	}

	t := builder.BuildTree(os.Stdin)
	output.Print(os.Stdout, t, noQueries)
}
