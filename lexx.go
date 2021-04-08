package lexx

import (
	"bufio"
	matchers_class "github.com/JeffThomas/lexx/matchers"
	"log"
)

type Lexx struct {
	Input       *bufio.Reader
	matchers    []func() func(r rune, currentText *string) matchers_class.MatcherResult
	Token       *matchers_class.Token
	NextToken   *matchers_class.Token
	CurrentText string
	cache       []rune
	Line        int
	Column      int
	oldLine     int
	oldColumn   int
	precedence  int8
}

func (lexx *Lexx) Clone() *Lexx {
	return &Lexx{
		lexx.Input,
		lexx.matchers[:],
		lexx.Token,
		lexx.NextToken,
		lexx.CurrentText,
		lexx.cache,
		0,
		0,
		lexx.oldLine,
		lexx.oldColumn,
		lexx.precedence,
	}
}

func (lexx *Lexx) AddMatcher(newMatcher func() func(r rune, currentText *string) matchers_class.MatcherResult) {
	lexx.matchers = append(lexx.matchers, newMatcher)
}

func (lexx *Lexx) Rewind() {
	lexx.cache = append([]rune(lexx.Token.Value), lexx.cache...)
	lexx.Line = lexx.oldLine
	lexx.Column = lexx.oldColumn
	lexx.Token = nil
}

func (lexx *Lexx) PushToken() {
	lexx.NextToken = lexx.Token
	lexx.Token = nil
}

func (lexx *Lexx) GetNextToken() (*matchers_class.Token, error) {
	lexx.Token = nil

	if lexx.NextToken != nil {
		lexx.Token = lexx.NextToken
		lexx.NextToken = nil
		return lexx.Token, nil
	}

	runningMatchers := make([]func(r rune, currentText *string) matchers_class.MatcherResult, len(lexx.matchers))
	isMatcherRunning := make([]bool, len(lexx.matchers))

	for i, makeMatcher := range lexx.matchers {
		runningMatchers[i] = makeMatcher()
		isMatcherRunning[i] = true
	}

	isRunning := true
	lexx.CurrentText = ""

	for isRunning {
		isRunning = false

		var r rune
		var err error

		if len(lexx.cache) > 0 {
			r = lexx.cache[0]
			lexx.cache = lexx.cache[1:len(lexx.cache)]
		} else {
			r, _, err = lexx.Input.ReadRune()
			if err != nil {
				if err.Error() == "EOF" {
					// this isn't an error
					r = 0
				} else {
					log.Fatal(err.Error())
				}
			}
		}

		matched := false
		for i, matcher := range runningMatchers {
			if isMatcherRunning[i] {
				matchResult := matcher(r, &lexx.CurrentText)
				if matchResult.Token != nil {
					if !matched {
						lexx.Token = matchResult.Token
						lexx.precedence = matchResult.Precedence
						matched = true
					} else if lexx.precedence < matchResult.Precedence {
						lexx.Token = matchResult.Token
						lexx.precedence = matchResult.Precedence
					}
					isMatcherRunning[i] = false
				} else if matchResult.Err != nil {
					isMatcherRunning[i] = false
				} else {
					isRunning = true
				}
			}
		}

		if r == 0 {
			if lexx.Token == nil {
				lexx.Token = &matchers_class.Token{Type: matchers_class.SYSTEM, Value: "EOF"}
			}
		}

		currentLen := len(lexx.CurrentText)
		if !isRunning && lexx.Token != nil && lexx.Token.Type != matchers_class.SYSTEM {
			tokenLen := len(lexx.Token.Value)
			if tokenLen > 0 && tokenLen < currentLen {
				cache := append([]rune(lexx.CurrentText)[tokenLen:currentLen], r)
				lexx.cache = append(cache, lexx.cache...)
			} else {
				lexx.cache = append([]rune{r}, lexx.cache...)
			}
		}

		lexx.CurrentText = lexx.CurrentText + string(r)
	}
	if lexx.Token == nil {
		return nil, nil
	}
	lineAdvance := lexx.Token.Line
	columnAdvance := lexx.Token.Column
	lexx.oldLine = lexx.Line
	lexx.oldColumn = lexx.Column
	lexx.Token.Line = lexx.Line
	lexx.Token.Column = lexx.Column
	lexx.Line += lineAdvance
	if lineAdvance > 0 {
		lexx.Column = columnAdvance
	} else {
		lexx.Column += columnAdvance
	}
	return lexx.Token, nil
}
