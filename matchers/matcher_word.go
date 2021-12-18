package matchers

import (
	"unicode"
)

func StartWordMatcher() LexxMatcherMatch {
	length := 0
	return func(r rune, currentText []rune) (token *Token, precedence int8, run bool) {
		if r == 0 {
			if length > 0 {
				return &Token{Type: WORD, Value: string(currentText), Column: length}, 0, false
			} else {
				return nil, 0, false
			}
		}

		if unicode.IsLetter(r) || (length > 0 && unicode.IsDigit(r)) {
			length++
			return nil, 0, true
		} else if length == 0 {
			return nil, 0, false
		} else {
			return &Token{Type: WORD, Value: string(currentText), Column: length}, 0, false
		}
	}
}
