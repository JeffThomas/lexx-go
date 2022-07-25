package matchers

func ConfigStringMatcher(useSingleQuote bool) LexxMatcherInitialize {
	singleQuote := useSingleQuote
	return func() LexxMatcherMatch {
		length := 0
		lines := 0
		done := false
		skip := false
		var startSymbol rune
		return func(r rune, currentText []rune) (token *Token, precedence int8, run bool) {

			if length == 0 && lines == 0 {
				if r != '"' && (!singleQuote || r != '\'') {
					return nil, 0, false
				}
				startSymbol = r
				length++
				return nil, 0, true
			}

			if r == 0 && !done {
				return nil, 0, false
			}

			if done {
				return &Token{Type: STRING, Value: string(currentText), Line: lines, Column: length}, 3, false
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
			if r == '\n' {
				lines++
				length = 0
			}
			return nil, 0, true
		}
	}
}
