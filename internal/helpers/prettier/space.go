package prettier

import (
	"strings"
	"unicode"
)

// RemoveSpaces removes all spaces from a string.
func RemoveSpaces(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}
