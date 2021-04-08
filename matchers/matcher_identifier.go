package matchers

import (
	"errors"
	"unicode"
)

func StartIdentifierMatcher() func(r rune, currentText *string) MatcherResult {
	length := 0
	return func(r rune, currentText *string) MatcherResult {

		if length == 0 {
			if unicode.IsDigit(r) {
				return MatcherResult{
					Token:      nil,
					Err:        errors.New("not an identifier"),
					Precedence: 0,
				}
			}

			if r == '_' || unicode.IsLetter(r) {
				length++
				return MatcherResult{
					Token:      nil,
					Err:        nil,
					Precedence: 0,
				}
			}
		}

		if r == 0 {
			if len(*currentText) > 0 {
				return MatcherResult{
					Token:      &Token{Type: IDENTIFIER, Value: *currentText + "", Column: length},
					Err:        nil,
					Precedence: 0,
				}
			} else {
				return MatcherResult{
					Token:      nil,
					Err:        errors.New("not an identifier"),
					Precedence: 0,
				}
			}
		}

		if r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r) {
			length++
			return MatcherResult{
				Token:      nil,
				Err:        nil,
				Precedence: 0,
			}
		}

		if len(*currentText) > 0 {
			return MatcherResult{
				Token: &Token{Type: IDENTIFIER, Value: *currentText + "", Column: length},
				Err:   nil,
			}
		} else {
			return MatcherResult{
				Token:      nil,
				Err:        errors.New("not an identifier"),
				Precedence: 0,
			}
		}
	}
}
