package matchers_tests

import (
	"github.com/JeffThomas/lexx/lexx"
	"github.com/JeffThomas/lexx/matchers"
	"testing"
)

////////////////////////////////
// StringMatcher
////////////////////////////////

func TestLexxDoesNotFindUnterminatedString(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("\"st", []matchers.LexxMatcherInitialize{
		matchers.ConfigStringMatcher(true),
	})
	to, err := lexxx.GetNextToken()
	if to.Type != matchers.SYSTEM {
		t.Errorf("EOF tokens should be of type  SYSTEM: %s.\n", to.Type)
	}
	if err != nil {
		t.Errorf("This isn't an error situtation %s\n", err.Error())
	}
}

func TestLexxDoesNotFindNonString(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("a", []matchers.LexxMatcherInitialize{
		matchers.ConfigStringMatcher(true),
	})
	to, _ := lexxx.GetNextToken()
	if to.Type != matchers.UNDEFINED {
		t.Errorf("Lexx found wrong thing %s\n", to.Value)
	}
}

func TestLexxDoesNotFindEmptyCase(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("", []matchers.LexxMatcherInitialize{
		matchers.ConfigStringMatcher(true),
	})
	to, err := lexxx.GetNextToken()
	if to.Type != matchers.SYSTEM {
		t.Errorf("EOF tokens should be of type  SYSTEM: %s.\n", to.Type)
	}
	if err != nil {
		t.Errorf("This isn't an error situtation %s\n", err.Error())
	}
}

func TestLexxDoesNotFindUnterminatedStringWithOtherQuote(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("\"st'", []matchers.LexxMatcherInitialize{
		matchers.ConfigStringMatcher(true),
	})
	to, err := lexxx.GetNextToken()

	if to.Type != matchers.SYSTEM {
		t.Errorf("EOF tokens should be of type  SYSTEM: %s.\n", to.Type)
	}
	if err != nil {
		t.Errorf("This isn't an error situtation %s\n", err.Error())
	}
}

func TestLexxFindsString(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("\"string\"", []matchers.LexxMatcherInitialize{
		matchers.ConfigStringMatcher(true),
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find String\n")
	} else if to.Value != "\"string\"" {
		t.Errorf("Found thing should be \"string\" not %s\n", to.Value)
	} else if to.Type != matchers.STRING {
		t.Errorf("Found tokens should be type FLOAT not %s\n", to.Type)
	}
}

func TestLexxFindsStringWithSubstring(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("\"st'r'ing\"", []matchers.LexxMatcherInitialize{
		matchers.ConfigStringMatcher(true),
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find String\n")
	} else if to.Value != "\"st'r'ing\"" {
		t.Errorf("Found thing should be \"st'r'ing\" not %s\n", to.Value)
	} else if to.Type != matchers.STRING {
		t.Errorf("Found tokens should be type FLOAT not %s\n", to.Type)
	}
}

func TestLexxFindsStringWithEscapedSubstring(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("\"st\\\"r\\\"ing\"", []matchers.LexxMatcherInitialize{
		matchers.ConfigStringMatcher(true),
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find String\n")
	} else if to.Value != "\"st\\\"r\\\"ing\"" {
		t.Errorf("Found thing should be \"st\\\"r\\\"ing\" not %s\n", to.Value)
	} else if to.Type != matchers.STRING {
		t.Errorf("Found tokens should be type STRING not %s\n", to.Type)
	}
}

func TestLexxFindsStringWithPunctuation(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("\"penny\ndreadful\"", []matchers.LexxMatcherInitialize{
		matchers.StartWordMatcher,
		matchers.ConfigStringMatcher(false),
		matchers.StartWhitespaceMatcher,
		matchers.StartPunctuationMatcher,
		matchers.StartSymbolicMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find String\n")
	} else if to.Value != "\"penny\ndreadful\"" {
		t.Errorf("Found thing should be \"penny\ndreadful\" not %s\n", to.Value)
	} else if to.Type != matchers.STRING {
		t.Errorf("Found tokens should be type STRING not %s\n", to.Type)
	}
}

func TestLexxFailedMatchContinues(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("\"penny\ndreadful", []matchers.LexxMatcherInitialize{
		matchers.StartWordMatcher,
		matchers.ConfigStringMatcher(false),
		matchers.StartWhitespaceMatcher,
		matchers.StartPunctuationMatcher,
		matchers.StartSymbolicMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Punctuation\n")
	} else if to.Value != "\"" {
		t.Errorf("Found thing should be \" not %s\n", to.Value)
	} else if to.Type != matchers.PUNCTUATION {
		t.Errorf("Found tokens should be type PUNCTUATION not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find word\n")
	} else if to.Value != "penny" {
		t.Errorf("Found thing should be \"penny\" not %s\n", to.Value)
	} else if to.Type != matchers.WORD {
		t.Errorf("Found tokens should be type WORD not %s\n", to.Type)
	}
}
