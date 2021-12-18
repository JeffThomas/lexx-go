package matchers

import (
	"unicode"
)

func StartSymbolicMatcher() LexxMatcherMatch {
	return func(r rune, currentText []rune) (token *Token, precedence int8, run bool) {
		if r == 0 {
			return nil, 0, false
		}

		if unicode.IsSymbol(r) {
			return &Token{Type: SYMBOLIC, Value: string(r), Column: 1}, 0, false
		}
		return nil, 0, false
	}
}
