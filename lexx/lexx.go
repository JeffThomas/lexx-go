// Package lexx contains the main lexx object and tests.
package lexx

import (
	"bufio"
	"github.com/JeffThomas/lexx/matchers"
	"log"
	"strings"
)

// StateLexx contains the state of the lexer for a single match.
// A linked list of them is maintained as a history and is used by the rewind function.
type StateLexx struct {
	// CurrentText contains the text that was pulled from the input to try and match.
	// If the match fails this can be used to diagnose the failure. If the match succeeds any remaining
	// text left in here will be pushed into a cache for use when GetToken is next called
	CurrentText []rune
	// Line is the line this match was made at.
	Line int
	// Column is the column where this match was made.
	Column int
	// LineNext is the line on which this match ended.
	LineNext int
	// ColumnNext is the column on which this match ended.
	ColumnNext int
	// Token is the matched token.
	Token *matchers.Token
	// Previous forms the linked list history of matches.
	Previous *StateLexx
}

// Lexx is the main object of the lexer.
// To use Lexx you instantiate a Lexx and initialize the Matchers field with the LexxMatcherInitialize methods you wish to use
type Lexx struct {
	// Matchers is an array of matcher initializing functions that are called at the beginning of a match.
	Matchers []matchers.LexxMatcherInitialize
	// Token is the most recently matched Token.
	Token *matchers.Token
	// NextToken is a one time cache, any token stored in here will be returned the next time
	// GetToken() is called and this will be set to nil.
	NextToken *matchers.Token
	// State is the root of the StateLexx struct history of matches when a match is made a new StateLexx gets added here.
	State *StateLexx
	// precedence is the current running precedence.
	precedence int8

	// input is the input buffer for the lexer. Runes are pulled out of it one at a time and are fed to the matchers.
	input *bufio.Reader
	// Cache is where input over-reads are stored for use on the next match. It's also used by the Rewind function.
	cache []rune
	// runningMatchers is the array of currently running matchers during a match.
	runningMatchers []matchers.LexxMatcherMatch
	// isMatcherRunning is an array set to true when the matchers are running, when a matcher finds a match or fails its
	// corresponding boolean in this array is set to false.
	isMatcherRunning []bool
}

// SetInput sets the input for the lexer and resets the state, any history will be lost.
func (lexx *Lexx) SetInput(input *bufio.Reader) {
	lexx.input = input
	lexx.cache = lexx.cache[:0]
	lexx.NextToken = nil
	lexx.Token = nil
	lexx.precedence = 0
	lexx.State = &StateLexx{}
}

// SetStringInput sets the input for the lexer using a string. It just creates a Reader and then calls SetInput.
func (lexx *Lexx) SetStringInput(input string) {
	lexx.SetInput(bufio.NewReader(strings.NewReader(input)))
}

// AddMatcher adds a matcher to the array of matchers
func (lexx *Lexx) AddMatcher(newMatcher matchers.LexxMatcherInitialize) {
	lexx.Matchers = append(lexx.Matchers, newMatcher)
	lexx.runningMatchers = make([]matchers.LexxMatcherMatch, len(lexx.Matchers))
	lexx.isMatcherRunning = make([]bool, len(lexx.Matchers))
}

// AddMatchers adds an array of matchers to the array of matchers
func (lexx *Lexx) AddMatchers(matchersToAdd []matchers.LexxMatcherInitialize) {
	for _, matcher := range matchersToAdd {
		lexx.Matchers = append(lexx.Matchers, matcher)
	}
	lexx.runningMatchers = make([]matchers.LexxMatcherMatch, len(lexx.Matchers))
	lexx.isMatcherRunning = make([]bool, len(lexx.Matchers))
}

// Rewind undoes a match by pulling a StateLexx off the history and using it to reset the state of Lexx. It should
// be able to rewind all the matches previously made. Note that the previously used inputs are cached, so modifying
// lexx.input after rewinding will not work. However, you can change the matchers.
func (lexx *Lexx) Rewind() {
	if lexx.State.Token == nil {
		return
	}
	rewind := len(lexx.State.Token.Value) - 1
	for rewind > -1 {
		lexx.cache = append(lexx.cache, rune(lexx.State.Token.Value[rewind]))
		rewind -= 1
	}
	lexx.State = lexx.State.Previous
	lexx.Token = lexx.State.Token
}

