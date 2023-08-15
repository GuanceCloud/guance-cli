package promql

import (
	"regexp"
)

// GenName generates a name for a promql expression used on Guance Dashboard
func GenName(query string) (string, error) {
	matches := namePattern.FindStringSubmatch(query)
	if len(matches) == 0 {
		return query, nil
	}
	return matches[1], nil
}

var namePattern = regexp.MustCompile(`^[^:]+:([^{]+)`)
