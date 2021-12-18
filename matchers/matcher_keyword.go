package matchers

import (
	"unicode"
)

func ConfigKeywordMatcher(newKeywords []string) LexxMatcherInitialize {
	keywords := make([][]rune, len(newKeywords))
	for i, keyword := range newKeywords {
		keywords[i] = []rune(keyword)
	}

	current := make([][]rune, len(keywords))

	return func() LexxMatcherMatch {
		runeCount := 0
		found := ""
		for i, keyword := range keywords {
			current[i] = keyword
		}

		keywordCount := len(current)

		return func(r rune, currentText []rune) (token *Token, precedence int8, run bool) {

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
					return &Token{Type: KEYWORD, Value: found, Column: runeCount}, 10, false
				} else {
					return nil, 0, false
				}
			}

			runeCount++
			return nil, 0, true
		}
	}
}
