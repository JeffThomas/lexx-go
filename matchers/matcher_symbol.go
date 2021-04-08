package matchers

import (
	"errors"
)

func InitSymbolMatcher(newSymbols []string) func() func(r rune, currentText *string) MatcherResult {

	symbols := make([][]rune, len(newSymbols))
	for i, symbol := range newSymbols {
		symbols[i] = []rune(symbol)
	}

	current := make([][]rune, len(symbols))

	return func() func(r rune, currentText *string) MatcherResult {
		runeCount := 0
		found := ""
		for i, symbol := range symbols {
			current[i] = symbol
		}

		symbolCount := len(current)

		return func(r rune, currentText *string) MatcherResult {
			if r == 0 || symbolCount == 0 {
				if found != "" {
					return MatcherResult{
						Token: &Token{Type: SYMBOL, Value: found, Column: runeCount},
						Err:   nil,
					}
				} else {
					return MatcherResult{
						Token: nil,
						Err:   errors.New("not a symbol"),
					}
				}
			}

			for i, symbol := range current {
				if symbol == nil {
					continue
				}

				match := current[i][runeCount] == r
				finished := runeCount+1 == len(current[i])
				if match && finished {
					found = string(current[i])
					current[i] = nil
					symbolCount--
				} else if !match || finished {
					current[i] = nil
					symbolCount--
				}
			}

			runeCount++
			return MatcherResult{
				Token: nil,
				Err:   nil,
			}

		}
	}
}
