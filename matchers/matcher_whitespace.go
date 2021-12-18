package matchers

import (
	"unicode"
)

func StartWhitespaceMatcher() LexxMatcherMatch {
	length := 0
	lines := 0
	return func(r rune, currentText []rune) (token *Token, precedence int8, run bool) {
		if r == 0 {
			if length > 0 {
				return &Token{Type: WHITESPACE, Value: string(currentText), Line: lines, Column: length}, 0, false
			} else {
				return nil, 0, false
			}
		}

		if unicode.IsSpace(r) {
			length++
			if r == '\n' {
				lines++
				length = 0
			}
			return nil, 0, true
		} else if lines == 0 && length == 0 {
			return nil, 0, false
		} else {
			return &Token{Type: WHITESPACE, Value: string(currentText), Line: lines, Column: length}, 1, false
		}
	}
}
