package matchers

func ConfigOperatorMatcher(newOperators []string) LexxMatcherInitialize {
	operators := make([][]rune, len(newOperators))
	lengths := make([]int, len(newOperators))
	for i, operator := range newOperators {
		operators[i] = []rune(operator)
		lengths[i] = len(operator)
	}

	running := make([]bool, len(operators))

	return func() LexxMatcherMatch {
		runeCount := 0
		found := ""
		for i := range operators {
			running[i] = true
		}

		operatorCount := len(operators)

		return func(r rune, currentText []rune) (token *Token, precedence int8, run bool) {
			if r != 0 {

				for i, operator := range operators {
					if !running[i] {
						continue
					}

					match := operator[runeCount] == r
					finished := runeCount+1 == lengths[i]
					if match && finished {
						found = string(operators[i])
						running[i] = false
						operatorCount--
					} else if !match || finished {
						running[i] = false
						operatorCount--
					}
				}
			}
			if r == 0 || operatorCount == 0 {
				if found != "" {
					return &Token{Type: OPERATOR, Value: found, Column: runeCount + 1}, 2, false
				} else {
					return nil, 0, false
				}
			}
			runeCount++
			return nil, 0, true
		}
	}
}
