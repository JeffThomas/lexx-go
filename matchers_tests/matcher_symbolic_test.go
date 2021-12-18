package matchers_tests

import (
	"github.com/JeffThomas/lexx/lexx"
	"github.com/JeffThomas/lexx/matchers"
	"testing"
)

////////////////////////////////
// SymbolicMatcher
////////////////////////////////

func TestLexxFindsSymbol(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("=", []matchers.LexxMatcherInitialize{
		matchers.StartSymbolicMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbol\n")
	} else if to.Value != "=" {
		t.Errorf("Found thing should be '=' not %s\n", to.Value)
	} else if to.Type != matchers.SYMBOLIC {
		t.Errorf("Found tokens should be type SYMBOLIC not %s\n", to.Type)
	}
}

func TestLexxDoesntFindWrongSymbol(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("a", []matchers.LexxMatcherInitialize{
		matchers.StartSymbolicMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to.Type != matchers.UNDEFINED {
		t.Errorf("Lexx found wrong thing %s\n", to.Value)
	}
}

func TestLexxDoesntFindEmptySymbol(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("", []matchers.LexxMatcherInitialize{
		matchers.StartSymbolicMatcher,
	})
	lexxx.AddMatcher(matchers.StartSymbolicMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find EOF\n")
	} else if to.Value != "EOF" {
		t.Errorf("Found thing should be 'EOF' not %s\n", to.Value)
	} else if to.Type != matchers.SYSTEM {
		t.Errorf("Found tokens should be type SYSTEM not %s\n", to.Type)
	}
}

func TestLexxFindsLeadingSymbol(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("=893a", []matchers.LexxMatcherInitialize{
		matchers.StartSymbolicMatcher,
		matchers.StartIntegerMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbol\n")
	} else if to.Value != "=" {
		t.Errorf("Found thing should be '=' not %s\n", to.Value)
	} else if to.Type != matchers.SYMBOLIC {
		t.Errorf("Found tokens should be type SYMBOLIC not %s\n", to.Type)
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

func TestLexxFindsMultipleSymbols(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("= + ~ ^", []matchers.LexxMatcherInitialize{
		matchers.StartSymbolicMatcher,
		matchers.StartWhitespaceMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbol\n")
	} else if to.Value != "=" {
		t.Errorf("Found token should be '=' not %s\n", to.Value)
	} else if to.Type != matchers.SYMBOLIC {
		t.Errorf("Found tokens should be type SYMBOLIC not %s\n", to.Type)
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
		t.Errorf("Lexx did not find Symbol\n")
	} else if to.Value != "+" {
		t.Errorf("Found token should be '+' not %s\n", to.Value)
	} else if to.Type != matchers.SYMBOLIC {
		t.Errorf("Found tokens should be type SYMBOLIC not %s\n", to.Type)
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
		t.Errorf("Lexx did not find Symbol\n")
	} else if to.Value != "~" {
		t.Errorf("Found token should be '~' not %s\n", to.Value)
	} else if to.Type != matchers.SYMBOLIC {
		t.Errorf("Found tokens should be type SYMBOLIC not %s\n", to.Type)
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
		t.Errorf("Lexx did not find Symbol\n")
	} else if to.Value != "^" {
		t.Errorf("Found token should be '^' not %s\n", to.Value)
	} else if to.Type != matchers.SYMBOLIC {
		t.Errorf("Found tokens should be type SYMBOLIC not %s\n", to.Type)
	}
}
