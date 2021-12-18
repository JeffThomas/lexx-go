package matchers

import (
	"unicode"
)

func StartIdentifierMatcher() LexxMatcherMatch {
	length := 0
	return func(r rune, currentText []rune) (token *Token, precedence int8, run bool) {

		if length == 0 {
			if unicode.IsDigit(r) {
				return nil, 0, false
			}

			if r == '_' || unicode.IsLetter(r) {
				length++
				return nil, 0, true
			}
		}

		if r == 0 {
			if length > 0 {
				return &Token{Type: IDENTIFIER, Value: string(currentText), Column: length}, 5, false
			} else {
				return nil, 0, false
			}
		}

		if r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r) {
			length++
			return nil, 0, true
		}

		if length > 0 {
			return &Token{Type: IDENTIFIER, Value: string(currentText), Column: length}, 5, false
		} else {
			return nil, 0, false
		}
	}
}
