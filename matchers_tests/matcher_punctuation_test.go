package matchers_tests

import (
	"github.com/JeffThomas/lexx/lexx"
	"github.com/JeffThomas/lexx/matchers"
	"testing"
)

////////////////////////////////
// PunctuationMatcher
////////////////////////////////

func TestLexxFindsPunctuation(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("!", []matchers.LexxMatcherInitialize{
		matchers.StartPunctuationMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Punctuation\n")
	} else if to.Value != "!" {
		t.Errorf("Found thing should be '!' not %s\n", to.Value)
	} else if to.Type != matchers.PUNCTUATION {
		t.Errorf("Found tokens should be type PUNCTUATION not %s\n", to.Type)
	}
}

func TestLexxDoesntFindWrongPunctuation(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("a", []matchers.LexxMatcherInitialize{
		matchers.StartPunctuationMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to.Type != matchers.UNDEFINED {
		t.Errorf("Lexx found wrong thing %s\n", to.Value)
	}
}

func TestLexxDoesntFindEmptyPunctuation(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("", []matchers.LexxMatcherInitialize{
		matchers.StartPunctuationMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find EOF\n")
	} else if to.Value != "EOF" {
		t.Errorf("Found thing should be 'EOF' not %s\n", to.Value)
	} else if to.Type != matchers.SYSTEM {
		t.Errorf("Found tokens should be type SYSTEM not %s\n", to.Type)
	}
}

func TestLexxFindsLeadingPunctuation(t *testing.T) {
	lexxx := lexx.BuildLexxWithString(".893a", []matchers.LexxMatcherInitialize{
		matchers.StartPunctuationMatcher,
		matchers.StartIntegerMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Punctuation\n")
	} else if to.Value != "." {
		t.Errorf("Found thing should be '.' not %s\n", to.Value)
	} else if to.Type != matchers.PUNCTUATION {
		t.Errorf("Found tokens should be type PUNCTUATION not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "893" {
		t.Errorf("Found thing should be '893' not %s\n", to.Value)
	} else if to.Type != matchers.INTEGER {
		t.Errorf("Found tokens should be type INTEGER not %s\n", to.Type)
	}
}

func TestLexxFindsMultiplePunctuations(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("( # @ ?", []matchers.LexxMatcherInitialize{
		matchers.StartWhitespaceMatcher,
		matchers.StartPunctuationMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Punctuation\n")
	} else if to.Value != "(" {
		t.Errorf("Found token should be '(' not %s\n", to.Value)
	} else if to.Type != matchers.PUNCTUATION {
		t.Errorf("Found tokens should be type PUNCTUATION not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " " {
		t.Errorf("Found token should be ' ' not %s\n", to.Value)
	} else if to.Type != matchers.WHITESPACE {
		t.Errorf("Found tokens should be type WHITESPACE not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Punctuation\n")
	} else if to.Value != "#" {
		t.Errorf("Found token should be '#' not %s\n", to.Value)
	} else if to.Type != matchers.PUNCTUATION {
		t.Errorf("Found tokens should be type PUNCTUATION not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " " {
		t.Errorf("Found token should be ' ' not %s\n", to.Value)
	} else if to.Type != matchers.WHITESPACE {
		t.Errorf("Found tokens should be type WHITESPACE not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Punctuation\n")
	} else if to.Value != "@" {
		t.Errorf("Found token should be '@' not %s\n", to.Value)
	} else if to.Type != matchers.PUNCTUATION {
		t.Errorf("Found tokens should be type PUNCTUATION not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " " {
		t.Errorf("Found token should be ' ' not %s\n", to.Value)
	} else if to.Type != matchers.WHITESPACE {
		t.Errorf("Found tokens should be type WHITESPACE not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Punctuation\n")
	} else if to.Value != "?" {
		t.Errorf("Found token should be '?' not %s\n", to.Value)
	} else if to.Type != matchers.PUNCTUATION {
		t.Errorf("Found tokens should be type PUNCTUATION not %s\n", to.Type)
	}
}
