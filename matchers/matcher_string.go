package matchers

import (
	"errors"
	token_class "github.com/JeffThomas/lexx/token"
)

func StartStringMatcher() func(r rune, currentText *string) MatcherResult {
	length := 0
	done := false
	skip := false
	var startSymbol rune
	return func(r rune, currentText *string) MatcherResult {

		if length == 0 {
			if r != '"' && r != '\'' {
				return MatcherResult{
					Token:      nil,
					Err:        errors.New("not a string"),
					Precedence: 0,
				}
			}
			startSymbol = r
			length++
			return MatcherResult{
				Token:      nil,
				Err:        nil,
				Precedence: 0,
			}
		}

		if r == 0 && !done {
			return MatcherResult{
				Token:      nil,
				Err:        errors.New("not a string"),
				Precedence: 0,
			}
		}

		if done {
			return MatcherResult{
				Token: &token_class.Token{Type: token_class.STRING, Value: *currentText + "", Column: length},
				Err:   nil,
			}
		}

		if !skip && r == startSymbol {
			done = true
		}

		if r == '\\' {
			skip = true
		} else {
			skip = false
		}

		length++
		return MatcherResult{
			Token:      nil,
			Err:        nil,
			Precedence: 0,
		}
	}
}
