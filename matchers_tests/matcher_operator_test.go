package matchers_tests

import (
	"github.com/JeffThomas/lexx/lexx"
	"github.com/JeffThomas/lexx/matchers"
	"testing"
)

////////////////////////////////
// OperatorMatcher
////////////////////////////////

func TestLexxFindsOperator(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("!", []matchers.LexxMatcherInitialize{
		matchers.ConfigOperatorMatcher([]string{"!"}),
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Operators\n")
	} else if to.Value != "!" {
		t.Errorf("Found token should be '!' not %s\n", to.Value)
	} else if to.Type != matchers.OPERATOR {
		t.Errorf("Found tokens should be type OPERATOR not %s\n", to.Type)
	}
}

func TestLexxFindsSubOperator(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("***** ***", []matchers.LexxMatcherInitialize{
		matchers.ConfigOperatorMatcher([]string{"**", "*"}),
		matchers.StartWhitespaceMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Operators\n")
	} else if to.Value != "**" {
		t.Errorf("Found token should be '**' not %s\n", to.Value)
	} else if to.Type != matchers.OPERATOR {
		t.Errorf("Found tokens should be type OPERATOR not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Operators\n")
	} else if to.Value != "**" {
		t.Errorf("Found token should be '**' not %s\n", to.Value)
	} else if to.Type != matchers.OPERATOR {
		t.Errorf("Found tokens should be type OPERATOR not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Operators\n")
	} else if to.Value != "*" {
		t.Errorf("Found token should be '*' not %s\n", to.Value)
	} else if to.Type != matchers.OPERATOR {
		t.Errorf("Found tokens should be type OPERATOR not %s\n", to.Type)
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
		t.Errorf("Lexx did not find Operators\n")
	} else if to.Value != "**" {
		t.Errorf("Found token should be '**' not %s\n", to.Value)
	} else if to.Type != matchers.OPERATOR {
		t.Errorf("Found tokens should be type OPERATOR not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Operators\n")
	} else if to.Value != "*" {
		t.Errorf("Found token should be '*' not %s\n", to.Value)
	} else if to.Type != matchers.OPERATOR {
		t.Errorf("Found tokens should be type OPERATOR not %s\n", to.Type)
	}
}

func TestLexxFindsLongestOperator(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("+=", []matchers.LexxMatcherInitialize{
		matchers.ConfigOperatorMatcher([]string{"+", "!", "+="}),
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Operators\n")
	} else if to.Value != "+=" {
		t.Errorf("Found token should be '+=' not %s\n", to.Value)
	} else if to.Type != matchers.OPERATOR {
		t.Errorf("Found tokens should be type OPERATOR not %s\n", to.Type)
	}
}

func TestLexxFindsMultipleOperators(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("!++=+", []matchers.LexxMatcherInitialize{
		matchers.ConfigOperatorMatcher([]string{"+", "!", "+="}),
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Operators\n")
	} else if to.Value != "!" {
		t.Errorf("Found token should be '!' not %s\n", to.Value)
	} else if to.Type != matchers.OPERATOR {
		t.Errorf("Found tokens should be type OPERATOR not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Operators\n")
	} else if to.Value != "+" {
		t.Errorf("Found token should be '+' not %s\n", to.Value)
	} else if to.Type != matchers.OPERATOR {
		t.Errorf("Found tokens should be type OPERATOR not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Operators\n")
	} else if to.Value != "+=" {
		t.Errorf("Found token should be '+=' not %s\n", to.Value)
	} else if to.Type != matchers.OPERATOR {
		t.Errorf("Found tokens should be type OPERATOR not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Operators\n")
	} else if to.Value != "+" {
		t.Errorf("Found token should be '+' not %s\n", to.Value)
	} else if to.Type != matchers.OPERATOR {
		t.Errorf("Found tokens should be type OPERATOR not %s\n", to.Type)
	}
}
