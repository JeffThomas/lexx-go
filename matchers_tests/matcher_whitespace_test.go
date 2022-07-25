package matchers_tests

import (
	"github.com/JeffThomas/lexx/lexx"
	"github.com/JeffThomas/lexx/matchers"
	"testing"
)

////////////////////////////////
// Whitespace Matcher
////////////////////////////////

func TestLexxFindsWhitespaces(t *testing.T) {
	lexxx := lexx.BuildLexxWithString(" ", []matchers.LexxMatcherInitialize{
		matchers.StartWhitespaceMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " " {
		t.Errorf("Found token should be ' ' not %s\n", to.Value)
	} else if to.Type != matchers.WHITESPACE {
		t.Errorf("Found tokens should be type WHITESPACE not %s\n", to.Type)
	}
}

func TestLexxFindsMultipleWhitespaces(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("  ", []matchers.LexxMatcherInitialize{
		matchers.StartWhitespaceMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != "  " {
		t.Errorf("Found token should be '  ' not %s\n", to.Value)
	} else if to.Type != matchers.WHITESPACE {
		t.Errorf("Found tokens should be type WHITESPACE not %s\n", to.Type)
	}
}

func TestLexxFindsMultipleWhitespacesWithTabs(t *testing.T) {
	lexxx := lexx.BuildLexxWithString(" 	 ", []matchers.LexxMatcherInitialize{
		matchers.StartWhitespaceMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " 	 " {
		t.Errorf("Found token should be ' \t ' not %s\n", to.Value)
	} else if to.Type != matchers.WHITESPACE {
		t.Errorf("Found tokens should be type WHITESPACE not %s\n", to.Type)
	}
}

func TestLexxFindsMultipleWhitespacesWithTabsAndLineBreaks(t *testing.T) {
	lexxx := lexx.BuildLexxWithString(" 	\n ", []matchers.LexxMatcherInitialize{
		matchers.StartWhitespaceMatcher,
		matchers.StartIntegerMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " 	\n " {
		t.Errorf("Found token should be ' \t\n ' not %s\n", to.Value)
	} else if to.Type != matchers.WHITESPACE {
		t.Errorf("Found tokens should be type WHITESPACE not %s\n", to.Type)
	}
}

func TestLexxFindsMultipleWhitespacesWithTabsAndLineBreaksCheckLineAdvance(t *testing.T) {
	lexxx := lexx.BuildLexxWithString(" 	\n1 ", []matchers.LexxMatcherInitialize{
		matchers.StartWhitespaceMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " 	\n" {
		t.Errorf("Found token should be ' \t\n' not %s\n", to.Value)
	} else if to.Type != matchers.WHITESPACE {
		t.Errorf("Found tokens should be type WHITESPACE not %s\n", to.Type)
	}
	to, _ = lexxx.GetNextToken() // skip the 1
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Whitespace\n")
	} else if to.Value != " " {
		t.Errorf("Found token should be ' ' not %s\n", to.Value)
	} else if to.Type != matchers.WHITESPACE {
		t.Errorf("Found tokens should be type WHITESPACE not %s\n", to.Type)
	} else if to.Line != 1 {
		t.Errorf("Line did not advance to 1, it is %d\n", to.Line)
	}
}

func TestLexxFindsEOF(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("", []matchers.LexxMatcherInitialize{
		matchers.StartWhitespaceMatcher,
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

func TestLexxFindsWhitespacesBetweenNumbers(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("4 5 6 7", []matchers.LexxMatcherInitialize{
		matchers.StartWhitespaceMatcher,
		matchers.StartIntegerMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Integer\n")
	} else if to.Value != "4" {
		t.Errorf("Found token should be '436.2' not %s\n", to.Value)
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
	} else if to.Value != "5" {
		t.Errorf("Found token should be '33343.444' not %s\n", to.Value)
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
	} else if to.Value != "6" {
		t.Errorf("Found token should be '42.42' not %s\n", to.Value)
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
	} else if to.Value != "7" {
		t.Errorf("Found token should be '2.0001' not %s\n", to.Value)
	} else if to.Type != matchers.INTEGER {
		t.Errorf("Found tokens should be type INTEGER not %s\n", to.Type)
	}
}
