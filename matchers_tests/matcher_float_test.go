package matchers_tests

import (
	"github.com/JeffThomas/lexx/lexx"
	"github.com/JeffThomas/lexx/matchers"
	"testing"
)

////////////////////////////////
// FloatMatcher
////////////////////////////////

func TestLexxFindsFloat(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("1.0", []matchers.LexxMatcherInitialize{
		matchers.StartFloatMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "1.0" {
		t.Errorf("Found thing should be '1.0' not %s\n", to.Value)
	} else if to.Type != matchers.FLOAT {
		t.Errorf("Found tokens should be type FLOAT not %s\n", to.Type)
	}
}

func TestLexxFindsFloatLeadingDot(t *testing.T) {
	lexxx := lexx.BuildLexxWithString(".234", []matchers.LexxMatcherInitialize{
		matchers.StartFloatMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Float\n")
	} else if to.Value != ".234" {
		t.Errorf("Found thing should be '.234' not %s\n", to.Value)
	} else if to.Type != matchers.FLOAT {
		t.Errorf("Found tokens should be type FLOAT not %s\n", to.Type)
	}
}

func TestLexxDoesntFindWrongFloat(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("a", []matchers.LexxMatcherInitialize{
		matchers.StartFloatMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to.Type != matchers.UNDEFINED {
		t.Errorf("Lexx found wrong thing %s\n", to.Value)
	}
}

func TestLexxDoesntFindWrongFloatInsteadOfInteger(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("1.", []matchers.LexxMatcherInitialize{
		matchers.StartFloatMatcher,
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

func TestLexxDoesntFindWrongFloatInsteadOfInteger2(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("1. ", []matchers.LexxMatcherInitialize{
		matchers.StartFloatMatcher,
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

func TestLexxFindsFloatNotInteger(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("1.1 ", []matchers.LexxMatcherInitialize{
		matchers.StartFloatMatcher,
		matchers.StartIntegerMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Float\n")
	} else if to.Value != "1.1" {
		t.Errorf("Found thing should be '1.1' not %s\n", to.Value)
	} else if to.Type != matchers.FLOAT {
		t.Errorf("Found tokens should be type FLOAT not %s\n", to.Type)
	}
}

func TestLexFindsTwoDots(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("1..", []matchers.LexxMatcherInitialize{
		matchers.StartFloatMatcher,
		matchers.StartIntegerMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Float\n")
	} else if to.Value != "1" {
		t.Errorf("Found thing should be '1' not %s\n", to.Value)
	} else if to.Type != matchers.INTEGER {
		t.Errorf("Found tokens should be type FLOAT not %s\n", to.Type)
	}
}

func TestLexxFindsLeadingFloat(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("893.5a", []matchers.LexxMatcherInitialize{
		matchers.StartFloatMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "893.5" {
		t.Errorf("Found thing should be '893.5' not %s\n", to.Value)
	} else if to.Type != matchers.FLOAT {
		t.Errorf("Found tokens should be type FLOAT not %s\n", to.Type)
	}
}

func TestLexxFindsMultipleFloats(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("436.2 33343.444 42.42 2.0001", []matchers.LexxMatcherInitialize{
		matchers.StartFloatMatcher,
		matchers.StartWhitespaceMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "436.2" {
		t.Errorf("Found token should be '436.2' not %s\n", to.Value)
	} else if to.Type != matchers.FLOAT {
		t.Errorf("Found tokens should be type FLOAT not %s\n", to.Type)
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
	} else if to.Value != "33343.444" {
		t.Errorf("Found token should be '33343.444' not %s\n", to.Value)
	} else if to.Type != matchers.FLOAT {
		t.Errorf("Found tokens should be type FLOAT not %s\n", to.Type)
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
	} else if to.Value != "42.42" {
		t.Errorf("Found token should be '42.42' not %s\n", to.Value)
	} else if to.Type != matchers.FLOAT {
		t.Errorf("Found tokens should be type FLOAT not %s\n", to.Type)
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
	} else if to.Value != "2.0001" {
		t.Errorf("Found token should be '2.0001' not %s\n", to.Value)
	} else if to.Type != matchers.FLOAT {
		t.Errorf("Found tokens should be type FLOAT not %s\n", to.Type)
	}
}
