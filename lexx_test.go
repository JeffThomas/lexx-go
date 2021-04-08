package lexx

import (
	"bufio"
	"fmt"
	matchers_class "github.com/JeffThomas/lexx/matchers"
	token_class "github.com/JeffThomas/lexx/token"
	"strings"
	"testing"
)

func TestLexxNoMatchers(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("adfasdfasf"))
	lexxx := Lexx{}
	lexxx.Input = r
	to, err := lexxx.GetNextToken()

	if to != nil {
		t.Errorf("What could this be? %s\n", to.Value)
	}
	if err != nil {
		t.Errorf("This isn't an error situtation %s\n", err.Error())
	}
}

func TestLexxNullString(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(""))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartWordMatcher)
	to, err := lexxx.GetNextToken()

	if to == nil {
		t.Errorf("LexxToken should not be empty on EOF.\n")
	} else if to.Type != token_class.SYSTEM {
		t.Errorf("EOF token should be of type  SYSTEM: %s.\n", to.Type)
	}
	if err != nil {
		t.Errorf("This isn't an error situtation %s\n", err.Error())
	}
}

func TestLexxFindsAToken(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("Text"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartWordMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Word\n")
	} else if to.Value != "Text" {
		t.Errorf("Found word should be 'Text' not %s\n", to.Value)
	} else if to.Type != token_class.WORD {
		t.Errorf("Found word token should be type WORD not %s\n", to.Type)
	}
}

func TestLexxLineCountToken(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("\n \n\n \n "))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartWhitespaceMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Word\n")
	} else if to.Value != "\n \n\n \n " {
		t.Errorf("Found word should be 'Text' not %s\n", to.Value)
	} else if to.Type != token_class.WHITESPACE {
		t.Errorf("Found word token should be type WHITESPACE not %s\n", to.Type)
	}
	if to != nil && (lexxx.Line != 4 || lexxx.Column != 1) {
		t.Errorf("Line should be 4 and Column should be 1. Found: %d, %d", to.Line, to.Column)
	}
}

