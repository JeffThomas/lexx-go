package matchers_tests

import (
	"github.com/JeffThomas/lexx/lexx"
	"github.com/JeffThomas/lexx/matchers"
	"testing"
)

////////////////////////////////
// KeywordMatcher
////////////////////////////////

func TestLexxFindsKeyword(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("for", []matchers.LexxMatcherInitialize{
		matchers.ConfigKeywordMatcher([]string{"for"}),
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Keyword\n")
	} else if to.Value != "for" {
		t.Errorf("Found token should be 'for' not %s\n", to.Value)
	} else if to.Type != matchers.KEYWORD {
		t.Errorf("Found tokens should be type KEYWORD not %s\n", to.Type)
	}
}

func TestLexxDoesntFindWrongKeyword(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("far", []matchers.LexxMatcherInitialize{
		matchers.ConfigKeywordMatcher([]string{"for"}),
	})
	to, _ := lexxx.GetNextToken()
	if to.Type != matchers.UNDEFINED {
		t.Errorf("Lexx found wrong keyword %s\n", to.Value)
	}
}

func TestLexxDoesntFindPartOfKeyword(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("forward", []matchers.LexxMatcherInitialize{
		matchers.ConfigKeywordMatcher([]string{"for"}),
	})
	to, _ := lexxx.GetNextToken()
	if to.Type != matchers.UNDEFINED {
		t.Errorf("Lexx found wrong keyword %s\n", to.Value)
	}
}

func TestLexxFindsLongestKeyword(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("forward", []matchers.LexxMatcherInitialize{
		matchers.ConfigKeywordMatcher([]string{"for", "forward", "new"}),
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Keyword\n")
	} else if to.Value != "forward" {
		t.Errorf("Found token should be 'forward' not %s\n", to.Value)
	} else if to.Type != matchers.KEYWORD {
		t.Errorf("Found tokens should be type KEYWORD not %s\n", to.Type)
	}
}

func TestLexxFindsKeywordNotIdentifier(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("forward", []matchers.LexxMatcherInitialize{
		matchers.StartIdentifierMatcher,
		matchers.ConfigKeywordMatcher([]string{"for", "forward", "new"}),
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Keyword\n")
	} else if to.Value != "forward" {
		t.Errorf("Found token should be 'forward' not %s\n", to.Value)
	} else if to.Type != matchers.KEYWORD {
		t.Errorf("Found tokens should be type KEYWORD not %s\n", to.Type)
	}
}

func TestLexxFindsKeywordNotIdentifierOrWord(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("forward", []matchers.LexxMatcherInitialize{
		matchers.StartIdentifierMatcher,
		matchers.StartWordMatcher,
		matchers.ConfigKeywordMatcher([]string{"for", "forward", "new"}),
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Keyword\n")
	} else if to.Value != "forward" {
		t.Errorf("Found token should be 'forward' not %s\n", to.Value)
	} else if to.Type != matchers.KEYWORD {
		t.Errorf("Found tokens should be type KEYWORD not %s\n", to.Type)
	}
}

func TestLexxFindsMultipleKeywords(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("The quick brown fox", []matchers.LexxMatcherInitialize{
		matchers.StartWhitespaceMatcher,
		matchers.ConfigKeywordMatcher([]string{"The", "fox", "quick", "brown"}),
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Keyword\n")
	} else if to.Value != "The" {
		t.Errorf("Found token should be 'The' not %s\n", to.Value)
	} else if to.Type != matchers.KEYWORD {
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
		t.Errorf("Lexx did not find Keyword\n")
	} else if to.Value != "quick" {
		t.Errorf("Found token should be 'quick' not %s\n", to.Value)
	} else if to.Type != matchers.KEYWORD {
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
		t.Errorf("Lexx did not find Keyword\n")
	} else if to.Value != "brown" {
		t.Errorf("Found token should be 'brown' not %s\n", to.Value)
	} else if to.Type != matchers.KEYWORD {
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
		t.Errorf("Lexx did not find Keyword\n")
	} else if to.Value != "fox" {
		t.Errorf("Found token should be 'fox' not %s\n", to.Value)
	} else if to.Type != matchers.KEYWORD {
		t.Errorf("Found tokens should be type KEYWORD not %s\n", to.Type)
	}
}
