package matchers

import (
	"unicode"
)

func StartFloatMatcher() LexxMatcherMatch {
	length := 0
	dot := -1
	return func(r rune, currentText []rune) (token *Token, precedence int8, run bool) {

		if r == 0 {
			if length > 0 && dot > -1 && length > dot {
				return &Token{Type: FLOAT, Value: string(currentText), Column: length}, 0, false
			} else {
				return nil, 0, false
			}
		}

		if r == '.' {
			if dot > -1 && length == dot {
				return nil, 0, false
			}
			if dot == -1 {
				dot = length
				return nil, 0, true
			}
		}
		if unicode.IsDigit(r) {
			length++
			return nil, 0, true
		} else if length == 0 || dot == -1 || dot == length {
			return nil, 0, false
		} else {
			return &Token{Type: FLOAT, Value: string(currentText), Column: length}, 0, false
		}
	}
}