func TestLexxParsesMultiples(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("Text this	\nthing"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartWordMatcher)
	lexxx.AddMatcher(matchers_class.StartWhitespaceMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Word\n")
	} else if to.Value != "Text" {
		t.Errorf("Found word should be 'Text' not %s\n", to.Value)
	} else if to.Type != token_class.WORD {
		t.Errorf("Found word token should be type WORD not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 0) {
		t.Errorf("Line should be 0 and Column should be 0. Found: %d, %d", to.Line, to.Column)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find first Whitespace\n")
	} else if to.Value != " " {
		t.Errorf("Found word should be ' ' not %s\n", to.Value)
	} else if to.Type != token_class.WHITESPACE {
		t.Errorf("Found whitespace token should be type WHITESPACE not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 4) {
		t.Errorf("Line should be 0 and Column should be 4. Found: %d, %d", to.Line, to.Column)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Word\n")
	} else if to.Value != "this" {
		t.Errorf("Found word should be 'this' not %s\n", to.Value)
	} else if to.Type != token_class.WORD {
		t.Errorf("Found word token should be type WORD not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 5) {
		t.Errorf("Line should be 0 and Column should be 5. Found: %d, %d", to.Line, to.Column)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find second Whitespace\n")
	} else if to.Value != "	\n" {
		t.Errorf("Found word should be '	\n' not %s\n", to.Value)
	} else if to.Type != token_class.WHITESPACE {
		t.Errorf("Found whitespace token should be type WHITESPACE not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 9) {
		t.Errorf("Line should be 0 and Column should be 9. Found: %d, %d", to.Line, to.Column)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Word\n")
	} else if to.Value != "thing" {
		t.Errorf("Found word should be 'thing' not %s\n", to.Value)
	} else if to.Type != token_class.WORD {
		t.Errorf("Found word token should be type WORD not %s\n", to.Type)
	}
	if to != nil && (to.Line != 1 || to.Column != 0) {
		t.Errorf("Line should be 1 and Column should be 0. Found: %d, %d", to.Line, to.Column)
	}
}

func TestLexxRewind(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("!++=+"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.InitSymbolMatcher([]string{"!++"}))
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "!++" {
		t.Errorf("Found word should be '!++' not %s\n", to.Value)
	} else if to.Type != token_class.SYMBOL {
		t.Errorf("Found word token should be type SYMBOL not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 0) {
		t.Errorf("Line should be 0 and Column should be 0. Found: %d, %d", to.Line, to.Column)
	}
	lexxx.Rewind()
	lexxx.matchers = lexxx.matchers[0:0]
	lexxx.AddMatcher(matchers_class.InitSymbolMatcher([]string{"+", "!", "+="}))
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "!" {
		t.Errorf("Found word should be '!' not %s\n", to.Value)
	} else if to.Type != token_class.SYMBOL {
		t.Errorf("Found word token should be type SYMBOL not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 0) {
		t.Errorf("Line should be 0 and Column should be 0. Found: %d, %d", to.Line, to.Column)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "+" {
		t.Errorf("Found word should be '+' not %s\n", to.Value)
	} else if to.Type != token_class.SYMBOL {
		t.Errorf("Found word token should be type SYMBOL not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 1) {
		t.Errorf("Line should be 0 and Column should be 1. Found: %d, %d", to.Line, to.Column)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "+=" {
		t.Errorf("Found word should be '+=' not %s\n", to.Value)
	} else if to.Type != token_class.SYMBOL {
		t.Errorf("Found word token should be type SYMBOL not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 3) {
		t.Errorf("Line should be 0 and Column should be 3. Found: %d, %d", to.Line, to.Column)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "+" {
		t.Errorf("Found word should be '+' not %s\n", to.Value)
	} else if to.Type != token_class.SYMBOL {
		t.Errorf("Found word token should be type SYMBOL not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 5) {
		t.Errorf("Line should be 0 and Column should be 5. Found: %d, %d", to.Line, to.Column)
	}
}

////////////////////////////////
// SymbolMatcher
////////////////////////////////

func TestLexxFindsSymbol(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("!"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.InitSymbolMatcher([]string{"!"}))
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "!" {
		t.Errorf("Found word should be '!' not %s\n", to.Value)
	} else if to.Type != token_class.SYMBOL {
		t.Errorf("Found word token should be type SYMBOL not %s\n", to.Type)
	}
}

func TestLexxFindsSubSymbol(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("***** ***"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.InitSymbolMatcher([]string{"**", "*"}))
	lexxx.AddMatcher(matchers_class.StartWhitespaceMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "**" {
		t.Errorf("Found word should be '**' not %s\n", to.Value)
	} else if to.Type != token_class.SYMBOL {
		t.Errorf("Found word token should be type SYMBOL not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "**" {
		t.Errorf("Found word should be '**' not %s\n", to.Value)
	} else if to.Type != token_class.SYMBOL {
		t.Errorf("Found word token should be type SYMBOL not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "*" {
		t.Errorf("Found word should be '*' not %s\n", to.Value)
	} else if to.Type != token_class.SYMBOL {
		t.Errorf("Found word token should be type SYMBOL not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " " {
		t.Errorf("Found word should be ' ' not %s\n", to.Value)
	} else if to.Type != token_class.WHITESPACE {
		t.Errorf("Found word token should be type WHITESPACE not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "**" {
		t.Errorf("Found word should be '**' not %s\n", to.Value)
	} else if to.Type != token_class.SYMBOL {
		t.Errorf("Found word token should be type SYMBOL not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "*" {
		t.Errorf("Found word should be '*' not %s\n", to.Value)
	} else if to.Type != token_class.SYMBOL {
		t.Errorf("Found word token should be type SYMBOL not %s\n", to.Type)
	}
}

func TestLexxFindsLongestSymbol(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("+="))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.InitSymbolMatcher([]string{"+", "!", "+="}))
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "+=" {
		t.Errorf("Found word should be '+=' not %s\n", to.Value)
	} else if to.Type != token_class.SYMBOL {
		t.Errorf("Found word token should be type SYMBOL not %s\n", to.Type)
	}
}

func TestLexxFindsMultipleSymbols(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("!++=+"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.InitSymbolMatcher([]string{"+", "!", "+="}))
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "!" {
		t.Errorf("Found word should be '!' not %s\n", to.Value)
	} else if to.Type != token_class.SYMBOL {
		t.Errorf("Found word token should be type SYMBOL not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "+" {
		t.Errorf("Found word should be '+' not %s\n", to.Value)
	} else if to.Type != token_class.SYMBOL {
		t.Errorf("Found word token should be type SYMBOL not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "+=" {
		t.Errorf("Found word should be '+=' not %s\n", to.Value)
	} else if to.Type != token_class.SYMBOL {
		t.Errorf("Found word token should be type SYMBOL not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "+" {
		t.Errorf("Found word should be '+' not %s\n", to.Value)
	} else if to.Type != token_class.SYMBOL {
		t.Errorf("Found word token should be type SYMBOL not %s\n", to.Type)
	}
}

////////////////////////////////
// KeywordMatcher
////////////////////////////////

func TestLexxFindsKeyword(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("for"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.InitKeywordMatcher([]string{"for"}))
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Keyword\n")
	} else if to.Value != "for" {
		t.Errorf("Found word should be 'for' not %s\n", to.Value)
	} else if to.Type != token_class.KEYWORD {
		t.Errorf("Found word token should be type KEYWORD not %s\n", to.Type)
	}
}

func TestLexxDoesntFindWrongKeyword(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("far"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.InitKeywordMatcher([]string{"for"}))
	to, _ := lexxx.GetNextToken()
	if to != nil {
		t.Errorf("Lexx found wrong word %s\n", to.Value)
	}
}

func TestLexxDoesntFindPartOfKeyword(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("forward"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.InitKeywordMatcher([]string{"for"}))
	to, _ := lexxx.GetNextToken()
	if to != nil {
		t.Errorf("Lexx found wrong word %s\n", to.Value)
	}
}

func TestLexxFindsLongestKeyword(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("forward"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.InitKeywordMatcher([]string{"for", "forward", "new"}))
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Keyword\n")
	} else if to.Value != "forward" {
		t.Errorf("Found word should be 'forward' not %s\n", to.Value)
	} else if to.Type != token_class.KEYWORD {
		t.Errorf("Found word token should be type KEYWORD not %s\n", to.Type)
	}
}

func TestLexxFindsKeywordNotIdentifier(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("forward"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.InitKeywordMatcher([]string{"for", "forward", "new"}))
	lexxx.AddMatcher(matchers_class.StartIdentifierMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Keyword\n")
	} else if to.Value != "forward" {
		t.Errorf("Found word should be 'forward' not %s\n", to.Value)
	} else if to.Type != token_class.KEYWORD {
		t.Errorf("Found word token should be type KEYWORD not %s\n", to.Type)
	}
}

func TestLexxFindsMultipleKeywords(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("The quick brown fox"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartWhitespaceMatcher)
	lexxx.AddMatcher(matchers_class.InitKeywordMatcher([]string{"The", "fox", "quick", "brown"}))
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Keyword\n")
	} else if to.Value != "The" {
		t.Errorf("Found word should be 'The' not %s\n", to.Value)
	} else if to.Type != token_class.KEYWORD {
		t.Errorf("Found word token should be type KEYWORD not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " " {
		t.Errorf("Found word should be '+' not %s\n", to.Value)
	} else if to.Type != token_class.WHITESPACE {
		t.Errorf("Found word token should be type WHITESPACE not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Keyword\n")
	} else if to.Value != "quick" {
		t.Errorf("Found word should be 'quick' not %s\n", to.Value)
	} else if to.Type != token_class.KEYWORD {
		t.Errorf("Found word token should be type KEYWORD not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " " {
		t.Errorf("Found word should be '+' not %s\n", to.Value)
	} else if to.Type != token_class.WHITESPACE {
		t.Errorf("Found word token should be type WHITESPACE not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Keyword\n")
	} else if to.Value != "brown" {
		t.Errorf("Found word should be 'brown' not %s\n", to.Value)
	} else if to.Type != token_class.KEYWORD {
		t.Errorf("Found word token should be type KEYWORD not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " " {
		t.Errorf("Found word should be '+' not %s\n", to.Value)
	} else if to.Type != token_class.WHITESPACE {
		t.Errorf("Found word token should be type WHITESPACE not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Keyword\n")
	} else if to.Value != "fox" {
		t.Errorf("Found word should be 'fox' not %s\n", to.Value)
	} else if to.Type != token_class.KEYWORD {
		t.Errorf("Found word token should be type KEYWORD not %s\n", to.Type)
	}
}

////////////////////////////////
// IntegerMatcher
////////////////////////////////

func TestLexxFindsInteger(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("1"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartIntegerMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "1" {
		t.Errorf("Found thing should be '1' not %s\n", to.Value)
	} else if to.Type != token_class.INTEGER {
		t.Errorf("Found word token should be type INTEGER not %s\n", to.Type)
	}
}

func TestLexxDoesntFindWrongInteger(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("a"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartIntegerMatcher)
	to, _ := lexxx.GetNextToken()
	if to != nil {
		t.Errorf("Lexx found wrong thing %s\n", to.Value)
	}
}

func TestLexxFindsLeadingNumber(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("893a"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartIntegerMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "893" {
		t.Errorf("Found thing should be '893' not %s\n", to.Value)
	} else if to.Type != token_class.INTEGER {
		t.Errorf("Found token should be type INTEGER not %s\n", to.Type)
	}
}

func TestLexxFindsMultipleNumbers(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("436 33343 42 2"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartWhitespaceMatcher)
	lexxx.AddMatcher(matchers_class.StartIntegerMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "436" {
		t.Errorf("Found word should be '436' not %s\n", to.Value)
	} else if to.Type != token_class.INTEGER {
		t.Errorf("Found word token should be type INTEGER not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " " {
		t.Errorf("Found word should be ' ' not %s\n", to.Value)
	} else if to.Type != token_class.WHITESPACE {
		t.Errorf("Found word token should be type WHITESPACE not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "33343" {
		t.Errorf("Found word should be '33343' not %s\n", to.Value)
	} else if to.Type != token_class.INTEGER {
		t.Errorf("Found word token should be type INTEGER not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " " {
		t.Errorf("Found word should be ' ' not %s\n", to.Value)
	} else if to.Type != token_class.WHITESPACE {
		t.Errorf("Found word token should be type WHITESPACE not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "42" {
		t.Errorf("Found word should be '42' not %s\n", to.Value)
	} else if to.Type != token_class.INTEGER {
		t.Errorf("Found word token should be type INTEGER not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " " {
		t.Errorf("Found word should be ' ' not %s\n", to.Value)
	} else if to.Type != token_class.WHITESPACE {
		t.Errorf("Found word token should be type WHITESPACE not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "2" {
		t.Errorf("Found word should be '2' not %s\n", to.Value)
	} else if to.Type != token_class.INTEGER {
		t.Errorf("Found word token should be type INTEGER not %s\n", to.Type)
	}
}

////////////////////////////////
// FloatMatcher
////////////////////////////////

func TestLexxFindsFloat(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("1.0"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartFloatMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "1.0" {
		t.Errorf("Found thing should be '1.0' not %s\n", to.Value)
	} else if to.Type != token_class.FLOAT {
		t.Errorf("Found word token should be type FLOAT not %s\n", to.Type)
	}
}

func TestLexxFindsFloatLeadingDot(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(".234"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartFloatMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Float\n")
	} else if to.Value != ".234" {
		t.Errorf("Found thing should be '.234' not %s\n", to.Value)
	} else if to.Type != token_class.FLOAT {
		t.Errorf("Found word token should be type FLOAT not %s\n", to.Type)
	}
}

func TestLexxDoesntFindWrongFloat(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("a"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartFloatMatcher)
	to, _ := lexxx.GetNextToken()
	if to != nil {
		t.Errorf("Lexx found wrong thing %s\n", to.Value)
	}
}

func TestLexxDoesntFindWrongFloatInsteadOfInteger(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("1."))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartFloatMatcher)
	lexxx.AddMatcher(matchers_class.StartIntegerMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "1" {
		t.Errorf("Found thing should be '1' not %s\n", to.Value)
	} else if to.Type != token_class.INTEGER {
		t.Errorf("Found word token should be type INTEGER not %s\n", to.Type)
	}
}

func TestLexxDoesntFindWrongFloatInsteadOfInteger2(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("1. "))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartFloatMatcher)
	lexxx.AddMatcher(matchers_class.StartIntegerMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "1" {
		t.Errorf("Found thing should be '1' not %s\n", to.Value)
	} else if to.Type != token_class.INTEGER {
		t.Errorf("Found word token should be type INTEGER not %s\n", to.Type)
	}
}

func TestLexxFindsFloatNotInteger(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("1.1 "))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartFloatMatcher)
	lexxx.AddMatcher(matchers_class.StartIntegerMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Float\n")
	} else if to.Value != "1.1" {
		t.Errorf("Found thing should be '1.1' not %s\n", to.Value)
	} else if to.Type != token_class.FLOAT {
		t.Errorf("Found word token should be type FLOAT not %s\n", to.Type)
	}
}

func TestLexxFindsLeadingFloat(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("893.5a"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartFloatMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "893.5" {
		t.Errorf("Found thing should be '893.5' not %s\n", to.Value)
	} else if to.Type != token_class.FLOAT {
		t.Errorf("Found token should be type FLOAT not %s\n", to.Type)
	}
}

func TestLexxFindsMultipleFloats(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("436.2 33343.444 42.42 2.0001"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartWhitespaceMatcher)
	lexxx.AddMatcher(matchers_class.StartFloatMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "436.2" {
		t.Errorf("Found word should be '436.2' not %s\n", to.Value)
	} else if to.Type != token_class.FLOAT {
		t.Errorf("Found word token should be type INTEGER not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " " {
		t.Errorf("Found word should be ' ' not %s\n", to.Value)
	} else if to.Type != token_class.WHITESPACE {
		t.Errorf("Found word token should be type WHITESPACE not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "33343.444" {
		t.Errorf("Found word should be '33343.444' not %s\n", to.Value)
	} else if to.Type != token_class.FLOAT {
		t.Errorf("Found word token should be type INTEGER not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " " {
		t.Errorf("Found word should be ' ' not %s\n", to.Value)
	} else if to.Type != token_class.WHITESPACE {
		t.Errorf("Found word token should be type WHITESPACE not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "42.42" {
		t.Errorf("Found word should be '42.42' not %s\n", to.Value)
	} else if to.Type != token_class.FLOAT {
		t.Errorf("Found word token should be type INTEGER not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " " {
		t.Errorf("Found word should be ' ' not %s\n", to.Value)
	} else if to.Type != token_class.WHITESPACE {
		t.Errorf("Found word token should be type WHITESPACE not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "2.0001" {
		t.Errorf("Found word should be '2.0001' not %s\n", to.Value)
	} else if to.Type != token_class.FLOAT {
		t.Errorf("Found word token should be type INTEGER not %s\n", to.Type)
	}
}

////////////////////////////////
// StringMatcher
////////////////////////////////

func TestLexxDoesNotFindUnterminatedString(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("\"st"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartStringMatcher)
	to, err := lexxx.GetNextToken()
	fmt.Printf("%+v\n", to)
	if to.Type != token_class.SYSTEM {
		t.Errorf("EOF token should be of type  SYSTEM: %s.\n", to.Type)
	}
	if err != nil {
		t.Errorf("This isn't an error situtation %s\n", err.Error())
	}
}

func TestLexxDoesNotFindUnterminatedStringWithOtherQuote(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("\"st'"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartStringMatcher)
	to, err := lexxx.GetNextToken()

	if to.Type != token_class.SYSTEM {
		t.Errorf("EOF token should be of type  SYSTEM: %s.\n", to.Type)
	}
	if err != nil {
		t.Errorf("This isn't an error situtation %s\n", err.Error())
	}
}

func TestLexxFindsString(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("\"string\""))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartStringMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find String\n")
	} else if to.Value != "\"string\"" {
		t.Errorf("Found thing should be \"string\" not %s\n", to.Value)
	} else if to.Type != token_class.STRING {
		t.Errorf("Found token should be type FLOAT not %s\n", to.Type)
	}
}

func TestLexxFindsStringWithSubstring(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("\"st'r'ing\""))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartStringMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find String\n")
	} else if to.Value != "\"st'r'ing\"" {
		t.Errorf("Found thing should be \"st'r'ing\" not %s\n", to.Value)
	} else if to.Type != token_class.STRING {
		t.Errorf("Found token should be type FLOAT not %s\n", to.Type)
	}
}

func TestLexxFindsStringWithEscapedSubstring(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("\"st\\\"r\\\"ing\""))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartStringMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find String\n")
	} else if to.Value != "\"st\\\"r\\\"ing\"" {
		t.Errorf("Found thing should be \"st\\\"r\\\"ing\" not %s\n", to.Value)
	} else if to.Type != token_class.STRING {
		t.Errorf("Found token should be type STRING not %s\n", to.Type)
	}
}

////////////////////////////////
// IdentifierMatcher
////////////////////////////////

func TestLexxFindsIdentifier(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("identifier"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartIdentifierMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Identifier\n")
	} else if to.Value != "identifier" {
		t.Errorf("Found thing should be 'identifier' not %s\n", to.Value)
	} else if to.Type != token_class.IDENTIFIER {
		t.Errorf("Found token should be type IDENTIFIER not %s\n", to.Type)
	}
}

func TestLexxFindsIdentifierWithLeadingUnderscore(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("_identifier"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartIdentifierMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Identifier\n")
	} else if to.Value != "_identifier" {
		t.Errorf("Found thing should be '_identifier' not %s\n", to.Value)
	} else if to.Type != token_class.IDENTIFIER {
		t.Errorf("Found token should be type IDENTIFIER not %s\n", to.Type)
	}
}

func TestLexxFindsIdentifierWithLeadingAndEmbeddedUnderscore(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("__ident_ifi_er_"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartIdentifierMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Identifier\n")
	} else if to.Value != "__ident_ifi_er_" {
		t.Errorf("Found thing should be '__ident_ifi_er_' not %s\n", to.Value)
	} else if to.Type != token_class.IDENTIFIER {
		t.Errorf("Found token should be type IDENTIFIER not %s\n", to.Type)
	}
}

func TestLexyFindsIdentifierWithNumberAndEmbeddedUnderscore(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("__2ident_ifi_er_"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartIdentifierMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Identifier\n")
	} else if to.Value != "__2ident_ifi_er_" {
		t.Errorf("Found thing should be '__2ident_ifi_er_' not %s\n", to.Value)
	} else if to.Type != token_class.IDENTIFIER {
		t.Errorf("Found token should be type IDENTIFIER not %s\n", to.Type)
	}
}

func TestLexxFindsIdentifierNotKeyword(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("newThing"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.InitKeywordMatcher([]string{"new"}))
	lexxx.AddMatcher(matchers_class.StartIdentifierMatcher)
	to, _ := lexxx.GetNextToken()

	if to == nil {
		t.Errorf("Lexx did not find Identifier\n")
	} else if to.Value != "newThing" {
		t.Errorf("Found identifier should be 'newThing' not %s\n", to.Value)
	} else if to.Type != token_class.IDENTIFIER {
		t.Errorf("Found word token should be type IDENTIFIER not %s\n", to.Type)
	}
}

func TestLexxDoesNotFindIdentifierWithLeadingNumber(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("2ident_ifi_er_"))
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartIdentifierMatcher)
	to, _ := lexxx.GetNextToken()

	if to != nil {
		t.Errorf("Lexx found incorrect token, should be null but it's %s\n", to.Value)
	}
}
