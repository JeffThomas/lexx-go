package matchers_tests

import (
	"github.com/JeffThomas/lexx/lexx"
	"github.com/JeffThomas/lexx/matchers"
	"testing"
)

////////////////////////////////
// WordMatcher
////////////////////////////////

func TestLexxFindsWord(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("for", []matchers.LexxMatcherInitialize{
		matchers.StartWordMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find word\n")
	} else if to.Value != "for" {
		t.Errorf("Found token should be 'for' not %s\n", to.Value)
	} else if to.Type != matchers.WORD {
		t.Errorf("Found tokens should be type WORD not %s\n", to.Type)
	}
}

func TestLexxDoesntFindNumber(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("342", []matchers.LexxMatcherInitialize{
		matchers.StartWordMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to.Type != matchers.UNDEFINED {
		t.Errorf("Lexx found wrong word %s\n", to.Value)
	}
}

func TestLexxDoesntFindEmptyEOF(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("", []matchers.LexxMatcherInitialize{
		matchers.StartWordMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find EOF\n")
	} else if to.Value != "EOF" {
		t.Errorf("Found token should be 'EOF' not %s\n", to.Value)
	} else if to.Type != matchers.SYSTEM {
		t.Errorf("Found tokens should be type SYSTEM not %s\n", to.Type)
	}
}

func TestLexxFindsWordWithNumber(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("four4", []matchers.LexxMatcherInitialize{
		matchers.StartWordMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find word\n")
	} else if to.Value != "four4" {
		t.Errorf("Found token should be 'four4' not %s\n", to.Value)
	} else if to.Type != matchers.WORD {
		t.Errorf("Found tokens should be type WORD not %s\n", to.Type)
	}
}

func TestLexxDoesNotFindNumberWithWord(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("4four", []matchers.LexxMatcherInitialize{
		matchers.StartWordMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to.Type != matchers.UNDEFINED {
		t.Errorf("Lexx found wrong word %s\n", to.Value)
	}
}

func TestLexxFindsMultipleWords(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("The quick brown fox", []matchers.LexxMatcherInitialize{
		matchers.StartWordMatcher,
		matchers.StartWhitespaceMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find word\n")
	} else if to.Value != "The" {
		t.Errorf("Found token should be 'The' not %s\n", to.Value)
	} else if to.Type != matchers.WORD {
		t.Errorf("Found tokens should be type KEYWORD not %s\n", to.Type)
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
		t.Errorf("Lexx did not find word\n")
	} else if to.Value != "quick" {
		t.Errorf("Found token should be 'quick' not %s\n", to.Value)
	} else if to.Type != matchers.WORD {
		t.Errorf("Found tokens should be type KEYWORD not %s\n", to.Type)
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
		t.Errorf("Lexx did not find word\n")
	} else if to.Value != "brown" {
		t.Errorf("Found token should be 'brown' not %s\n", to.Value)
	} else if to.Type != matchers.WORD {
		t.Errorf("Found tokens should be type KEYWORD not %s\n", to.Type)
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
		t.Errorf("Lexx did not find word\n")
	} else if to.Value != "fox" {
		t.Errorf("Found token should be 'fox' not %s\n", to.Value)
	} else if to.Type != matchers.WORD {
		t.Errorf("Found tokens should be type KEYWORD not %s\n", to.Type)
	}
}
