package matchers_tests

import (
	"github.com/JeffThomas/lexx/lexx"
	"github.com/JeffThomas/lexx/matchers"
	"testing"
)

////////////////////////////////
// IdentifierMatcher
////////////////////////////////

func TestLexxFindsIdentifier(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("identifier", []matchers.LexxMatcherInitialize{
		matchers.StartIdentifierMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Identifier\n")
	} else if to.Value != "identifier" {
		t.Errorf("Found thing should be 'identifier' not %s\n", to.Value)
	} else if to.Type != matchers.IDENTIFIER {
		t.Errorf("Found tokens should be type IDENTIFIER not %s\n", to.Type)
	}
}

func TestLexxFindsIdentifierWithLeadingUnderscore(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("_identifier", []matchers.LexxMatcherInitialize{
		matchers.StartIdentifierMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Identifier\n")
	} else if to.Value != "_identifier" {
		t.Errorf("Found thing should be '_identifier' not %s\n", to.Value)
	} else if to.Type != matchers.IDENTIFIER {
		t.Errorf("Found tokens should be type IDENTIFIER not %s\n", to.Type)
	}
}

func TestLexxFindsIdentifierWithLeadingAndEmbeddedUnderscore(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("__ident_ifi_er_", []matchers.LexxMatcherInitialize{
		matchers.StartIdentifierMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Identifier\n")
	} else if to.Value != "__ident_ifi_er_" {
		t.Errorf("Found thing should be '__ident_ifi_er_' not %s\n", to.Value)
	} else if to.Type != matchers.IDENTIFIER {
		t.Errorf("Found tokens should be type IDENTIFIER not %s\n", to.Type)
	}
}

func TestLexyFindsIdentifierWithNumberAndEmbeddedUnderscore(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("__2ident_ifi_er_", []matchers.LexxMatcherInitialize{
		matchers.StartIdentifierMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Identifier\n")
	} else if to.Value != "__2ident_ifi_er_" {
		t.Errorf("Found thing should be '__2ident_ifi_er_' not %s\n", to.Value)
	} else if to.Type != matchers.IDENTIFIER {
		t.Errorf("Found tokens should be type IDENTIFIER not %s\n", to.Type)
	}
}

func TestLexxFindsIdentifierNotKeyword(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("newThing", []matchers.LexxMatcherInitialize{
		matchers.ConfigKeywordMatcher([]string{"new"}),
		matchers.StartIdentifierMatcher,
	})
	to, _ := lexxx.GetNextToken()

	if to == nil {
		t.Errorf("Lexx did not find Identifier\n")
	} else if to.Value != "newThing" {
		t.Errorf("Found identifier should be 'newThing' not %s\n", to.Value)
	} else if to.Type != matchers.IDENTIFIER {
		t.Errorf("Found tokens should be type IDENTIFIER not %s\n", to.Type)
	}
}

func TestLexxFindsIdentifierNotKeywordOrWord(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("newThing", []matchers.LexxMatcherInitialize{
		matchers.ConfigKeywordMatcher([]string{"new"}),
		matchers.StartWordMatcher,
		matchers.StartIdentifierMatcher,
	})
	to, _ := lexxx.GetNextToken()

	if to == nil {
		t.Errorf("Lexx did not find Identifier\n")
	} else if to.Value != "newThing" {
		t.Errorf("Found identifier should be 'newThing' not %s\n", to.Value)
	} else if to.Type != matchers.IDENTIFIER {
		t.Errorf("Found tokens should be type IDENTIFIER not %s\n", to.Type)
	}
}

func TestLexxFindsMultipleIdentifiers(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("ident_ifi_er_ _id ", []matchers.LexxMatcherInitialize{
		matchers.StartIdentifierMatcher,
		matchers.StartWhitespaceMatcher,
	})
	to, _ := lexxx.GetNextToken()

	if to == nil {
		t.Errorf("Lexx did not find Identifier\n")
	} else if to.Value != "ident_ifi_er_" {
		t.Errorf("Found identifier should be 'ident_ifi_er_' not %s\n", to.Value)
	} else if to.Type != matchers.IDENTIFIER {
		t.Errorf("Found tokens should be type IDENTIFIER not %s\n", to.Type)
	}

	to, _ = lexxx.GetNextToken()
	to, _ = lexxx.GetNextToken()

	if to == nil {
		t.Errorf("Lexx did not find Identifier\n")
	} else if to.Value != "_id" {
		t.Errorf("Found identifier should be '_id' not %s\n", to.Value)
	} else if to.Type != matchers.IDENTIFIER {
		t.Errorf("Found tokens should be type IDENTIFIER not %s\n", to.Type)
	}
}

func TestLexxDoesNotFindIdentifierWithLeadingNumber(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("2ident_ifi_er_", []matchers.LexxMatcherInitialize{
		matchers.StartIdentifierMatcher,
	})
	to, _ := lexxx.GetNextToken()

	if to.Type != matchers.UNDEFINED {
		t.Errorf("Lexx found incorrect tokens, should be nil but it's %s\n", to.Value)
	}
}

func TestLexxDoesNotFindeEmpty(t *testing.T) {
	lexxx := lexx.BuildLexxWithString("", []matchers.LexxMatcherInitialize{
		matchers.StartIdentifierMatcher,
	})
	to, _ := lexxx.GetNextToken()

	if to == nil {
		t.Errorf("Lexx did not find EOF\n")
	} else if to.Value != "EOF" {
		t.Errorf("Found identifier should be 'EOF' not %s\n", to.Value)
	} else if to.Type != matchers.SYSTEM {
		t.Errorf("Found tokens should be type SYSTEM not %s\n", to.Type)
	}
}
