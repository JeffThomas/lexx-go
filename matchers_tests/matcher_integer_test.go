package matchers_tests

import (
	"github.com/JeffThomas/lexx/lexx"
	"github.com/JeffThomas/lexx/matchers"
	"testing"
)

////////////////////////////////
// IntegerMatcher
////////////////////////////////

func TestLexxFindsInteger(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("1", []matchers.LexxMatcherInitialize{
		matchers.StartIntegerMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "1" {
		t.Errorf("Found thing should be '1' not %s\n", to.Value)
	} else if to.Type != matchers.INTEGER {
		t.Errorf("Found tokens should be type INTEGER not %s\n", to.Type)
	}
}

func TestLexxDoesntFindWrongInteger(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("a", []matchers.LexxMatcherInitialize{
		matchers.StartIntegerMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to.Type != matchers.UNDEFINED {
		t.Errorf("Lexx found wrong thing %s\n", to.Value)
	}
}

func TestLexxDoesntFindEmpty(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("", []matchers.LexxMatcherInitialize{
		matchers.StartIntegerMatcher,
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

func TestLexxFindsLeadingNumber(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("893a", []matchers.LexxMatcherInitialize{
		matchers.StartIntegerMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "893" {
		t.Errorf("Found thing should be '893' not %s\n", to.Value)
	} else if to.Type != matchers.INTEGER {
		t.Errorf("Found tokens should be type INTEGER not %s\n", to.Type)
	}
}

func TestLexxFindsMultipleNumbers(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("436 33343 42 2", []matchers.LexxMatcherInitialize{
		matchers.StartWhitespaceMatcher,
		matchers.StartIntegerMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "436" {
		t.Errorf("Found token should be '436' not %s\n", to.Value)
	} else if to.Type != matchers.INTEGER {
		t.Errorf("Found tokens should be type INTEGER not %s\n", to.Type)
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
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "33343" {
		t.Errorf("Found token should be '33343' not %s\n", to.Value)
	} else if to.Type != matchers.INTEGER {
		t.Errorf("Found tokens should be type INTEGER not %s\n", to.Type)
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
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "42" {
		t.Errorf("Found token should be '42' not %s\n", to.Value)
	} else if to.Type != matchers.INTEGER {
		t.Errorf("Found tokens should be type INTEGER not %s\n", to.Type)
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
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "2" {
		t.Errorf("Found token should be '2' not %s\n", to.Value)
	} else if to.Type != matchers.INTEGER {
		t.Errorf("Found tokens should be type INTEGER not %s\n", to.Type)
	}
}