// PushToken stores a single token that will be returned the next time GetNextToken is called. Note that this will not
// affect the StateLexx history in any way.
func (lexx *Lexx) PushToken() {
	lexx.NextToken = lexx.Token
	lexx.Token = nil
}

// GetNextToken Returns the next Token pulled from the input. If a match isn't made it will return an Undefined token.
func (lexx *Lexx) GetNextToken() (*matchers.Token, error) {
	lexx.Token = nil

	if lexx.NextToken != nil {
		lexx.Token = lexx.NextToken
		lexx.NextToken = nil
		return lexx.Token, nil
	}

	runningMatchers := lexx.runningMatchers
	isMatcherRunning := lexx.isMatcherRunning

	for i, makeMatcher := range lexx.Matchers {
		runningMatchers[i] = makeMatcher()
		isMatcherRunning[i] = true
	}

	isRunning := true
	currentText := make([]rune, 100)
	currentText = currentText[:0]
	var token *matchers.Token

	for isRunning {
		isRunning = false

		var r rune
		var err error

		if len(lexx.cache) > 0 {
			r, lexx.cache = lexx.cache[len(lexx.cache)-1], lexx.cache[:len(lexx.cache)-1]
		} else {
			r, _, err = lexx.input.ReadRune()
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
				t, precedence, running := matcher(r, currentText)
				isMatcherRunning[i] = running
				if !isRunning {
					isRunning = running
				}
				if t != nil {
					if !matched {
						token = t
						lexx.precedence = precedence
						matched = true
					} else if lexx.precedence < precedence {
						token = t
						lexx.precedence = precedence
					}
				}
			}
		}

		if r == 0 {
			if token == nil {
				token = &matchers.Token{Type: matchers.SYSTEM, Value: "EOF"}
			}
		}
		currentText = append(currentText, r)
	}

	if token == nil {
		token = &matchers.Token{Type: matchers.UNDEFINED, Value: string(currentText[0]), Column: 1}
	}

	rewind := len(currentText) - 1
	length := len(token.Value) - 1
	for rewind > length {
		lexx.cache = append(lexx.cache, currentText[rewind])
		rewind -= 1
	}
	lineAdvance := token.Line
	columnAdvance := token.Column
	token.Line = lexx.State.LineNext
	token.Column = lexx.State.ColumnNext
	lineNext := token.Line + lineAdvance
	columnNext := 0
	if lineAdvance > 0 {
		columnNext = columnAdvance
	} else {
		columnNext = token.Column + columnAdvance
	}
	state := StateLexx{
		CurrentText: currentText,
		Line:        lexx.State.Line,
		Column:      lexx.State.Column,
		LineNext:    lineNext,
		ColumnNext:  columnNext,
		Token:       token,
		Previous:    lexx.State,
	}
	lexx.State = &state
	lexx.Token = token
	return token, nil
}

// BuildLexxWithString creates a Lexx while setting the input and list of matchers at the same time.
func BuildLexxWithString(input string, matchersToAdd []matchers.LexxMatcherInitialize) *Lexx {
	r := bufio.NewReader(strings.NewReader(input))
	l := Lexx{input: r, State: &StateLexx{}}
	l.AddMatchers(matchersToAdd)
	return &l
}

// BuildLexxWithReader creates a Lexx while setting the input and list of matchers at the same time.
func BuildLexxWithReader(input *bufio.Reader, matchersToAdd []matchers.LexxMatcherInitialize) *Lexx {
	l := Lexx{input: input, State: &StateLexx{}}
	l.AddMatchers(matchersToAdd)
	return &l
}

// NewLexx creats a Lexx and sets the matchers
func NewLexx(matchersToAdd []matchers.LexxMatcherInitialize) *Lexx {
	l := Lexx{State: &StateLexx{}}
	l.AddMatchers(matchersToAdd)
	return &l
}

// NewDefaultLexx creates a Lexx with a default set of matchers which should tokenize standard text files just fine.
func NewDefaultLexx() *Lexx {
	l := NewLexx([]matchers.LexxMatcherInitialize{
		matchers.StartWordMatcher,
		matchers.StartIntegerMatcher,
		matchers.StartFloatMatcher,
		matchers.StartWhitespaceMatcher,
		matchers.StartPunctuationMatcher,
		matchers.StartSymbolicMatcher,
	})
	return l
}
