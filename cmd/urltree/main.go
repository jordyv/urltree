package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"sort"
	"strings"
)

var (
	noQueries bool
)

type URLTree map[string]map[string][]string

func main() {
	flag.BoolVar(&noQueries, "noQueries", false, "don't print query strings")
	flag.Parse()

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "no input, use 'cat urls | urltree'")
		os.Exit(1)
	}

	t := buildTree(os.Stdin)
	printTree(t)
}

func buildTree(input io.Reader) URLTree {
	t := make(URLTree)

	s := bufio.NewScanner(input)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			continue
		}

		u, err := url.Parse(line)
		if err != nil {
			continue
		}

		r := fmt.Sprintf("%s://%s", u.Scheme, u.Host)
		p := u.Path
		q := u.RawQuery

		if _, trFound := t[r]; !trFound {
			t[r] = make(map[string][]string)
		}

		if _, pFound := t[r][p]; !pFound {
			t[r][p] = make([]string, 0)
		}

		if q != "" && !contains(t[r][p], q) {
			t[r][p] = append(t[r][p], q)
		}
	}

	return t
}

func printTree(t URLTree) {
	var domains []string
	for d, _ := range t {
		domains = append(domains, d)
	}
	domains = sort.StringSlice(domains)

	for _, d := range domains {
		fmt.Printf("- %s\n", d)

		var paths []string
		for p, _ := range t[d] {
			paths = append(paths, p)
		}
		paths = sort.StringSlice(paths)

		for _, p := range paths {
			fmt.Printf("  -- %s\n", p)

			if !noQueries {
			qs := sort.StringSlice(t[d][p])
				for _, q := range qs {
					fmt.Printf("    --- %s\n", q)
				}
			}
		}
	}
}

func contains[T comparable](s []T, v T) bool {
	for _, i := range s {
		if i == v {
			return true
		}
	}
	return false
}
