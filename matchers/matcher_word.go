package matchers

import (
	"errors"
	"unicode"
)

func StartWordMatcher() func(r rune, currentText *string) MatcherResult {
	length := 0
	return func(r rune, currentText *string) MatcherResult {

		if r == 0 {
			if len(*currentText) > 0 {
				return MatcherResult{
					Token:      &Token{Type: WORD, Value: *currentText + "", Column: length},
					Err:        nil,
					Precedence: 0,
				}
			} else {
				return MatcherResult{
					Token:      nil,
					Err:        errors.New("not a word"),
					Precedence: 0,
				}
			}
		}

		if unicode.IsLetter(r) || (length > 0 && unicode.IsDigit(r)) {
			length++
			return MatcherResult{
				Token:      nil,
				Err:        nil,
				Precedence: 0,
			}
		} else if len(*currentText) == 0 {
			return MatcherResult{
				Token:      nil,
				Err:        errors.New("not a word"),
				Precedence: 0,
			}
		} else {
			return MatcherResult{
				Token:      &Token{Type: WORD, Value: *currentText + "", Column: length},
				Err:        nil,
				Precedence: 0,
			}
		}
	}
}
