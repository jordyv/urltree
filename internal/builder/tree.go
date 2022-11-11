package builder

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/jordyv/urltree/internal/entities"
)

func BuildTree(input io.Reader) entities.URLTree {
	t := make(entities.URLTree)

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

func contains[T comparable](s []T, v T) bool {
	for _, i := range s {
		if i == v {
			return true
		}
	}
	return false
}
