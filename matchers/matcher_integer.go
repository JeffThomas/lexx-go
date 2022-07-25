package matchers

import (
	"unicode"
)

func StartIntegerMatcher() LexxMatcherMatch {
	length := 0
	return func(r rune, currentText []rune) (token *Token, precedence int8, run bool) {

		if r == 0 {
			if length > 0 {
				return &Token{Type: INTEGER, Value: string(currentText), Column: length}, 0, false
			} else {
				return nil, 0, false
			}
		}

		if unicode.IsDigit(r) {
			length++
			return nil, 0, true
		} else if length == 0 {
			return nil, 0, false
		} else {
			return &Token{Type: INTEGER, Value: string(currentText), Column: length}, 0, false
		}
	}
}
