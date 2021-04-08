package matchers

import (
	"errors"
	token_class "github.com/JeffThomas/lexx/token"
	"unicode"
)

func StartFloatMatcher() func(r rune, currentText *string) MatcherResult {
	length := 0
	dot := -1
	return func(r rune, currentText *string) MatcherResult {

		if r == 0 {
			if len(*currentText) > 0 && dot > -1 && length > dot {
				return MatcherResult{
					Token:      &token_class.Token{Type: token_class.FLOAT, Value: *currentText + "", Column: length},
					Err:        nil,
					Precedence: 0,
				}
			} else {
				return MatcherResult{
					Token:      nil,
					Err:        errors.New("not a float"),
					Precedence: 0,
				}
			}
		}

		if r == '.' {
			if dot > -1 && length == dot {
				return MatcherResult{
					Token:      nil,
					Err:        errors.New("not a float"),
					Precedence: 0,
				}
			}
			if dot == -1 {
				dot = length
				return MatcherResult{
					Token:      nil,
					Err:        nil,
					Precedence: 0,
				}
			}
		}
		if unicode.IsDigit(r) {
			length++
			return MatcherResult{
				Token:      nil,
				Err:        nil,
				Precedence: 0,
			}
		} else if len(*currentText) == 0 || dot == -1 || dot == length {
			return MatcherResult{
				Token:      nil,
				Err:        errors.New("not a float"),
				Precedence: 0,
			}
		} else {
			return MatcherResult{
				Token:      &token_class.Token{Type: token_class.FLOAT, Value: *currentText + "", Column: length},
				Err:        nil,
				Precedence: 0,
			}
		}
	}
}
