package matchers

import (
	"errors"
	token_class "github.com/JeffThomas/lexx/token"
	"unicode"
)

func InitKeywordMatcher(newKeywords []string) func() func(r rune, currentText *string) MatcherResult {

	keywords := make([][]rune, len(newKeywords))
	for i, keyword := range newKeywords {
		keywords[i] = []rune(keyword)
	}

	current := make([][]rune, len(keywords))

	return func() func(r rune, currentText *string) MatcherResult {
		runeCount := 0
		found := ""
		for i, keyword := range keywords {
			current[i] = keyword
		}

		keywordCount := len(current)

		return func(r rune, currentText *string) MatcherResult {

			if r != 0 {
				for i, keyword := range current {
					if keyword == nil {
						continue
					}

					if runeCount == len(current[i]) {
						if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
							found = string(current[i])
						}
						current[i] = nil
						keywordCount--
					} else {
						if current[i][runeCount] != r {
							current[i] = nil
							keywordCount--
						}
					}
				}
			}

			if r == 0 || keywordCount == 0 {
				for i := range current {
					if runeCount == len(current[i]) {
						found = string(current[i])
						current[i] = nil
						keywordCount--
					}
				}
				if found != "" {
					return MatcherResult{
						Token:      &token_class.Token{Type: token_class.KEYWORD, Value: found, Column: runeCount},
						Err:        nil,
						Precedence: 1,
					}
				} else {
					return MatcherResult{
						Token:      nil,
						Err:        errors.New("not a keyword"),
						Precedence: 0,
					}
				}
			}

			runeCount++
			return MatcherResult{
				Token:      nil,
				Err:        nil,
				Precedence: 0,
			}
		}
	}
}
