package output

import (
	"fmt"
	"io"
	"sort"

	"github.com/jordyv/urltree/internal/entities"
)

func Print(w io.Writer, t entities.URLTree, noQueries bool) {
	var domains []string
	for d, _ := range t {
		domains = append(domains, d)
	}
	sort.Strings(domains)

	for _, d := range domains {
		fmt.Fprintf(w, "- %s\n", d)

		var paths []string
		for p, _ := range t[d] {
			paths = append(paths, p)
		}
		sort.Strings(paths)

		for _, p := range paths {
			fmt.Fprintf(w, "  -- %s\n", p)

			if !noQueries {
				qs := t[d][p]
				sort.Strings(qs)
				for _, q := range qs {
					fmt.Fprintf(w, "    --- ?%s\n", q)
				}
			}
		}
	}
}
