package matchers

import (
	"unicode"
)

func StartPunctuationMatcher() LexxMatcherMatch {
	return func(r rune, currentText []rune) (token *Token, precedence int8, run bool) {
		if r == 0 {
			return nil, 0, false
		}

		if unicode.IsPunct(r) {
			return &Token{Type: PUNCTUATION, Value: string(r), Column: 1}, 1, false
		}
		return nil, 0, false
	}
}
